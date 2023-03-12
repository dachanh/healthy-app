# arent-healthy-app

# arent-healthy-app

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

API Document :

# Login

Request :

```
POST /user/login
Host: 127.0.0.1:8176
Content-Type: application/json
Content-Length: 66
{
	"email":"johndoe@example.com",
	"password":"password123"
}
```
If Login successfull, the api will response user 's information

Example with curl:

```
curl --request POST \
  --url http://127.0.0.1:8176/user/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email":"johndoe@example.com",
	"password":"password123"
}'
```

# Diary

```
GET /diary/:id?page=[page]&limit=[limit]

id : user id 
page, limit: pagination 
```
response if successful, api will response user 's diary 


Exmaple with curl:
```
curl --request GET \
  --url http://127.0.0.1:8176/diary/1
```

```
[
  {
    "ID": 1,
    "content": "Today was a good day!",
    "created_at": "2023-03-11T12:40:31.848111Z",
    "updated_at": "2023-03-11T12:40:31.848111Z"
  }
]
```

# Excerise History 

```
GET /excerice-history/:id

id : User 's id
```

```
curl --request GET \
  --url http://127.0.0.1:8176/excerice-history/1
```

# News 

```
curl --request GET \
  --url http://127.0.0.1:8176/news
```

