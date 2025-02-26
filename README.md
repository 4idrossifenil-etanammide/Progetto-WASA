# ⚠️ Disclaimer: Uni Project Repo 🚧
Before you dive into this repo, here are a few things you should know:

📌 This repository was created for the course _Web and software architecture (2023)_ – meaning it was built with deadlines, specific exam requirements, and a whole lot of learning along the way.

🛠️ Expect rough edges – You might find some less-than-ideal code, typos, missing documentation, or even bits of unused logic hanging around. It happens.

📚 Not a polished product – This isn't an actively maintained project, so don’t expect updates or refactors. If something looks weird, it probably has a reason (or I just didn’t know better at the time 🤷).

💡 I was learning! – The code here reflects my understanding at that point in time, so don’t judge too hard. Everyone starts somewhere!

🗣️ **Italian language ahead!** – The project was developed for a university course in a **Computer Science Bachelor's program taught in Italian**, so you’ll find comments, documentation, and possibly variable names in Italian.

***

## Running the application

For building the docker file, navigate to the folder of the project, then use: <br />
	&emsp; docker build -t nomeImmagine:latest -f nomeFile . <br />
	&emsp; ( ex. docker build -t wasaphoto-backend:latest -f Dockerfile.backend . <br />
	&emsp;      docker build -t wasaphoto-frontend:latest -f Dockerfile.frontend . ) 
	
Running the container: <br />
	&emsp; docker run -it -p porta:porta --rm nomeImmagine:latest <br />
	&emsp; ( ex. docker run -it -p 3000:3000 --rm wasaphoto-backend:latest <br />
	&emsp;      docker run -it -p 8080:80 --rm wasaphoto-frontend:latest ) 
	     
After running the container, go to http://localhost:8080/ to see the project
