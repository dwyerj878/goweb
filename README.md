
# Testbed App for GO development


## Web API

uses gin

### ToDo

* ~~implement gin~~
* ~~implement env variable settings~~
* ~~implement redis/valkey~~
* implement redis/valkey settings
* implement mongo
* ~~implement swagger/openapi~~


### sample payloads

#### encrypt data

`curl -X POST 127.0.0.1:3000/encrypt --data '{"text" : "some text", "key" : "passphrasewhichneedstobe32bytes!" } '`

#### decrypt data

`curl -X POST 127.0.0.1:3000/decrypt --data '{"text" : "JsBtosJ3BqMQW9GsX0sWntFsvlI3cq422uyF4XUzKj84HN/etQ==", "key" : "passphrasewhichneedstobe32bytes!" } '`

#### decrypt fail

 `curl -v -X POST 127.0.0.1:3000/decrypt --data '{"texxxt" : "JsBtosJ3BqMQW9GsX0sWntFsvlI3cq422uyF4XUzKj84HN/etQ==" } '`

 #### get settings

 `curl -X GET 127.0.01:3000/keys`



## Swagger impl

see [https://github.com/swaggo/gin-swagger]

## VSCode

### launch.json

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
                "MONGO_HOST" : "127.0.0.1:27017",

            }
        }
    ]
}
```