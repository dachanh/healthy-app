
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

this api will be responsed on `My record` page

```
GET /diary/:id

id : user id 
```
List all diary information of the user
response if successful, api will response user 's diary 


Exmaple with curl:
```
curl --request GET \
  --url http://127.0.0.1:8176/diary/1
```

# Excerise History 
this api will be responsed on `My record` page

```
GET /exercise-history/:id

id : User 's id
```

Example with curl:
```
curl --request GET \
  --url http://127.0.0.1:8176/exercise-history/1
```

# News 
This api will response News on the `Column page`

```
GET  /news
```

```
curl --request GET \
  --url http://127.0.0.1:8176/news
```

# Meal history

this api will be reponsed the meal of user and display on `Top page`

```
GET  /meal-history/:id
id : the user 's id
```

Example with curl:
```
curl --request GET \
  --url http://127.0.0.1:8176/meal-history/1
```

# Meal history by Session

if session = 1 , the meal for morning <br/>
if session = 2 , the meal for lunch <br/>
if session = 3 , the meal for dinner <br/>

```
GET /meal-history/:id/:session
id : the user 's id
session: the session id that I define above
```

Example with curl:

```
curl --request GET \
  --url http://127.0.0.1:8176/meal-history/1/1
```
