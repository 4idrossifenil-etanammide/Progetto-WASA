/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// Struct che verranno usate nelle operazioni per il db
type UserName struct {
	Name string `json:"name"`
}

type ID struct {
	Id string `json:"id"`
}

type Photo struct {
	PhotoID       string    `json:"photoID"`
	UploadingDate time.Time `json:"date"`
	LikeNumber    int       `json:"likeNumber"`
	CommentNumber int       `json:"commentNumber"`
}

type Comment struct {
	Name string `json:"name"`
	Text string `json:"comment"`
}

type Profile struct {
	Photos      []Photo
	PhotoNumber int
	Follower    []UserName
	Following   []UserName
}

type Stream struct {
	Photos []Photo `json:"photos"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	CheckToken(string, string) error
	LoginUser(UserName) (ID, error)

	SetName(string, string) error

	UploadPhoto(string, string) (Photo, error)
	DeletePhoto(string, string) error

	LikePhoto(string, string) error
	UnlikePhoto(string, string) error

	BanUser(string, string) error
	UnbanUser(string, string) error
	CheckBan(string, string) error

	CommentPhoto(string, Comment) (UserName, error)
	DeleteComment(string, string) error

	Follow(string, string) error
	Unfollow(string, string) error

	GetStream(string) (Stream, error)
	GetProfile(string, string) (Profile, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// ==================== ATTIVAZIONE VINCOLI REFERENZIALI ====================

	sqlStmt := `PRAGMA foreign_keys = ON;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// ==================== CREAZIONE TABELLA UTENTE ====================

	err = createTable(db, "Utente", `CREATE TABLE Utente(ID TEXT NOT NULL PRIMARY KEY, Nome TEXT NOT NULL);`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// ==================== CREAZIONE TABELLA FOTO ====================

	err = createTable(db, "Foto", `CREATE TABLE Foto(
									FotoID TEXT NOT NULL,
									Utente TEXT NOT NULL,
									Data DATE NOT NULL, 
									nLikes INTEGER NOT NULL, 
									nCommenti INTEGER NOT NULL,
									Path TEXT NOT NULL,
									CONSTRAINT fk_utente FOREIGN KEY(Utente) REFERENCES Utente(ID) ON DELETE CASCADE ON UPDATE CASCADE,
									PRIMARY KEY(FotoID, Utente));`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	/* Questo pezzo di codice serve a far funzionare i vincoli di integrità refernziali.
	Incredibile come il SQLite permetta di definire tabelle senza chiavi primarie, chiavi primarie con valori nulli (e allora che differenza c'è tra
	UNIQUE e PRIMARY KEY) o ancora che non abbia i vincoli di integrità referenziali attivati di default (sennò non ti fai chiamare db relazionale) e
	tante altre cose senza senso, tra cui l'impossibilità di definire foreign keys su valori non UNIQUE o PRIMARY KEYS.
	*/
	createUniqueIndex(db, "Foto", "j", "FotoID")
	if err != nil {
		return nil, fmt.Errorf("Error creating unique index: %w", err)
	}

	// ==================== CREAZIONE TABELLA LIKE ====================

	err = createTable(db, "Like", `CREATE TABLE Like(
									LikeID INTEGER PRIMARY KEY AUTOINCREMENT,
									FotoReference TEXT NOT NULL, 
									Utente TEXT NOT NULL, 
									CONSTRAINT fk_foto FOREIGN KEY(FotoReference) REFERENCES Foto(FotoID) ON DELETE CASCADE ON UPDATE CASCADE, 
									CONSTRAINT fk_utente FOREIGN KEY(Utente) REFERENCES Utente(ID) ON DELETE CASCADE ON UPDATE CASCADE,
									UNIQUE(FotoReference, Utente));`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// ==================== CREAZIONE TABELLA COMMENTI ====================

	err = createTable(db, "Commenti", `CREATE TABLE Commenti(
										CommentiID INTEGER PRIMARY KEY autoincrement, 
										FotoReference TEXT NOT NULL, 
										Utente TEXT NOT NULL, 
										Testo TEXT NOT NULL, 
										CONSTRAINT fk_foto FOREIGN KEY(FotoReference) REFERENCES Foto(FotoID) ON DELETE CASCADE ON UPDATE CASCADE, 
										CONSTRAINT fk_utente FOREIGN KEY(Utente) REFERENCES Utente(ID) ON DELETE CASCADE ON UPDATE CASCADE);`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// ==================== CREAZIONE TABELLA SEGUE ====================

	err = createTable(db, "Segue", `CREATE TABLE Segue(
										Utente TEXT NOT NULL, 
										Follower TEXT NOT NULL, 
										CONSTRAINT fk_utente FOREIGN KEY(Utente) REFERENCES Utente(ID) ON DELETE CASCADE ON UPDATE CASCADE, 
										CONSTRAINT fk_follower FOREIGN KEY(Follower) REFERENCES Utente(ID) ON DELETE CASCADE ON UPDATE CASCADE, 
										PRIMARY KEY(Utente, Follower));`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// ==================== CREAZIONE TABELLA BAN ====================

	err = createTable(db, "Ban", `CREATE TABLE Ban(
									Utente TEXT NOT NULL, 
									Bannato TEXT NOT NULL, 
									CONSTRAINT fk_utente FOREIGN KEY(Utente) REFERENCES Utente(ID) ON DELETE CASCADE ON UPDATE CASCADE, 
									CONSTRAINT fk_bannato FOREIGN KEY(Bannato) REFERENCES Utente(ID) ON DELETE CASCADE ON UPDATE CASCADE, 
									PRIMARY KEY(Utente, Bannato));`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func createTable(db *sql.DB, tableName string, query string) error {
	initialQuery := fmt.Sprintf(`SELECT name FROM sqlite_master WHERE type='table' AND name='%s';`, tableName)
	err := db.QueryRow(initialQuery).Scan(&tableName)
	if err != nil {
		return err
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	if errors.Is(err, sql.ErrNoRows) {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func createUniqueIndex(db *sql.DB, tableName string, indexName string, column string) error {

	var seq string
	var name string
	var unique string
	var creationMode string
	var partialIndex string
	var createNew bool = true

	query := fmt.Sprintf(`PRAGMA INDEX_LIST(%s);`, tableName)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		err = rows.Scan(&seq, &name, &unique, &creationMode, &partialIndex)
		if err != nil {
			return err
		}

		if name == indexName {
			createNew = false
		}

	}
	if err := rows.Err(); err != nil {
		return err
	}

	if createNew {
		query = fmt.Sprintf(`CREATE UNIQUE INDEX %s ON %s(%s)`, indexName, tableName, column)
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
