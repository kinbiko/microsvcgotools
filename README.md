# microsvcgotools

## Run the panicking app

Run the following to compare the output with and without [panicparse](https://github.com/maruel/panicparse).

```bash
go run ./boom/cmd/main.go #without panicparse
go run ./boom/cmd/main.go 2>&1 | pp
```

## Run and manually test gRPC server with gRPCurl

The `fizzbuzz` package contains a server that exposes a gRPC endpoint for getting "fizzbuzz" values for numbers in a certain interval.

Use [`gRPCurl`](https://github.com/fullstorydev/grpcurl) to hit the endpoint without writing a gRPC client.

```bash
# Run the gRPC server
go run ./fizzbuzz/grpc/main.go
# Hit the gRPC server endpoint with JSON
grpcurl -proto fizzbuzz/fizzbuzz.proto -plaintext -d '{"start": 1, "end": 100}' localhost:1234 fizzbuzz.FizzBuzz/GetFizzBuzzSequence

# Simplify this by creating a temporary alias:
alias grpcall="grpcurl -proto fizzbuzz/fizzbuzz.proto -plaintext -d @ localhost:1234 fizzbuzz.FizzBuzz/GetFizzBuzzSequence"
# Then simply:
echo '{"start": 0, "end": 100}' | grpcall
```

### Manually test through a web interface

Take `gRPCurl` a step further and get a nice web UI with [`gRPCui`](https://github.com/fullstorydev/grpcui).

```bash
# Start a web UI for hitting the same server (works because reflection is enabled)
grpcui -plaintext localhost:1234
```

### Load test the gRPC server with ghz

Use [`ghz`](https://github.com/bojand/ghz) to load test your gRCP server.

```bash
ghz \
--insecure \
-d '{"start": 1, "end": 600}' \
-c 100 \
-n 5000 \
--call=fizzbuzz.FizzBuzz/GetFizzBuzzSequence \
--connections=5 \
localhost:1234
```

### Generate and explore the gRPC server under load

Assuming you have built and are running the binary `main` (built from ./fizzbuzz/grpc/main.go), run [pprof](https://github.com/google/pprof) to get profile your application.
You probably want to use ghz to generate some load so your application isn't idle.

```bash
go tool pprof -http=":" main http://localhost:4321/debug/pprof/profile
```
