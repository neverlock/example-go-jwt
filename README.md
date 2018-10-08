# example-go-jwt

curl -i -v -H "Content-Type: application/json" -X POST --data '{"Username":"haku","Password":"testing"}' http://localhost:5000/token-auth

curl -i -v -H "Authorization: Bearer xxxxxxxxxx" http://localhost:5000/test/hello
