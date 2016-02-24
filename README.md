## Building
To build an executable, you can use the standard `go install` from your GOPATH root

```bash
go install github.com/jessicalevine/hellocgo
```

However, if you would like to see the intermediate .c and .go files output by cgo, you can use (assuming you're in the project directory):

```bash
go tool cgo -objdir output_directory hellocgo.go
```
