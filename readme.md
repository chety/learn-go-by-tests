### DOCUMENTATION

- Creating documentation file `godoc -http :8000`
- If you do not have godoc then install `go install golang.org/x/tools/cmd/godoc@latest`
- Navigate to `http://localhost:8000/pkg/` and enjoy documentation

### TESTING

- To run tests execute `go test -v`
- If you want to run benchmarks execute `go test -bench=.`
- If you want to see coverage execute `go test -cover`
- We can create a coverage report(here for fmt package) by running `go test -covermode=count -coverprofile=count.out fmt`
- To open coverage report in your browser type `go tool cover -html=count.out`. You can rename coverage profile as you wish
- We do not have to provide `covermode`. You can read details [here](https://go.dev/blog/cover)
- If you want to see each function coverage report type `go tool cover -func=count.out `
