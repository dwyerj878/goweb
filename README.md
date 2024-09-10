
# Testbed App for GO development


## Web API

uses 
* gin
* valkey
* mongo


### ToDo / Roadmap

- [x] implement gin
- [x] implement env variable settings
- [x] implement redis/valkey
- [ ] implement redis/valkey settings
- [x] implement mongo
- [x] implement swagger/openapi
- [x] implement authentication
- [ ] add ssl/tls
- [ ] re-implement settings as a singleton
- [ ] add api testing with mocks and or test containers
- [ ] add otel metrics/tracing
- [ ] add custom logging with trac/span id

### Sample payloads

#### encrypt data

`curl -X POST 127.0.0.1:3000/encrypt --data '{"text" : "some text", "key" : "passphrasewhichneedstobe32bytes!" } '`

#### decrypt data

`curl -X POST 127.0.0.1:3000/decrypt --data '{"text" : "JsBtosJ3BqMQW9GsX0sWntFsvlI3cq422uyF4XUzKj84HN/etQ==", "key" : "passphrasewhichneedstobe32bytes!" } '`

#### decrypt fail

 `curl -v -X POST 127.0.0.1:3000/decrypt --data '{"texxxt" : "JsBtosJ3BqMQW9GsX0sWntFsvlI3cq422uyF4XUzKj84HN/etQ==" } '`

 #### get settings

 `curl -X GET 127.0.01:3000/keys`

#### create user

`POST': curl -v -X POST 127.0.0.1:3000/user --data '{"user_name" : "fred2", "plain_text_password" : "password", "full_name" : "fred bear" }'`

#### use protected route with above credentials

`curl -v -X GET 127.0.0.1:3000/user/fredxx --header 'Authorization:Basic ZnJlZDI6cGFzc3dvcmQ='`



## Swagger implementation

see [https://github.com/swaggo/gin-swagger]

### generate swagger 

`~/go/bin/swag init -g hello.go`

### swagger URL

`http://127.0.0.1:3000/swagger/index.html`

## VSCode

### launch.json

assumes mongo and valkey running locally

```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch file",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "src/hello.go",
            "console": "integratedTerminal",
            
            "env": {
                "VALKEY_HOST" : "127.0.0.1:6379",
                "MONGO_HOST" : "mongodb://127.0.0.1:27017",

            }
        }
    ]
}
```

### docker setup

docker run --name valkey -d -p 6379:6379 valkey/valkey

docker run --name mongodb -d -p 27017:27017 mongo


### docker start 

docker restart valkey

docker restart mongodb

