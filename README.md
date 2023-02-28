## Running the application

For building the docker file, navigate to the folder of the project, then use: 
	docker build -t nomeImmagine:latest -f nomeFile . 
	( ex. docker build -t wasaphoto-backend:latest -f Dockerfile.backend . 
	      docker build -t wasaphoto-frontend:latest -f Dockerfile.frontend . ) 
	
Running the container: 
	docker run -it -p porta:porta --rm nomeImmagine:latest 
	( ex. docker run -it -p 3000:3000 --rm wasaphoto-backend:latest 
	      docker run -it -p 8080:80 --rm wasaphoto-frontend:latest ) 
	     
After running the container, go to http://localhost:8080/ to see the project
