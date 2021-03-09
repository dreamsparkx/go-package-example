# Introduction

Go was designed to encourage good software engineering practices. One of the guiding principles of high-quality software is the DRY principle - Don’t Repeat Yourself, which basically means that you should never write the same code twice. You should reuse and build upon existing code as much as possible.

Functions are the most basic building blocks that allow code reuse. Packages are the next step into code reusability. They help you organize related Go source files together into a single unit, making them modular, reusable, and maintainable.

# Go Package

In the most basic terms, A package is nothing but a directory inside your Go workspace containing one or more Go source files, or other Go packages.

Every Go source file belongs to a package. To declare a source file to be part of a package, we use the following syntax -

```
package <packagename>
```

The above package declaration must be the first line of code in your Go source file. All the functions, types, and variables defined in your Go source file become part of the declared package.

You can choose to export a member defined in your package to outside packages, or keep them private to the same package. Other packages can import and reuse the functions or types that are exported from your package.

### Let’s see an example

Almost all the code that we have seen so far in this tutorial series include the following line -

```
import "fmt"
```

fmt is a core library package that contains functionalities related to formatting and printing output or reading input from various I/O sources. It exports functions like Println(), Printf(), Scanf() etc, for other packages to reuse.

Packaging functionalities in this way has the following benefits -

It reduces naming conflicts. You can have the same function names in different packages. This keeps our function names short and concise.

It organizes related code together so that it is easier to find the code you want to reuse.

It speeds up the compilation process by only requiring recompilation of smaller parts of the program that has actually changed. Although we use the fmt package, we don’t need to recompile it every time we change our program.

## The main package and main() function

Go programs start running in the main package. It is a special package that is used with programs that are meant to be executable.

By convention, Executable programs (the ones with the main package) are called Commands. Others are called simply Packages.

The main() function is a special function that is the entry point of an executable program.

## Importing Packages

There are two ways to import packages in Go -

```
// Multiple import statements
import "fmt"
import "time"
import "math"
import "math/rand"
```

```
// Factored import statements
import (
	"fmt"
	"time"
	"math"
    "math/rand"
)
```

Go’s convention is that - the package name is the same as the last element of the import path. For example, the name of the package imported as math/rand is rand. It is imported with path math/rand because It is nested inside the math package as a subdirectory.

## Exported vs Unexported names

Anything (variable, type, or function) that starts with a capital letter is exported, and visible outside the package.

Anything that does not start with a capital letter is not exported, and is visible only inside the same package.

When you import a package, you can only access its exported names.

```
package main

import (
	"fmt"
	"math"
)

func main() {
	// MaxInt64 is an exported name
	fmt.Println("Max value of int64: ", int64(math.MaxInt64))

	// Phi is an exported name
	fmt.Println("Value of Phi (ϕ): ", math.Phi)

	// pi starts with a small letter, so it is not exported
	fmt.Println("Value of Pi (𝜋): ", math.pi)
}
```

```
# Output
./exported_names.go:16:38: cannot refer to unexported name math.pi
./exported_names.go:16:38: undefined: math.pi
```

To fix the above error, you need to change math.pi to math.Pi.

Note: Go module is Go’s new dependency management system. A module is a collection of Go packages stored in a directory with a go.mod file at its root. The go.mod file defines the module’s path, which is also the import path used while importing packages that are part of this module.

Before Go module got introduced in Go 1.11, every project needed to be created inside the so-called GOPATH. The path of the project inside GOPATH was considered its import path.

Let’s initialize a Go module by typing the following commands:

```
$ go mod init example.com
```

Credit:

https://www.callicoder.com/golang-packages/
