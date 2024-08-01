# Go Lang client app to connect to redis or valkey server and get, set, patch entries using following APIs
To Run the application:
1. cd redis-client-app
2. Update the Redis/Valkey redisConnectionString in file repository/repository.go
3. Run the app using command ```go run main.go```

**APIs**
1. **POST /api/entries:** To add a new entry(key, value) on Redis/Valkey Server <br />
Request Body:<br /> 
```
{
"Key":"key1",
"Value":"value1"
}
```
Status Code: 202 <br />
Response Body:<br /> 
```
{
"Key":"key1",
"Value":"value1"
}
```
2. **GET /api/entries:** Get all entries(key, value) from Redis/Valkey Server <br />
Status Code: 200 <br />
Response Body:<br /> 
```
[
{
"Key":"key1",
"Value":"value1"
},
"Key":"key2",
"Value":"value2"
}
]
```
3. **GET /api/entries/{key}:** Get entrt(key, value) by Key Name from Redis/Valkey Server <br />
Status Code: 200 <br />
Response Body:<br /> 
```
{
"Key":"key1",
"Value":"value1"
}
```
4. **PATCH /api/entries/{key}:** Update value by Key Name on Redis/Valkey Server <br />
Request Body:<br /> 
```
{
"Value":"value1"
}
```
Status Code: 200 <br />
Response Body:<br /> 
```
{
"Key":"key1",
"Value":"value1"
}
```
5. **DELETE /api/entries/{key}:** Delete entry(key, value) by Key Name from Redis/Valkey Server <br />
Status Code: 200 <br />
