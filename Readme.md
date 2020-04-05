# Go Tutorials(Go 语言教程)

此教程是为初学 Go 语言的小伙伴准备的简明版本的教程。  
This docs is for the beginner who start to learn Go Language.

> 教程会有中英文的翻译。

## Setup(Go 语言安装)

安装 Go 语言环境。

Install Go [click here](https://golang.org/doc/install)

## Editor plugins and IDEs

个人来说建议使用 [VS Code](https://code.visualstudio.com/) 开发 go 程序。

Personaly I recommend the [VS Code](https://code.visualstudio.com/) as your editor.

> vs code plugin: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go

Read more [here](https://golang.org/doc/editors.html)

## Hellow World(第一个 Go 程序)

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

编译上面的程序，并运行。

Build and run the code.

```sh
cd workspace
go build index.go
./index
```

> Go 命令行工具，除了编译之外，还有格式化、测试等一系列工具。 文档地址：[这里](https://golang.org/doc/cmd)
> Go command tool has several commands, e.g., doc, build, fix, fmt. Read more [here](https://golang.org/doc/cmd)

在控制台中获得如下输出：

We can get the output as below:

```sh
Hello, world, aicoder.com
```

## The Go Programming Language Specification(Go 语言语法规范)

### Notation (文档符号规范)

- `|`: alternation(或)
- `()`: grouping(分组)
- `[]`: option (0 or 1 times)(出现 0 次或者 1 次)
- `{}`: repetition (0 to n times)(出现次数范围，出现 0 次到 n 次)

### Comments（注释）

Comments serve as program documentation. There are two forms:

注释是写给程序员来阅读代码的，不会编译到最终程序中。Go 语言有两种注释（类似 C、Java、C#、JavaScript)

- Line comments start with the character sequence `//` and stop at the end of the line.
- General comments start with the character sequence `/*` and stop with the first subsequent character sequence `*/`.

```go
i:= 3 // Line comments , 行注释

/*
  General comments, mutiple lines comments
  多行注释
*/
```

### 标识符

标识符用于声明函数、接口、类型、变量等的名字。

#### 保留字和关键字

```sh
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

#### 预定义标识符

- Types:

```sh
bool byte complex64 complex128 error float32 float64
int int8 int16 int32 int64 rune string
uint uint8 uint16 uint32 uint64 uintptr
```

- Constants:

```sh
true false iota
```

- Zero value:

```sh
nil
```

- Functions:

```sh
append cap close complex copy delete imag len
make new panic print println real recover
```

### 运算符

```sh
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```
