# Server_CLI
## Server

Build Image
```
docker build -t server .
```
Start docker
```
docker run -p 9000:9000 server
```
## Client

Build client
```
make build
```
Start client
```
./bin/client -str=<string> -url=<http://localhost:9000/api/substring>
```
