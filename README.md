# Golang and dropzone

This is a basic example about how to implement [dropzonejs](http://www.dropzonejs.com) using [golang](https://golang.org/)
as server language.

### Usage

1. Go to your `GOPATH` path and run `git clone https://github.com/gustavocd/golang-and-dropzone`.
2. Change directory `cd golang-and-dropzone` and run `go get` to download the dependencies.
3. Run `go run main.go`, you should see the message *server running* in your terminal.
4. Visit `http://localhost:8081` and start uploading images :dancer:.

### Note:

This is a basic implementation as I said before, it doesn't implement success and error handlers, the main goal of this example
is to set up a base for a better implementation feel free to improve and change this code, but if it helps you please let me know.