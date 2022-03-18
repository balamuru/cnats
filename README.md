# cnats
A C-Golang NATS client bridge

## Compilation
* Compile a C .so `go build -o cnats.so -buildmode=c-shared cnats.go`
* Compile the C executable using the .so created earlier `gcc -o client client.c ./cnats.so`

## Testing
* Spawn a temporary nats container listening on the port configured in the sample client (TCP 4555) `docker run --rm -it  -p 4555:4222  nats -DV`
* Launch the C client `./client` which publishes a test message to `nats://localhost:4555`
* Observe the NATS broker logs to see that the message has been received
```
[1] 2022/03/18 18:26:47.183212 [TRC] 172.17.0.1:59262 - cid:4 - "v1.13.0:go" - <<- [PUB test.subject 11]
[1] 2022/03/18 18:26:47.183249 [TRC] 172.17.0.1:59262 - cid:4 - "v1.13.0:go" - <<- MSG_PAYLOAD: ["hello world"]
```

