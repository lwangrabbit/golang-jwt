
### /login
```
curl -X POST http://localhost:9099/login -H "Content-type:application/json" -d '{"username": "admin", "password": "admin"}'
```
will return token
```
200
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEwMDEsImV4cCI6MTYzNjEwMTcwNiwiaXNzIjoiVEVTVDAwMSJ9.BFw-E-Pmq5mZ5aRCydIbpJQN6TF6k4618sQXGibi-0g"}
```

### /api/order
```
curl -X POST http://localhost:9099/api/order -H "Content-type:application/json" -H "token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEwMDEsImV4cCI6MTYzNjEwMTcwNiwiaXNzIjoiVEVTVDAwMSJ9.BFw-E-Pmq5mZ5aRCydIbpJQN6TF6k4618sQXGibi-0g" -d '{"product": "apple", "count": 3}'
```
will return response
```
200
"Hi 1001, I will give you 3 apple"
```

when give invalid token:
```
curl -X POST http://localhost:9099/api/order -H "Content-type:application/json" -H "token:xyz" -d '{"product": "apple", "count": 3}'
```
will return response
```
403
"Invalid token"
```
