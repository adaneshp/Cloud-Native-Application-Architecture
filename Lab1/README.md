# Lab 1: Go Basics

- This code was developed in January, 2022.

## Set up lab platform (for all labs)
Recommend using Ubuntu Linux (20.04) installed on a Virtualbox VM for all labs.
If your computer does not have sufficient resources (cores, memory) to smoothly run a VM, you could either 1) Use Linux as the only OS, or 2) Dual boot current OS with Linux
You are free to use other Linux distributions, MacOS, or even Windows for all labs. However, I may not be able to help you if you run into issues. 

## All instructions are for Linux.
Create  a CloudNative/labs directory for all your programs. I am assuming you are sufficiently familiar with Linux to do this from the command line.

In this lab we will be writing basic Go programs.

## Install Go
Follow instructions as provided here to download and install Go in your system
https://golang.org/doc/install

Create directory for GOPATH setting.
```
mkdir ~/go
```
For GOPATH environment variable, add the following lines to the ~/.bashrc
```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH 
```
If you need help on configuring paths, bashrc see -
https://linuxhint.com/export-a-path-in-bashrc/


## Configure your favorite editor for Go

For Vim editor (My editor of choice. Use only if you are familiar with Vim)
vim-go plugin
https://github.com/fatih/vim-go-tutorial

For Visual Studio Code 
https://dev.to/ko31/how-to-setup-golang-with-vscode-1i4i

Go plugins are also available for Atom and Sublime editors

## Writing a simple integer adder program as a package


Go uses a new dependency management system starting from Go 1.11 using Go Modules
A module is a collection of Go packages stored in a file tree with a go.mod file at its root. The go.mod file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build.


Type in following command under the labs directory:
```
go mod init labs
```

This creates a go.mod file with a name of the module and the Go version. All dependencies required for building the module will be automatically included here.

Create a lab1 directory under labs directory

We will develop the integer adder program as a Go package

Create a myadder directory under lab1. All adder package files will be in this directory. It is a good practice to name the directory have the same name as the package (myadder in our case)

Copy and paste the following code in a file add.go in the myadder directory.

```
// Returns the sum of two input integers
package myadder

func Add(x, y int) int {
    return x + y 
    //return 42
}
```

Testing is built into Go. The test file is in the same package and is named with “<filename>_test.go”. 
Copy and paste the following code in a file add_test.go in the myadder directory.
```
package myadder

import "testing"

func TestAdd(t *testing.T) {
    want := 7
    got := Add(3, 4)
    if want != got {
        t.Errorf("Error in myadder.Add; Want 7, Got %d", got)
    }   
}
```
Run the test as follows from command line as follows:
```
go test
```
This test should pass. 
Now modify add.go by commenting out the statement return x + y, and uncommenting the statement return 42.

Run the test again. The test should fail.
This completes the first version of the adder package. We can extend the package by adding other functions, and user defined data types. For each component added, you would modify the test to ensure that the component is functionally correct. 

Let’s now use the Add function from outside the package. Be sure that the test above is passing.
From the lab1 directory, cut and paste the following code to a file named main.go

```
package main

import (
    "fmt"
    "labs/lab1/myadder"
)

func main() {
    fmt.Println(myadder.Add(5, 6)) 
}
```

Notice how the myadder package is referred to with respect to the labs module path.
Now build and run main.go from the command line. Ensure that the test is passing prior to this.
```
go run main.go
```
This should print the output of 11.
Note: You can build an executable using go build, and then run it separately.

Congratulations! You have written your first Go program.
In summary, a typical set of steps to write Go code
In the development directory, initialize module with go mod init
Break your project into a collection of packages
Incrementally, develop the code for the package, testing along the way
From the module, integrate the packages in main.go for your overall project.


## Extra
Write a Go program that reads a text file, and outputs the top K occurrences of words in the file. Words are defined as a set of characters limited by a whitespace.
Write your code in a directory corresponding to the package name (textproc)

Hint 1: Use string package for string processing
Hint 2: Use os package for opening a file
Hint 3: Use bufio, and io package for file reading
Hint 4: Use a map for counting word occurrences in O(n) time
Hint 5: Helper function for sorting is provided


The code for the Extra section can be seen "textproc" directory.
