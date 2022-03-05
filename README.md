# learn-golang

Use `--target dev` to build a docker image based on the target defined in your dockerfile - go use case for using docker as a development environment.

Entry point of an application is the main() function.

Compiling (go build) for windows will give you an exe binary. 

Keep in mind the reduced impact on memory consumption when using local variables. 

Arrays are fixed in size in Go, usually better to use slices. Slices are dynamic and are not copied in memory like arrays. 

```
go mod init github.com/../..
```

'flag' package is used to build CLI tools in go. 

Building an http server with endpoints using the built in http package in go. 

Define your own webserver and define a port for it to serve over. 

Handle function (can be any type of function)

Read Header values using the HandleFunc

