binary := "orcabak"
binDir := "./bin"

alias b := build
alias t := test

@_default:
  just --list

# Build locally
build:
  mkdir -p {{binary}}
  go build -o {{binDir}}/{{binary}} ./cmd/orcabak

# Run local binary
run *ARGS:
  {{binDir}}/{{binary}} {{ARGS}}

# Run tests
test:
  go test ./...

# Install globally
install:
  go install

# Clean local binaries
clean:
  rm -rf {{binDir}}
