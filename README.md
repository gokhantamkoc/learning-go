# **Learning Go**

## **Hello Go Developer**

Let's analyse `/1-helloworld/main.go` file declaration by declaration:

The building block of a Go Program is package.

Every Go Program starts from `package main`. Below is a **package declaration**.

```
package main
```

Packages can be imported to a Go program with `import` declaration:

```
import "fmt"
```

Every Go program has a main function declaration. It is the entry point of a Go program.

```
func main() {
   fmt.Println("Hello, Go Developer!")
}
```

## **Package Definition**

By convention, the package name is the same as the last element of the import path.

For instance, the `math/rand` package comprises files that begin with the statement `package rand`.

## **Importing Packages**

There are two ways to import packages.

1. Using `factored import`:

```
import (
	"fmt"
	"math"
)
```

2. Using single `package import`:

```
import "fmt"
import "math"
```

It is a good practice to use the factored import statement.

## **Exported Names**

In Go Lang, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run `the code`. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again.

## **Functions**



## **Named Return Values**



## **Variables**



## **Variables with Initializers**



## **Short Variable Declarations**



## **Basic Types**

Go Lang has below basic types:

Name|Description|Type|Size
-|-|-|-
bool|Boolean data type|Logical| 1 bit
string|String literals|Character Sequence|Depends on character sequence
int|Integer data type|Numeric|8 byte
int8|Integer data type|Numeric|1 byte
int16|Integer data type|Numeric|2 byte
int32|Integer data type|Numeric|4 byte
int64|Integer data type|Numeric|8 byte
uint|Unsigned integer data type|Numeric|4 byte
uint8|Unsigned integer data type|Numeric|1 byte
uint16|Unsigned integer data type|Numeric|2 byte
uint32|Unsigned integer data type|Numeric|4 byte
uint64|Unsigned integer data type|Numeric|8 byte
uintptr|Unsigned integer data type|Numeric|4 byte
byte| This basic type is alias for int8|Numeric| 1 byte
rune| This basic type is alias for int8. This type also represents a Unicode code point|Numeric|1 byte
float32||Numeric|4 byte
float64||Numeric|8 byte
complex64||Numeric|8 byte
complex128||Numeric|16 byte

> Variable declarations can be *factored* into blocks, like `import` statements.

> When you need an integer value you should use `int` unless you have a specific reason to use a `sized` or `unsigned integer type`.

## **Zero (Default) Values**

Variables which initial value are not declared are given their zero (default) value.

Type| Default Value
-|-
Numeric| 0
String Literal|Empty String
Logical|false

## **Type Conversions**



## **Type Inference**



## **Constants**



## **Numeric Constants**



## **For Loop**



## **If Statement**

