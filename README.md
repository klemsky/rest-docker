# rest-docker
This app is to fetch total marks of every Mezink students in the record tables of Mezink database.
The database and sample data are pre-defined using docker compose

Pre-requisite:
1. Docker installed on your OS
2. Postman (optional)

How to run the app:
1. Simply run ```docker compose up --build```. Application will run on the your local on port: 8080
2. Hit the endpoint using below contract
3. To stop the service, run ```docker compose down``` or ```docker compose down -v``` (to remove the docker containers completely), and then run docker compose on point (1)

Here is the API contract of the records API:
https://docs.google.com/document/d/1dF-LNngEvJEuvYZiBp7s-NEzpu08iZ5GoHdHGbNsVfw/edit?usp=sharing
