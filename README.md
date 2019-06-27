# microsvcgotools

## Run the panicking app

Run the following to compare the output with and without [panicparse](https://github.com/maruel/panicparse).

```bash
go run ./boom/cmd/main.go #without panicparse
go run ./boom/cmd/main.go |& pp
```

## Run and manually test gRPC server with gRPCurl

The `fizzbuzz` package contains a server that exposes a gRPC endpoint for getting "fizzbuzz" values for numbers in a certain interval.

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

```bash
# Start a web UI for hitting the same server (works because reflection is enabled)
grpcui -plaintext localhost:1234
```

### Load test the gRPC server with ghz

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

```bash
go tool pprof http://localhost:4321/debug/pprof/profile
# Start hitting the gRPC endpoint with ghz as above
```
