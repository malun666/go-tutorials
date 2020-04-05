# Go Tutorials

此教程是为初学 Go 语言的小伙伴准备的简明版本的教程。
This docs is for the beginner who start to learn Go Language.

> 教程会有中英文的翻译。

## Setup

安装 Go 语言环境。
Install Go [click here](https://golang.org/doc/install)

## Editor plugins and IDEs

个人来说建议使用 [VS Code](https://code.visualstudio.com/) 开发 go 程序。

Personaly I recommend the [VS Code](https://code.visualstudio.com/) as your editor.

> vs code plugin: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go

Read more [here](https://golang.org/doc/editors.html)

## Hellow World

第一个 Go 程序。
Your first Go program.

```sh
mkdir workspace && cd workspace
touch index.go
```

用`vscode`打开`workspace`文件夹。

Open the `workspace` by `vscode`.

添加`index.go`文件，并添加以下代码：
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
