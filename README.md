# Go
***

### Package
---
Every go program is made up of packages.

Programs start running in package main.

### Import
---
import "fmt"
import "math"

import (
	"fmt"
	"math"
)

### Exported Names
---
A name exported if it starts with a capital name like math.Pi

### Function
---
func main () {
	fmt.Println (add (42, 13))
}

Integer return
func add (x, y int) int {
	return x + y, x
}

Multiple results
func add (x int, y int) (int, int) {
	return x + y, x
}


Named return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

### Variable
---
Variables without initializers

var c, python, java bool

Variables with initializers

var i, j int = 1, 2

var (
	i int = 1
	j int = 2
)

Short variable declaration (must be within a function)

k := 3

### Print
---
fmt.Println ()

fmt.Printf ()

fmt.Sprint ()

%b = binary

%t = boolean

%v = value

%T = type

### Types
---
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

### Constant
---
Contstant variables do not change
const Pi = 3.14

### For loop / While loop
---
init; condition; post

for i := 0; i < 10; i++ {
	fmt.Println (i)
}

init and post statments are optional

sum := 1
for sum < 1000 {
	sum += sum
}

### If
---
v = 5
if v < 10 {
}

If can start with variable declaration and also available in else blocks

if v := 5; v < 10 {
}
else {
	v += 5
}


### Switch
---
Switch runs the first case that satisfies condition

switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.\n", os)
}

Switch without a condition is the same as switch true. 

os := runtime.GOOS
switch {
case os == "darwin":
	fmt.Println("OS X.")
case os =="linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.\n", os)
}


### Defer
---
A defer statement defers the execution of a function until others are done and stacks

fmt.Println("counting")

for i := 0; i < 10; i++ {
	defer fmt.Println(i)
}

fmt.Println("done")

### Methods
---
type Vertex struct {

	X, Y float64
	
}

v = OBJECT VARIABLE Vertex = TYPE METHOD NAME

func (v Vertex) Abs() float64 {							func Abs(v Vertex) float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)						return math.Sqrt(v.X*v.X + v.Y*v.Y)
	
}										}

func main() {

	v := Vertex{3, 4}
	
	fmt.Println(Abs(v))
	
}

WITH NON-STRUCT TYPE

type MyFloat float64

func (f MyFloat) Abs() float64 {

	if f < 0 {
	
		return float64(-f)
		
	}
	
	return float64(f)
	
}

### Pointer
---
type Vertex struct {

	X, Y float64
	
}

func (v *Vertex) Scale(f float64) {

	v.X = v.X * f
	
	v.Y = v.Y * f
	
}

### Interface
---
type Abser interface {

	Abs() float64
	
}

### Error Handling
---
f = i.(float64) // ERROR

fmt.Println(f)

f, ok := i.(float64) // OK

fmt.Println(f, ok)


RESEARCH
INTERFACE
STRINGER/ERROR STATE
