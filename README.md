# Go
***
### TIPS
GO IS A STATIC TYPED LANGUAGE - TYPE OF A VARIABLE CANNOT CHANGE

### Resources
---
sentdex - https://www.youtube.com/playlist?list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX - 17

Tech With Tim - https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q - 4

### Research
---
INTERFACE

STRINGER/ERROR STATE

p  = &Vertex{1, 2}

### QUOTES
---
"" REPRESENTS A COLLECTION OF CHARACTERS (STRING)

'' REPRESENTS A CHARACTER

`` MULTILINE STRING

USE DOUBLE QUOTES

## COMPILE/RUN
---
TO COMILE AND RUN

go run script.go

TO JUST COMPILE. THE COMPILED FILE CAN BE RUN ON ANY OS.

go build script.go

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

TO IMPORT SUBPACKAGES

import "math/rand"

TO USE rand.Intn(100)

### Exported Names
---
A NAME IS EXPORTED IF IT STARTS WITH A CAPITAL NAME.

ALL THE FUNCTIONS YOU ARE USING FROM A PACKAGE START WITH CAPITAL LETTERS.

math.Pi, fmt.Println ()

### COMMENTS
---
// FOR SINGLE LINE
/* */ FOR MULTIPLE LINES

### Variable
---
VARIABLES WITHOUT INITALIZERS

var c, python, java bool

VARIABLES WITH INITIALIZERS

var i, j int = 1, 2

var (

	i int = 1
	
	j int = 2
	
)

### EXPLRESSION ASSIGNMENT OPERATOR/WALRUS OPERATOR
SHORT VARIABLE DECLARATION (MUST BE WITHIN A FUNCTION)

k := 3

### IMPLICIT/EXPLICIT

IMPLICIT - IMPLICITLY ASK GO TO DEFINE WHAT THIS VARIABLE SHOULD BE

var number = 260

EXPLICIT - EXPLICITLY DEFINED THE TYPE OF VARIABLE

var number uint16 = 260

IT IS POSSIBLE THAT WHAT GO DIFINED IS NOT WHAT YOU NEED

### Constant
---
Contstant variables do not change

const Pi = 3.14

### Function
---
MAIN FUNCTION WILL ALWAYS RUN

func main () {

	fmt.Println ("HELLO WORLD")
	
}

SINGLE INTEGER RETURN

func add (x, y int) int {
	return x + y, x
}

MULTIPULE RETURN

func add (x int, y int) (int, int) {

	return x + y, x
	
}


NAMED RETURN

func split(sum int) (x, y int) {

	x = sum * 4 / 9
	
	y = sum - x
	
	return
}

### Print
---
fmt.Println () - PRINT LINE

fmt.Sprint () - STORE FORMATTED STRING TO VARIABLE

fmt.Printf () - PRINT FORMATTED STRING

### FORMATTING
GENERAL

%v - VALUE

%T - TYPE

%% - PERCENT SIGN

BOOLEAN

%t - BOOLEARN

INTEGER

%b - BASE 2 (BINARY REPRESENTATION)

%o - BASE 8 (OCTAL DECIMAL)

%d - BASE 10 (DECIMAL)

%x - BASE 16 LOWER (HEXADECIMAL)

%X - BASE 16 UPPER (HEXADECIMAL)

FLOATING POINTS

%e - SCIENTIFIC NOTATION OF e

%f / %F - DECIMAL TO EXPONENT

%g - LARGET DECIMAL TO EXPONENT

STRINGS

%s - STRING

%q - QUOTED STRING

WIDTH & PRECISION

%f - DEFAULT WIDTH (PADDING LEFT, USE - FOR RIGHT), DEFAULT PRECISION (FLOATING POINTS)

%9f - WIDTH 9, DEFAULT PRECISION

%.2f - DEFAULT WIDTH, PRECISION 2

%9.2f - WIDTH 9, PRECISION 2

%9.F - WIDTH 9, PRECISION 0

PADDING

%09D - PAD DIGITS TO LENGTH 9 WITH PRECEEDING 0'S

