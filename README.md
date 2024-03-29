# set sql creds in .env config file

# userStore start(default server port 8080)
go run main.go -port=8000

# curl examples
curl -X GET http://localhost:8000/profile -H "Api-key: ffff-2918-xcas"
curl -X GET http://localhost:8000/profile/admin -H "Api-key: ffff-2918-xcas"