#  REST API app that continuously monitors the price of Bitcoin  

A Simple Golang project that continuously monitors the price of Bitcoin and performs the basic GET CRUD operation, the Bitcoin price is stored/persisted in a DB which is hosted on a volume.It also alerts a given email when the price either goes above or below given limits

## DevOps Steps 
All scripts to run Kuberentes files are located under **.k8s**. <br/>

Steps to execute in local:
Ensure you have golang, docker installed

1. Checkout code
2. Update .env file
3. docker compose build
4. docker compose up
5. Curl/Postman to test using localhost
   
- curl -X GET http://localhost:8080/api/prices/btc?date=03-10-2022&offset=2&limit=3

## High Level Overview
![image](https://user-images.githubusercontent.com/37391853/150505929-88d01921-360b-4dbe-aa33-1943e4f8911f.png)


For more information/references in detail:
- https://docs.docker.com/engine/reference/commandline/compose/
- https://docs.docker.com/compose/environment-variables/


