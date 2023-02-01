package database

import (
	"hash/fnv"
	"strconv"
	"time"
)

/*
La funzione esegue la query alla ricerca dell'id corrispondente al nome passato in input. La funzione Next() ritorna True fintanto
che ci sono valori da leggere, quindi si controlla che il risultato non sia vuoto e si restituisce il corrispondente id, altrimenti viene generato
e vengono inseriti i dati nel database
*/
func (db *appdbimpl) LoginUser(u UserName) (ID, error) {

	rows, err := db.c.Query(`SELECT id FROM Utente WHERE nome = ?`, u.Name)
	if err != nil {
		return ID{}, err
	}
	defer func() { _ = rows.Close() }()

	if rows.Next() {
		var id ID
		err = rows.Scan(&id.Id)
		if err != nil {
			return ID{}, err
		}
		return id, nil
	}
	if err := rows.Err(); err != nil {
		return ID{}, err
	}
	uniqueID := generateUniqueId(u.Name)
	_, err = db.c.Exec(`INSERT INTO Utente (id, nome) VALUES (?, ?)`, uniqueID, u.Name)
	if err != nil {
		return ID{}, err
	}
	return ID{Id: uniqueID}, nil
}

func generateUniqueId(name string) string {
	h := fnv.New32a()
	h.Write([]byte(time.Now().Format("02-01-2006 15:04:05")))
	return strconv.Itoa(int(h.Sum32()))
}
