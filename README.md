# Server_CLI
## Server

#uild Image
```
docker build -t server .
```
Start docker
```
docker run -p 9000:9000 server
```
## Client

#uild client
```
make build
```
Start client
```
./bin/client -str=<string> -url<http://localhost:9000/api/substring>
```
