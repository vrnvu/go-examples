# effective go

Notes taken from [effective-go](https://golang.org/doc/effective_go.html) at 05/01/2021

- [effective go](#effective-go)
  - [Formatting](#formatting)
  - [Commentary](#commentary)
  - [Names](#names)
    - [Package names](#package-names)
    - [Getters](#getters)
    - [Interface names](#interface-names)
    - [MixedCaps](#mixedcaps)
    - [Semicolons](#semicolons)
  - [Control structures](#control-structures)
    - [If](#if)
    - [Redeclaration and reassignment](#redeclaration-and-reassignment)
    - [For](#for)
    - [Switch](#switch)
    - [Type switch](#type-switch)
  - [Functions](#functions)
    - [Multiple return values](#multiple-return-values)
    - [Named result paraemters](#named-result-paraemters)
    - [Defer](#defer)
  - [Data](#data)
    - [Allocation with new](#allocation-with-new)
    - [Constructors and composite literals](#constructors-and-composite-literals)
    - [Allocation with make](#allocation-with-make)
    - [Arrays](#arrays)
    - [Slices](#slices)
    - [Two-dimensional slices](#two-dimensional-slices)
    - [Maps](#maps)
    - [Printing](#printing)
    - [Append](#append)
    - [Initialization](#initialization)
    - [Constants](#constants)
    - [Variables](#variables)
    - [The init function](#the-init-function)
  - [Methods](#methods)
    - [Pointers vs Values](#pointers-vs-values)
  - [Interfaces and other types](#interfaces-and-other-types)
    - [Interfaces](#interfaces)
    - [Conversions](#conversions)
    - [Interface conversions and type assertions](#interface-conversions-and-type-assertions)
    - [Generality](#generality)
    - [Interfaces and methods](#interfaces-and-methods)
  - [The blank identifier](#the-blank-identifier)

## Formatting

Use *gofmt* 

## Commentary

Go provides C-style /* */ block comments and C++-style // line comments. Line comments are the norm; block comments appear mostly as package comments, but are useful within an expression or to disable large swaths of code.

**Every package should have a package comment**. It will appear first on the godoc page and should set up the detailed documentation that follows. 

Inside a package, any comment immediately preceding a top-level declaration serves as a doc comment for that declaration. **Every exported (capitalized) name in a program should have a doc comment.**

If every doc comment begins with the name of the item it describes, **you can use the doc subcommand of the go tool and run the output through grep**. Imagine you couldn't remember the name "Compile" but were looking for the parsing function for regular expressions, so you ran the command, 

```bash
$ go doc -all regexp | grep -i parse
    Compile parses a regular expression and returns, if successful, a Regexp
    MustCompile is like Compile but panics if the expression cannot be parsed.
    parsed. It simplifies safe initialization of global variables holding
$
```

## Names

### Package names

**By convention, packages are given lower case, single-word names.**

The package name is only the default name for imports; it need not be unique across all source code, and in the rare case of a collision the importing package can choose a different name to use locally. In any case, confusion is rare because the file name in the import determines just which package is being used. 

### Getters

No

### Interface names

**By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification** to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc. 

### MixedCaps

Use MixedCaps or mixedCaps

### Semicolons

Normally the programmer does not need to write them explicitly.

Idiomatic Go programs have semicolons only in places such as for loop clauses, to separate the initializer, condition, and continuation elements. They are also necessary to separate multiple statements on a line, should you write code that way. 

One consequence of the semicolon insertion rules is that **you cannot put the opening brace of a control structure (if, for, switch, or select) on the next line**. If you do, a semicolon will be inserted before the brace, which could cause unwanted effects. 

## Control structures

The control structures of Go are related to those of C but differ in important ways. 

- There is no do or while loop, only a slightly generalized for

- switch is more flexible

- if and switch accept an optional initialization statement like that of for; 

- break and continue statements take an optional label to identify what to break or continue
  
- and there are new control structures including a type switch and a multiway communications multiplexer, select. The syntax is also slightly different: there are no parentheses and the bodies must always be brace-delimited.

### If

In Go a simple if looks like this: 

```go
if x > 0 {
    return y
}
```

Since if and switch accept an initialization statement, it's common to see one used to set up a local variable. 

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

**If an if statement doesn't flow into the next statement**—that is, the body ends in break, continue, goto, or return—**the unnecessary else is omitted.**

### Redeclaration and reassignment

In a := declaration a variable v may appear even if it has already been declared, provided: 

- this declaration is in the same scope as the existing declaration of v
- the corresponding value in the initialization is assignable to v 
- there is at least one other variable that is created by the declaration

### For 

It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons. 

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

**If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.**

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

If you only need the first item in the range (the key or index), drop the second.

If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first.

### Switch

The expressions need not be constants or even integers, the cases are evaluated top to bottom until a match is found, and if the switch has no expression it switches on true. It's therefore possible—and idiomatic—to write an if-else-if-else chain as a switch. 

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

There is no automatic fall through, but cases can be presented in comma-separated lists. 

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```
### Type switch

A switch can also be used to discover the dynamic type of an interface variable. Such a type switch uses the syntax of a type assertion with the keyword type inside the parentheses. If the switch declares a variable in the expression, the variable will have the corresponding type in each clause. It's also idiomatic to reuse the name in such cases, in effect declaring a new variable with the same name but a different type in each case. 


```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```

## Functions

### Multiple return values

One of Go's unusual features is that functions and methods can return multiple values. 

```go
func (file *File) Write(b []byte) (n int, err error)
```

### Named result paraemters

The return or result "parameters" of a Go function **can be given names and used as regular variables**, just like the incoming parameters. When named, they are **initialized to the zero values** for their types when the function begins; if the function executes a return statement with no arguments, the current values of the result parameters are used as the returned values. 

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

### Defer

Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns.

Deferring a call to a function such as Close has two advantages. First, it guarantees that you will never forget to close the file, a mistake that's easy to make if you later edit the function to add a new return path. Second, it means that the close sits near the open, which is much clearer than placing it at the end of the function. 

## Data

### Allocation with new

Go has two allocation primitives, the built-in functions new and make. 

Let's talk about **new** first. It's a built-in function that **allocates memory, but unlike its namesakes in some other languages it does not initialize the memory, it only zeros it.** 

### Constructors and composite literals

Sometimes the zero value isn't good enough and an initializing constructor is necessary.

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

We can simplify it using a composite literal, which is an expression that creates a new instance each time it is evaluated. 

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

**Note that, unlike in C, it's perfectly OK to return the address of a local variable; the storage associated with the variable survives after the function returns.**

```go
    return &File{fd, name, nil, 0}
```

### Allocation with make

**It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T)**. The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use. 

Remember that make applies only to maps, slices and channels and does not return a pointer. To obtain an explicit pointer allocate with new or take the address of a variable explicitly. 

### Arrays

There are major differences between the ways arrays work in Go and C. In Go:

- **Arrays are values. Assigning one array to another copies all the elements.**

- In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it. 

- The size of an array is part of its type. The types [10]int and [20]int are distinct. 

### Slices

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimension such as transformation matrices, most array programming in Go is done with slices rather than simple arrays. 

**Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.** If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array. 

### Two-dimensional slices

Go's arrays and slices are one-dimensional. To create the equivalent of a 2D array or slice, it is necessary to define an array-of-arrays or slice-of-slices, like this: 

```go
type Transform [3][3]float64  // A 3x3 array, really an array of arrays.
type LinesOfText [][]byte     // A slice of byte slices.
```

### Maps

The key can be of any type for which the equality operator is defined. Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller. 

### Printing

Formatted printing in Go uses a style similar to C's printf family but is richer and more general. **The functions live in the fmt package and have capitalized names: fmt.Printf, fmt.Fprintf, fmt.Sprintf and so on**. The string functions (Sprintf etc.) return a string rather than filling in a provided buffer. 

### Append

```go
func append(slice []T, elements ...T) []T
```

You can't actually write a function in Go where the type T is determined by the caller. That's why append is built in: it needs support from the compiler. 

### Initialization

Although it doesn't look superficially very different from initialization in C or C++, initialization in Go is more powerful. Complex structures can be built during initialization and the ordering issues among initialized objects, even among different packages, are handled correctly. 

### Constants

Constants in Go are just that—constant.**They are created at compile time**, even when defined as locals in functions, and can only be numbers, characters (runes), strings or booleans. Because of the compile-time restriction, **the expressions that define them must be constant expressions, evaluatable by the compiler**. 

### Variables

Variables can be initialized just like constants but the initializer can be a general expression computed at run time. 

### The init function

Finally, each source file can define its own niladic init function to set up whatever state is required. (Actually each file can have multiple init functions.) And finally means finally: init is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized. 

**Besides initializations that cannot be expressed as declarations, a common use of init functions is to verify or repair correctness of the program state before real execution begins.**

## Methods

### Pointers vs Values

The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers. 

## Interfaces and other types

### Interfaces

Interfaces with only one or two methods are common in Go code, and are usually given a name derived from the method, such as io.Writer for something that implements Write. 

### Conversions

It's an idiom in Go programs to convert the type of an expression to access a different set of methods. 

### Interface conversions and type assertions

Type switches are a form of conversion: they take an interface and, for each case in the switch, in a sense convert it to the type of that case. 

### Generality

If a type exists only to implement an interface and will never have exported methods beyond that interface, there is no need to export the type itself. Exporting just the interface makes it clear the value has no interesting behavior beyond what is described in the interface. It also avoids the need to repeat the documentation on every instance of a common method. 

In such cases, the constructor should return an interface value rather than the implementing type. 

### Interfaces and methods

Since almost anything can have methods attached, almost anything can satisfy an interface. One illustrative example is in the http package, which defines the Handler interface. Any object that implements Handler can serve HTTP requests. 

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Here's a trivial implementation of a handler to count the number of times the page is visited. 

```go
// Simple counter server.
type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```

For reference, here's how to attach such a server to a node on the URL tree.

```go
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
```

But why make Counter a struct? An integer is all that's needed. (The receiver needs to be a pointer so the increment is visible to the caller.) 

```go
// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}
```

What if your program has some internal state that needs to be notified that a page has been visited? Tie a channel to the web page.

```go
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")

```

Finally, let's say we wanted to present on /args the arguments used when invoking the server binary. It's easy to write a function to print the arguments. 

```go
func ArgServer() {
    fmt.Println(os.Args)
}
```

**Since we can define a method for any type except pointers and interfaces, we can write a method for a function**. 

```go
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
```

ArgServer now has same signature as HandlerFunc, so it can be converted to that type to access its methods. The code to set it up is concise: 

```go
http.Handle("/args", http.HandlerFunc(ArgServer))
```

## The blank identifier