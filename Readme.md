# Go Tutorials

This docs is for the beginner who start learn Go Language.

## Setup

Install Go [click here](https://golang.org/doc/install)

## Editor plugins and IDEs

Personaly I recommend the [VS Code](https://code.visualstudio.com/) as your editor.

> vs code plugin: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go

Read more [here](https://golang.org/doc/editors.html)

## Hellow World

```sh
mkdir workspace && cd workspace
touch index.go
```

Open the `workspace` by `vscode`.

Add the following code to `index.go`

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, world, aicoder.com")
}
```

Build and run the code

```sh
cd workspace
go build index.go
./index
```

We can get the output as below:

```sh
Hello, world, aicoder.com
```
