### Go modules test

Figuring out how dependency management works in Go.

`mkdir something`
`touch something.go`
`go mod init $something`
`go mod vendor`
`go build -mod=vendor`
`go mod tidy`
