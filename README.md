# Readme
## Guide to run functional test
### Prerequisite
1. Make sure you run installed go version >= `1.15.3`
2. Run your program to serve port `8080`.

### Run Functional Test
1. `cd parking-lot-golang`
2. `go mod download`
3. To run first case: `go run functional_testing.go -url=http://localhost:8080 -case=1`
4. To run second case: `go run functional_testing.go -url=http://localhost:8080 -case=2`
