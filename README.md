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