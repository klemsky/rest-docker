# rest-docker
This app is to fetch total marks of every Mezink students in the record tables of Mezink database.
The database and sample data are pre-defined using docker. 

Pre-requisite:
1. Docker installed on your OS
2. Postman (optional)
3. Golang (optional)

How to run the app:
1. Simply run ```docker compose up --build```. Application will be built & run in your local on port: 8080
2. Hit the endpoint using below contract attached on google docs.
3. To stop the service, run ```docker compose down``` or ```docker compose down -v``` (to remove the docker containers completely), and then run docker compose on point (1) to build & run the app again.

Here is the API contract of the records API:
https://docs.google.com/document/d/1X0n1H1t-1v_lWIxcB9FIHhsxrvJaoPfRuyI8EPozQ-k/edit?usp=sharing

You can import a Postman collection by downloading the ```postman.json``` file inside this repository and try import it using Postman application.
