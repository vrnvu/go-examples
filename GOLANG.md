# 1. Golang reference

- [1. Golang reference](#1-golang-reference)
- [2. Notation](#2-notation)
- [3. Keywords](#3-keywords)
- [4. Constants](#4-constants)
- [5. Variables](#5-variables)
- [6. Types](#6-types)
  - [6.1. Method sets](#61-method-sets)
  - [6.2. String types](#62-string-types)
  - [6.3. Array types](#63-array-types)
  - [6.4. Slice types](#64-slice-types)
  - [6.5. Struct types](#65-struct-types)
  - [6.6. Pointer types](#66-pointer-types)
  - [6.7. Function types](#67-function-types)
  - [6.8. Interfaces types](#68-interfaces-types)
  - [6.9. Map types](#69-map-types)
  - [6.10. Channel types](#610-channel-types)
- [7. Properties of types and values](#7-properties-of-types-and-values)
  - [7.1. Type identity](#71-type-identity)
- [8. Declarations and scope](#8-declarations-and-scope)
  - [8.1. Constant declarations](#81-constant-declarations)
  - [8.2. Iota](#82-iota)
  - [8.3. Type declarations](#83-type-declarations)
    - [8.3.1. Alias declaration](#831-alias-declaration)
    - [8.3.2. Type definitions](#832-type-definitions)
  - [8.4. Variable declarations](#84-variable-declarations)
  - [8.5. Short variable declarations](#85-short-variable-declarations)
  - [8.6. Method declarations](#86-method-declarations)
- [9. Expressions](#9-expressions)
  - [9.1. Qualified identifiers](#91-qualified-identifiers)
  - [9.2. Composite literals](#92-composite-literals)
  - [9.3. Function literals](#93-function-literals)
  - [9.4. Primary expressions](#94-primary-expressions)
  - [9.5. Selectors](#95-selectors)
  - [9.6. Index expressions](#96-index-expressions)
  - [9.7. Slice expressions](#97-slice-expressions)
  - [9.8. Type assertions](#98-type-assertions)
  - [9.9. Operators](#99-operators)
    - [9.9.1. Operators precedence](#991-operators-precedence)
    - [9.9.2. Integer overflow](#992-integer-overflow)
  - [9.10. Comparison operators](#910-comparison-operators)
  - [9.11. Address operators](#911-address-operators)
  - [9.12. Receive operator](#912-receive-operator)
  - [9.13. Constant expressions](#913-constant-expressions)
- [10. Statements](#10-statements)
  - [10.1. Termianting statements](#101-termianting-statements)
  - [10.2. Send statements](#102-send-statements)
  - [10.3. Assignaments](#103-assignaments)
  - [10.4. Go statements](#104-go-statements)
  - [10.5. Select statements](#105-select-statements)
  - [10.6. Return statements](#106-return-statements)
  - [10.7. Break statements](#107-break-statements)
  - [10.8. Defer statements](#108-defer-statements)
- [11. Built-in functions](#11-built-in-functions)
  - [11.1. Close](#111-close)
  - [11.2. Length and capacity](#112-length-and-capacity)
  - [11.3. Allocation](#113-allocation)
  - [11.4. Making slices, maps and channels](#114-making-slices-maps-and-channels)
  - [11.5. Appending to and copying slices](#115-appending-to-and-copying-slices)
  - [11.6. Deletion of map elements](#116-deletion-of-map-elements)
  - [11.7. Manipulating complex numbers](#117-manipulating-complex-numbers)
  - [11.8. Handling panics](#118-handling-panics)
- [12. Packages](#12-packages)
  - [12.1. Source file organization](#121-source-file-organization)
  - [12.2. Import declarations](#122-import-declarations)
- [13. Program initialization and execution](#13-program-initialization-and-execution)
  - [13.1. The zero value](#131-the-zero-value)
  - [13.2. Package initialization](#132-package-initialization)
  - [13.3. Program execution](#133-program-execution)

# 2. Notation 

```EBNF
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

 Productions are expressions constructed from terms and the following operators, in increasing precedence: 

 ```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
 ```

# 3. Keywords

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

# 4. Constants


[^1] There are boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants. Rune, integer, floating-point, and complex constants are collectively called numeric constants. 

In general, complex constants are a form of constant expression and are discussed in that section. 

Numeric constants represent exact values of arbitrary precision and do not overflow. Consequently, there are no constants denoting the IEEE-754 negative zero, infinity, and not-a-number values. 

[^1] [spec#Constants](https://golang.org/ref/spec#Constants)

# 5. Variables

A variable is a storage location for holding a value. The set of permissible values is determined by the variable's type. 



A variable's value is retrieved by referring to the variable in an expression; it is the most recent value assigned to the variable. If a variable has not yet been assigned a value, its value is the zero value for its type. 

# 6. Types

The language predeclares certain type names. Others are introduced with type declarations.

```EBNF
Type      = TypeName | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
	    SliceType | MapType | ChannelType .
```
## 6.1. Method sets

The method set of a type determines the interfaces that the type implements and the methods that can be called using a receiver of that type. 

## 6.2. String types

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type. 

## 6.3. Array types

```EBNF
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

Examples of arrays:

```go
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

TODO: Is N a constant?

## 6.4. Slice types

**A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array** fs. A slice type denotes the set of all slices of arrays of its element type. The number of elements is called the length of the slice and is never negative. **The value of an uninitialized slice is nil**. 

```EBNF
SliceType = "[" "]" ElementType .
```

A slice, once initialized, **is always associated with an underlying array that holds its elements**. A slice therefore **shares storage** with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage. 

A new, initialized slice value for a given element type T is made using the built-in function **make**, which takes a slice type and parameters specifying the length and optionally the capacity. **A slice created with make always allocates a new, hidden array to which the returned slice value refers.**

These two expressions are equivalent:

```go
make([]int, 50, 100)
new([100]int)[0:50]
```

The inner slices must be initialized individually. 

## 6.5. Struct types

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique. 

```EBNF
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName .
Tag           = string_lit .
```

A field declared with a type but no explicit field name is called an embedded field. An embedded field must be specified as a type name T or as a pointer to a non-interface type name *T, and T itself may not be a pointer type. The unqualified type name acts as the field name. 

## 6.6. Pointer types

A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer. The value of an uninitialized pointer is nil. 

```EBNF
PointerType = "*" BaseType .
BaseType    = Type .
```

## 6.7. Function types
A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is nil. 

The final incoming parameter in a function signature may have a type prefixed with .... A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter. 

## 6.8. Interfaces types

An interface type specifies a method set called its interface. A variable of interface type can store a value of any type with a method set that is any superset of the interface. Such a type is said to implement the interface. The value of an uninitialized variable of interface type is nil. 

An interface type may specify methods explicitly through method specifications, or it may embed methods of other interfaces through interface type names. 

An interface T may use a (possibly qualified) interface type name E in place of a method specification. This is called embedding interface E in T. The method set of T is the union of the method sets of T’s explicitly declared methods and of T’s embedded interfaces. 

```go
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader  // includes methods of Reader in ReadWriter's method set
	Writer  // includes methods of Writer in ReadWriter's method set
}
```

A union of method sets contains the (exported and non-exported) methods of each method set exactly once, and methods with the same names must have identical signatures. 

##  6.9. Map types

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type. The value of an uninitialized map is nil. 

**The comparison operators == and != must be fully defined for operands of the key type;** thus the key type must not be a function, map, or slice. If the key type is an interface type, these comparison operators must be defined for the dynamic key values.

## 6.10. Channel types

A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. The value of an uninitialized channel is nil.

The optional <- operator specifies the channel direction, send or receive. If no direction is given, the channel is bidirectional. 

A new, initialized channel value can be made using the built-in function make, which takes the channel type and an optional capacity as arguments: 

 ```go
make(chan int, 100)
 ```

The capacity, in number of elements, sets the size of the buffer in the channel. **If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready.**

# 7. Properties of types and values

## 7.1. Type identity

Two types are either identical or different. 

A defined type is always different from any other type. Otherwise, two types are identical if their underlying type literals are structurally equivalent; that is, they have the same literal structure and corresponding components have identical types. 

# 8. Declarations and scope

Go is lexically scoped using blocks: 

- The scope of a predeclared identifier is the universe block.
- The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block.
- The scope of the package name of an imported package is the file block of the file containing the import declaration.
- The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.
- The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.
- The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.

An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration. 

## 8.1. Constant declarations

A constant declaration binds a list of identifiers (the names of the constants) to the values of a list of constant expressions.

## 8.2. Iota

Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants. 

```go
const (
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)
```

## 8.3. Type declarations 

A type declaration binds an identifier, the type name, to a type. Type declarations come in two forms: alias declarations and type definitions. 

### 8.3.1. Alias declaration

An alias declaration binds an identifier to the given type. 

```go
type (
	nodeList = []*Node  // nodeList and []*Node are identical types
	Polar    = polar    // Polar and polar denote identical types
)
```

### 8.3.2. Type definitions

A type definition creates a new, distinct type with the same underlying type and operations as the given type, and binds an identifier to it. 

The new type is called a defined type. It is different from any other type, including the type it is created from. 

```go
type (
	Point struct{ x, y float64 }  // Point and struct{ x, y float64 } are different types
	polar Point                   // polar and Point denote different types
)

type TreeNode struct {
	left, right *TreeNode
	value *Comparable
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
```

A defined type may have methods associated with it. It does not inherit any methods bound to the given type, but the method set of an interface type or of elements of a composite type remains unchanged.

TODO: not clear

## 8.4. Variable declarations

A variable declaration creates one or more variables, binds corresponding identifiers to them, and gives each a type and an initial value. 

If a list of expressions is given, the variables are initialized with the expressions following the rules for assignments. Otherwise, each variable is initialized to its zero value. 

## 8.5. Short variable declarations 

It is shorthand for a regular variable declaration with initializer expressions but no types.

**Unlike regular variable declarations**, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. **Redeclaration does not introduce a new variable; it just assigns a new value to the original.**

## 8.6. Method declarations

A method is a function with a receiver. 

The receiver is specified via an extra parameter section preceding the method name. That parameter section must declare a single non-variadic parameter, the receiver. Its type must be a defined type T or a pointer to a defined type T. T is called the receiver base type. 

# 9. Expressions

## 9.1. Qualified identifiers 

A qualified identifier is an identifier qualified with a package name prefix. 

## 9.2. Composite literals 

Composite literals construct values for structs, arrays, slices, and maps and create a new value each time they are evaluated. 

Taking the address of a composite literal generates a pointer to a unique variable initialized with the literal's value. 

```go
var pointer *Point3D = &Point3D{y: 1000}
```

## 9.3. Function literals 

```go
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

Function literals are closures: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible. 

## 9.4. Primary expressions

Primary expressions are the operands for unary and binary expressions. 

## 9.5. Selectors 

For a primary expression x that is not a package name, the selector expression denotes the field or method f of the value x (or sometimes *x; see below). 

```go
x.f
```

A selector f may denote a field or method f of a type T, or it may refer to a field or method f of a nested embedded field of T. 

## 9.6. Index expressions

A primary expression of the form


```go
a[x]
```

denotes the element of the array, pointer to array, slice, string or map a indexed by x. The value x is called the index or map key, respectively. The following rules apply: 

If *a* is not a map:
- the index x must be of integer type or an untyped constant
- a constant index must be non-negative and representable by a value of type int
- a constant index that is untyped is given type int
- the index x is in range if 0 <= x < len(a), otherwise it is out of range

For *a* pf array type A:
- a constant index must be in range
- if x is out of range at run time, a run-time panic occurs
- a[x] is the array element at index x and the type of a[x] is the element type of A

For *a* of pointer to array type:
- a[x] is shorthand for (*a)[x]

For *a* slice of type S:
- if x is out of range at run time, a run-time panic occurs
- a[x] is the slice element at index x and the type of a[x] is the element type of S


For *a* of string type:
- a constant index must be in range if the string a is also constant
- if x is out of range at run time, a run-time panic occurs
- a[x] is the non-constant byte value at index x and the type of a[x] is byte
- a[x] may not be assigned to


For *a* of map type M:
- x's type must be assignable to the key type of M
- if the map contains an entry with key x, a[x] is the map element with key x and the type of a[x] is the element type of M
- if the map is nil or does not contain such an entry, a[x] is the zero value for the element type of M

Otherwise *a[x]* is illegal

## 9.7. Slice expressions

For a string, array, pointer to array, or slice a, the primary expression 

```go
a[low : high]
```

constructs a substring or slice. The indices low and high select which elements of operand a appear in the result. The result has indices starting at 0 and length equal to high - low.

## 9.8. Type assertions

For an expression x of interface type and a type T, the primary expression 

 ```go
x.(T)
 ```

asserts that x is not nil and that the value stored in x is of type T. The notation x.(T) is called a type assertion. 

## 9.9. Operators

Operators combine operands into expressions. 


### 9.9.1. Operators precedence

```EBNF
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||


```

### 9.9.2. Integer overflow

For unsigned integer values, the operations +, -, *, and << are computed modulo 2n, where n is the bit width of the unsigned integer's type. Loosely speaking, these unsigned integer operations discard high bits upon overflow, and programs may rely on "wrap around". 

For signed integers, the operations +, -, *, /, and << may legally overflow and the resulting value exists and is deterministically defined by the signed integer representation, the operation, and its operands. Overflow does not cause a run-time panic. A compiler may not optimize code under the assumption that overflow does not occur. For instance, it may not assume that x < x + 1 is always true. 

## 9.10. Comparison operators 

Comparison operators compare two operands and yield an untyped boolean value. 

The equality operators == and != apply to operands that are comparable. The ordering operators <, <=, >, and >= apply to operands that are ordered.

- Boolean values are comparable. Two boolean values are equal if they are either both true or both false.
- Integer values are comparable and ordered, in the usual way.
- Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.
- Complex values are comparable. Two complex values u and v are equal if both real(u) == real(v) and imag(u) == imag(v).
- String values are comparable and ordered, lexically byte-wise.
- Pointer values are comparable. Two pointer values are equal if they point to the same variable or if both have value nil. Pointers to distinct zero-size variables may or may not be equal.
- Channel values are comparable. Two channel values are equal if they were created by the same call to make or if both have value nil.
- Interface values are comparable. Two interface values are equal if they have identical dynamic types and equal dynamic values or if both have value nil.
- A value x of non-interface type X and a value t of interface type T are comparable when values of type X are comparable and X implements T. They are equal if t's dynamic type is identical to X and t's dynamic value is equal to x.
- Struct values are comparable if all their fields are comparable. Two struct values are equal if their corresponding non-blank fields are equal.
- Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.

A comparison of two interface values with identical dynamic types causes a run-time panic if values of that type are not comparable.

**Slice, map, and function values are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier nil.** Comparison of pointer, channel, and interface values to nil is also allowed and follows from the general rules above. 

## 9.11. Address operators

For an operand x of type T, the address operation &x generates a pointer of type *T to x. The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. As an exception to the addressability requirement, x may also be a (possibly parenthesized) composite literal. If the evaluation of x would cause a run-time panic, then the evaluation of &x does too. 

For an operand x of pointer type *T, the pointer indirection *x denotes the variable of type T pointed to by x. If x is nil, an attempt to evaluate *x will cause a run-time panic. 

```go
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic
```

## 9.12. Receive operator

For an operand ch of channel type, the value of the receive operation <-ch is the value received from the channel ch. T

## 9.13. Constant expressions

Constant expressions may contain only constant operands and are evaluated at compile time. 

# 10. Statements

Statements control execution. 

## 10.1. Termianting statements

A terminating statement prevents execution of all statements that lexically appear after it in the same block. The following statements are terminating: 

- A "return" or "goto" statement. 
- A call to the built-in function panic.
- A block in which the statement list ends in a terminating statement.
- An "if" statement in which:
  - the "else" branch is present, and
  - both branches are terminating statements.
- A "for" statement in which:
    - there are no "break" statements referring to the "for" statement, and
    - the loop condition is absent.
- A "switch" statement in which:
  - there are no "break" statements referring to the "switch" statement,
  - there is a default case, and
  - the statement lists in each case, including the default, end in a terminating statement, or a possibly labeled "fallthrough" statement.
- A "select" statement in which:
    - there are no "break" statements referring to the "select" statement, and
    - the statement lists in each case, including the default if present, end in a terminating statement.
- A labeled statement labeling a terminating statement.

## 10.2. Send statements

Both the channel and the value expression are evaluated before communication begins. Communication blocks until the send can proceed. A send on an unbuffered channel can proceed if a receiver is ready. A send on a buffered channel can proceed if there is room in the buffer. A send on a closed channel proceeds by causing a run-time panic. A send on a nil channel blocks forever. 

## 10.3. Assignaments

Each left-hand side operand must be addressable, a map index expression, or (for = assignments only) the blank identifier. 

A tuple assignment assigns the individual elements of a multi-valued operation to a list of variables. 

The blank identifier provides a way to ignore right-hand side values in an assignment: 

```go
_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value
```

TODO: move/copy semantics

## 10.4. Go statements

A "go" statement starts the execution of a function call as an independent concurrent thread of control, or goroutine, within the same address space. 

## 10.5. Select statements

A "select" statement chooses which of a set of possible send or receive operations will proceed. 

Execution of a "select" statement proceeds in several steps: 


- For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement. The result is a set of channels to receive from or send to, and the corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated.
- If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed.
- Unless the selected case is the default case, the respective communication operation is executed.
- If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions are evaluated and the received value (or values) are assigned.
- The statement list of the selected case is executed.

## 10.6. Return statements

A "return" statement in a function F terminates the execution of F, and optionally provides one or more result values. Any functions deferred by F are executed before F returns to its caller. 

## 10.7. Break statements 

A "break" statement terminates execution of the innermost "for", "switch", or "select" statement within the same function. 

## 10.8. Defer statements 

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking. 

# 11. Built-in functions

Built-in functions are predeclared. They are called like any other function but some of them accept a type instead of an expression as the first argument.

The built-in functions do not have standard Go types, so they can only appear in call expressions; they cannot be used as function values. 

## 11.1. Close

For a channel c, the built-in function close(c) records that no more values will be sent on the channel.

## 11.2. Length and capacity

The built-in functions len and cap take arguments of various types and return a result of type int. The implementation guarantees that the result always fits into an int. 

## 11.3. Allocation

The built-in function new takes a type T, allocates storage for a variable of that type at run time, and returns a value of type *T pointing to it. The variable is initialized as described in the section on initial values. 

## 11.4. Making slices, maps and channels

The built-in function make takes a type T, which must be a slice, map or channel type, optionally followed by a type-specific list of expressions. It returns a value of type T (not *T). The memory is initialized as described in the section on initial values. 

```go
Call             Type T     Result

make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for approximately n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n
```

## 11.5. Appending to and copying slices 

The built-in functions append and copy assist in common slice operations. For both functions, the result is independent of whether the memory referenced by the arguments overlaps. 

If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, append re-uses the underlying array. 

## 11.6. Deletion of map elements 

The built-in function delete removes the element with key k from a map m. The type of k must be assignable to the key type of m. 

 ```go
delete(m, k)
 ```

## 11.7. Manipulating complex numbers 

Three functions assemble and disassemble complex numbers. The built-in function complex constructs a complex value from a floating-point real and imaginary part, while real and imag extract the real and imaginary parts of a complex value. 

## 11.8. Handling panics

Two built-in functions, panic and recover, assist in reporting and handling run-time panics and program-defined error conditions. 

While executing a function F, an explicit call to panic or a run-time panic terminates the execution of F. Any functions deferred by F are then executed as usual. Next, any deferred functions run by F's caller are run, and so on up to any deferred by the top-level function in the executing goroutine. At that point, the program is terminated and the error condition is reported, including the value of the argument to panic. This termination sequence is called panicking. 

# 12. Packages

 Go programs are constructed by linking together packages. A package in turn is constructed from one or more source files that together declare constants, types, variables and functions belonging to the package and which are accessible in all files of the same package. Those elements may be exported and used in another package. 

## 12.1. Source file organization

Each source file consists of a package clause defining the package to which it belongs, followed by a possibly empty set of import declarations that declare packages whose contents it wishes to use, followed by a possibly empty set of declarations of functions, types, variables, and constants. 

## 12.2. Import declarations  

An import declaration states that the source file containing the declaration depends on functionality of the imported package and enables access to exported identifiers of that package. 

```go
Import declaration          Local name of Sin

import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin
```

# 13. Program initialization and execution

## 13.1. The zero value

When storage is allocated for a variable, either through a declaration or a call of new, or when a new value is created, either through a composite literal or a call of make, and no explicit initialization is provided, the variable or value is given a default value. Each element of such a variable or value is set to the zero value for its type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified. 

## 13.2. Package initialization

Within a package, package-level variable initialization proceeds stepwise, with each step selecting the variable earliest in declaration order which has no dependencies on uninitialized variables. 

## 13.3. Program execution

A complete program is created by linking a single, unimported package called the main package with all the packages it imports, transitively. The main package must have package name main and declare a function main that takes no arguments and returns no value. 

Program execution begins by initializing the main package and then invoking the function main. **When that function invocation returns, the program exits. It does not wait for other (non-main) goroutines to complete.**