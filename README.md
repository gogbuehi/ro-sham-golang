# Golang RoShamGolang

## Learning Go from the ground up

### Purpose
 - Learn basic patterns in Go
 - Establish a template for building applications using Docker
 - Make a worthwhile mini-game (in this case, Ro-Sham-Bo; aka Rock, Paper, Scissors)

### Status
 - 09 Mar 2022: A simple "Hello World" Go app that runs in a Docker container

### Prerequisites
 - Docker - [Install Docker](https://docs.docker.com/get-docker/)

### How to Run

Run the following commands:
 - `docker-compose build`
 - `docker-compose run`

If any changes are made to the application, `docker-compose build` will need to be run again.
This allows the changes to the Go application to be built inside the Docker container.

If the application just needs to be run and has already been built, then `docker-compose up` is all that is needed.