%-4D - PAD WITH SPACES (WIDTH 4, LEFT JUSTIFIED)

NEW LINE

\n

TAB

\t

### Types
---
bool

string

int  int8  int16  int32  int64

uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32, represents a Unicode code point

float32 float64

complex64 complex128

### For loop / While loop
---
INIT; CONDITION; POST

for i := 0; i < 10; i++ {

	fmt.Println (i)
	
}

INIT, CONDITION, POST ARE OPTIONAL (INIFINITE LOOP)

sum := 1

for sum < 1000 {

	sum += sum
	
}

x := 5

for {

	x += 3
	
	if x > 5 {
	
		break
	}
}

RANGE TO ITERATE SLICE OR MAP (YOU CAN USE _ AS VARIABLE TO UNASSIGN)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

for i, v := range pow {

	fmt.Printf("2**%d = %d\n", i, v)
	
}

### If
---
v = 5

if v < 10 {

}

IF CAN START WITH VARIABLE DECLARATION AND THIS VARIABLE IS ALSO AVAILABLE IN ELSE BLOCKS

if v := 5; v < 10 {

}

else {

	v += 5
	
}

### Switch
---
SWITCH RUNS THE FIRST CAST THAT SATISFIES CONDITION

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

SWITCH WITHOUT A CONDITION IS SAME AS SWITCH TRUE

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

### Pointer
---
i, j := 42, 2701

& IS POINT TO MEMORY ADDRESS. * IS READ THROUGH MEMORY ADDRESS

p := &i         // point to i

fmt.Println(*p) // read i through the pointer

*p = 21         // set i through the pointer

fmt.Println(i)  // see the new value of i

### Defer
---
A DEFER STATEMENT DEFERS THE EXECUTION OF A FUNCTION UNTIL OTHERS ARE DONE AND STACKS

fmt.Println("counting")

for i := 0; i < 10; i++ {

	defer fmt.Println(i)
	
}

fmt.Println("done")

### Struct
---
type Vertex struct {

	X int
	
	Y int
	
}

func main() {

	v := Vertex {1, 2}
	v := Vertex {X:1, Y: 2}
	
	v.X = 4
	
	fmt.Println(v.X)
}

### Array
---
ARRAYS CANNOT BE RESIZED

var a [2]string

a[0] = "Hello"

a[1] = "World"
	
primes := [6]int{2, 3, 5, 7, 11, 13}

### Slice
---
SLICE IS DYNAMICALLY SIZED
STARTING INDEX WILL COMPLETELY DROP PREVIOUS ELEMENTS

q := []int{2, 3, 5, 7, 11, 13}

r := []bool{true, false, true, true, false, true}

s := []struct {

	i int
	
	b bool
	
}{

	{2, true},
	
	{3, false}
}

SLICE IS LIKE REFERENCES TO ARRAYS (IF YOU CHANGE IT, ORIGINAL ARRAY/SLICE CHANGES)

names := [2]string{

	"John",
	
	"Paul",
	
}

b := names[1:3]

b[0] = "XXX"

SLICE HAS LENGTH AND CAPACITY

s := []int{2, 3, 5, 7, 11, 13}

printSlice(s)

// Slice the slice to give it zero length.

s = s[:0]

printSlice(s)

// Extend its length.

s = s[:4]

printSlice(s)

// Drop its first two values.

s = s[2:]

printSlice(s)

s = s[0:4]

printSlice(s)

func printSlice(s []int) {

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	
}

ZERO VALUE OF SLICE IS NIL

var s []int

if s == nil {

	fmt.Println("nil!")
	
}

MAKE CAN BE USED TO SET LENGTH AND CAPACITY

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

SLICE CAN HAVE SLICE

board := [][]string{

	[]string{"_", "_", "_"},

	[]string{"_", "_", "_"},

	[]string{"_", "_", "_"},

}

board[0][0] = "X"

APPEND

var s []int

s = append(s, 0)

### HASHTABLE/MAP

type Vertex struct {

	Lat, Long float64
	
}

