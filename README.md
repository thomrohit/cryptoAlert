#  REST API app that continuously monitors the price of Bitcoin  

A Simple Golang project that continuously monitors the price of Bitcoin and performs the basic GET CRUD operation, the Bitcoin price is stored/persisted in a DB which is hosted on a volume.It also alerts a given email when the price either goes above or below given limits

## Steps 
Steps to execute in local:
Ensure you have golang, docker installed

1. Checkout code
2. Update .env file
3. docker compose build
4. docker compose up
5. Curl/Postman to test using localhost
   
   - curl -X GET http://localhost:8080/api/prices/btc?date=03-10-2022&offset=2&limit=3

## High Level Overview
![highlevel](https://user-images.githubusercontent.com/37391853/193573162-c76bf12f-906c-4cab-9d54-bebc5a684eb1.png)


## For more information/references in detail:
- https://docs.docker.com/engine/reference/commandline/compose/
- https://docs.docker.com/compose/environment-variables/


