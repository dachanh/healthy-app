# arent-healthy-app

this is [diagram](https://dbdiagram.io/d/640414bf296d97641d8576f1) of this project

## Build 

first step, please make sure created folder `config/postgresql/data`

```
mkdir -p config/postgresql/data
```
Next step, Using docker-compose to build the system
```
docker-compose up -d
```
## Migrate

this step will migrate schema and create data seed, please follow the steps below

```
docker exec -it migrate.arent.healthyapp bash
make install
make migrateup
```

If you want delete database, please follow the steps below 

```
docker exec -it migrate.arent.healthyapp bash
migratedown
```

please check my api [document](./doc/document.md)

the undeveloped parts:

- Unit test : I not enough time to make unit test
- Login service : due to the lack of time and the purpose for assignment I don't have enough time the make the login service better
- CRUD : Some the api I don't have to make the full CRUD
- Optimize infrastructure: I think I make good enough with this time If I have more time I will it better
- Swagger : Because for the interview assignment I don't implement it to Front-End developer can use it 
- Document : the same as the `Swagger` I think If I write the document and give the `curl` example can help you understand the response API because I reference the data seed which is the correct data and API will response
