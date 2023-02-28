## Running the application

For building the docker file, navigate to the folder of the project, then use: <br />
	> docker build -t nomeImmagine:latest -f nomeFile . <br />
	> ( ex. docker build -t wasaphoto-backend:latest -f Dockerfile.backend . <br />
	      docker build -t wasaphoto-frontend:latest -f Dockerfile.frontend . ) 
	
Running the container: <br />
	> docker run -it -p porta:porta --rm nomeImmagine:latest <br />
	> ( ex. docker run -it -p 3000:3000 --rm wasaphoto-backend:latest <br />
	      docker run -it -p 8080:80 --rm wasaphoto-frontend:latest ) 
	     
After running the container, go to http://localhost:8080/ to see the project