#YOU NEED TO USE MAKE IF YOU ARE ONLY INITAILIZING AND THEN ASSIGN KEY VALUE LATER

m := make(map[string]Vertex)

m["Bell Labs"] = Vertex{

	40.68433, -74.39967,

}

OR

var m = map[string]Vertex{

	"Bell Labs": Vertex{
	
		40.68433, -74.39967,

	},
	"Google": Vertex{
	
		37.42202, -122.08408,

	},
}

OR

var m = map[string]Vertex{

	"Bell Labs": {40.68433, -74.39967},
	
	"Google":    {37.42202, -122.08408},
	
}

INSERT/RETRIEVE/DELETE

m[key] = elem / elem = m[key] / delete(m, key)

TEST IF KEY EXIST (ok = TRUE IF EXIST OTHERWISE FALSE AND v = 0)

elem, ok = m[key]

ITERATE

for key, value := range m:
	fmt.Println (key, value)

### FUNCTION VS METHOD
---
FUNCTION CAN BE USED ANYWHERE WHEREAS METHODS ARE ASSOCIATED WITH TYPE

IT IS POSSIBLE TO USE FUNCTION TO RECEIVE/MODIFY BY TAKING TYPE AS INPUT

### VALUE RECEIVE METHOD
---
VALUE RECEIVER ONLY RECEIVES VALUES AND THAT'S ALL THEY CAN DO. THEY CANNOT CHANGE VALUES OF THE TYPE

VALUCE RECEIVER CAN CHANGE VALUES BUT IT WILL ONLY BE WITHIN SCOPE

VALUE RECEIVER WORKS LIKE MAKING A COPY OF THE OBJECT.

IF THE TYPE IS HUGE, IT IS MORE EFFICIENT TO USE POINTER RECEIVER. WHEREAS IF THE OBJECT IS SMALL (ATTRIBUTES), IT CAN BE MORE EFFICIENT TO USE VALUE RECEIVER. 

type Vertex struct {

	X, Y float64
	
}

v = OBJECT VARIABLE Vertex = TYPE METHOD NAME

func (v Vertex) Abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
	
}

func main() {

	v := Vertex{3, 4}
	
	fmt.Println (v.Abs ())
	
}

WITH NON-STRUCT TYPE

type MyFloat float64

func (f MyFloat) Abs() float64 {

	if f < 0 {
	
		return float64(-f)
		
	}
	
	return float64(f)
	
}

### POINTER RECEIVER METHOD
---
POINTER RECEIVER CAN MODIFY VALUES/ATTRIBUTES OF THE TYPE

IF THE TYPE IS HUGE, IT IS MORE EFFICIENT TO USE POINTER RECEIVER. WHEREAS IF THE OBJECT IS SMALL (ATTRIBUTES), IT CAN BE MORE EFFICIENT TO USE VALUE RECEIVER.

type Vertex struct {

	X, Y float64
	
}

THIS METHOD IS NOT RETURNING ANYTHING SO RETURN TYPE IS EMPTY
func (v *Vertex) Scale(f float64) {

	v.X = v.X * f
	
	v.Y = v.Y * f
	
}

### FUNCTION AS VALUE
---
FUNCTIONS CAN BE PASSED AROUND LIKE VALUES

func compute(fn func(float64, float64) float64) float64 {

	return fn(3, 4)
	
}

func main() {

	hypot := func(x, y float64) float64 {
	
		return math.Sqrt(x*x + y*y)
		
	}
	
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	
	fmt.Println(compute(math.Pow))
}

### FUNCTION CLOSURE
---
A CLOSURE IS A FUNCTION VALUE THAT REFERENCES VARIABLE FROM OUTSIDE ITS BODY

func adder() func(int) int {

	sum := 0
	
	return func(x int) int {
	
		sum += x
		
		return sum
		
	}
	
}

func main() {

	pos, neg := adder(), adder()
	
	for i := 0; i < 10; i++ {
	
		fmt.Println(
		
			pos(i),
			
			neg(-2*i),
			
		)
		
	}
	
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
