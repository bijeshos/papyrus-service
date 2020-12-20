# Instructions

(In sequence)
### Terminal 1
Execute the following in Terminal 1
- `$ ./papyrus-service -listen=:8001 &`
    - Expected output:
        - `listen=:8001 caller=proxying.go:25 proxy_to=none`
        - `listen=:8001 caller=main.go:72 msg=HTTP addr=:8001`
- `$ ./papyrus-service -listen=:8002 &`
    - Expected output:
        - `listen=:8002 caller=proxying.go:25 proxy_to=none`
        - `listen=:8002 caller=main.go:72 msg=HTTP addr=:8002`
- `$ ./papyrus-service -listen=:8003 &`
    - Expected output:
        - `listen=:8003 caller=proxying.go:25 proxy_to=none`
        - `listen=:8003 caller=main.go:72 msg=HTTP addr=:8003`
- `$ ./papyrus-service -listen=:8080 -proxy=localhost:8001,localhost:8002,localhost:8003`
    - Expected output:
        - `listen=:8080 caller=proxying.go:29 proxy_to="[localhost:8001 localhost:8002 localhost:8003]"`
        - `listen=:8080 caller=main.go:72 msg=HTTP addr=:8080`        
        
### Terminal 2
Execute the following in Terminal 2
- `$ for s in foo bar baz ; do curl -d"{\"s\":\"$s\"}" localhost:8080/uppercase ; done`
    - Expected output:
        - `{"v":"FOO"}`
        - `{"v":"BAR"}`
        - `{"v":"BAZ"}`       

### Terminal 1
Verify the following in Terminal 1
- Expected output:
    - `listen=:8001 caller=logging.go:28 method=uppercase input=foo output=FOO err=null took=5.168µs`
    - `listen=:8080 caller=logging.go:28 method=uppercase input=foo output=FOO err=null took=4.39012ms`
    - `listen=:8002 caller=logging.go:28 method=uppercase input=bar output=BAR err=null took=5.445µs`
    - `listen=:8080 caller=logging.go:28 method=uppercase input=bar output=BAR err=null took=2.04831ms`
    - `listen=:8003 caller=logging.go:28 method=uppercase input=baz output=BAZ err=null took=3.285µs`
    - `listen=:8080 caller=logging.go:28 method=uppercase input=baz output=BAZ err=null took=1.388155ms`

## ----------------------
## v2
### Terminal 1
Execute the following in Terminal 1
- `$ go build`
    - to build the project
- `$ ./papyrus-service`
    - to run the app
    - Expected output: `msg=HTTP addr=:8080`

### Terminal 2
Execute the following in Terminal 2
- `$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase`
    - Expected output: `{"v":"HELLO, WORLD"}`
- `$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count`
    - Expected output: `{"v":12}`
        
        
