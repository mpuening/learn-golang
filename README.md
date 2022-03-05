Learn Go Language
=================

This project contains my musings learning Go. There will not be anything
fancy of note here. Just fumblings of a Java programmer figuring out how
to do equivalent techniques with Go.

Dirt Simple Hello World
=======================

Once Go is installed, one can create a simple file `main.go` as follows:

```
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World!")
}
```

And run the application as follows:

```
go run main.go
```

One can build the application into an executable 'myapp` with this command:

```
go build -o myapp main.go
```

Third Party Dependencies
========================
While the most simple Go project can consist of a `.go` file in the current directory,
conventions toward larger projects have the `main` program be placed into a `cmd`
subdirectory, and reference third party code as follows for an example of `cmd/main.go`

```
package main

import (
    c "github.com/TreyBastian/colourize"

    "fmt"
)

func main() {
    fmt.Println(c.Colourize("Hello World!", c.Red))
}
```

To help Go manage dependencies, one needs to initialize the module
system for the application as shown with this command:

```
go mod init learn-golang
```

Then to run that application, one would use the following command:

```
go run cmd/main.go
```

But that might cause an error due to the missing dependency on the `colourize`
library. To solve that issue, one would run this command:

```
go get github.com/TreyBastian/colourize
```

That library would get downloaded to the following example directory:

```
$GOPATH/pkg/mod/github.com/!trey!bastian/colourize@v0.0.0-20160510094957-c352b0ac60d7
```

Then, running the application again should succeed.

Makefile
========

A lot can go into running the Go commands, so it is logical to wrap them up into
a `Makefile`. The file in this project was gathered from free examples on the internet.

To see what the `Makefile` can do, run `make help`

```
$make help

 Choose a command run in server:

  build     Build app.
  start     Start in development mode.
  stop      Stop development mode.
  compile   Compile the binary.
  test      Test the binaries.
  clean     Clean build files. Runs `go clean` internally.
  ```

Code Structure
==============

Big applications are hardly contained in one single file. To split an
application into manageable components, Go uses the `internal` package
convention. One can place code into sub-directories under the `internal`
directory (sibling to the `cmd` directory) and import the code as follows:

```

import (
	"learn-golang/internal/subpackage"
)
```

where the prefix to `internal` matches the `mod` name.

At this point, a developer should now have parity with how Java
code is laid out and how third party code is managed. Now it is a
manner of learning the fine points of the language and what third
party libraries make one more productive.

Front End UI
============
A simple Vue.js front end UI is provided in the `frontend` directory.

Docker Support
==============

To build the docker image, run this command:

```
sudo docker build -t learn-golang .
```

To run the docker image, run this command:

```
sudo docker run -d -p 8080:8080 learn-golang
```
