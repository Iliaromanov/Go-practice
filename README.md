# Go-practice
This repo contains the work I do as I learn the Go programming language.

## Command Line

- Create .mod file:
   > `go mod init`

- Install package:
   > `go get [-u] [package(s)]`
   
   eg:
   > `go get -u google.golang.org/grpc`

   The -u flag instructs get to use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

<!--Hidden Notes:
    * Event-based architecture: production, consumtion, reaction to events (eg. transaction is event, update state is reaction)
-->
