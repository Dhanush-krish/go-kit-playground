### repo contains go api basics

### Go test commands with descriptions
```
# Run all tests in the current package
go test

# Run tests with detailed output (shows each test)
go test -v

# Run tests in all sub-packages (recursively)
go test ./...

# Run only a specific test by name
go test -run TestAdd

# Run tests and measure coverage (summary only)
go test -cover

# Run tests and save coverage profile to a file
go test -coverprofile=coverage.out

# Show function-level coverage report from the profile
go tool cover -func=coverage.out

# Open HTML view in browser (green = covered, red = uncovered)
go tool cover -html=coverage.out

# Run all tests and show coverage for all sub-packages
go test -cover ./...

# Run tests with a minimum coverage requirement (e.g., 80%)
# If coverage is below 80%, tests fail
go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//' | awk '{ if ($1 < 80) { print "Coverage below 80%!"; exit 1 } }'

# Run tests with race detector (detects race conditions)
go test -race

# Run benchmarks (functions with BenchmarkXxx)
go test -bench .

# Run tests with timeout (default is 10m)
go test -timeout 30s

```