# Go
***
### TIPS
GO IS A STATIC TYPED LANGUAGE - TYPE OF A VARIABLE CANNOT CHANGE (MOST OF THE TIME)

### Resources
---
sentdex - https://www.youtube.com/playlist?list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX - 17

Tech With Tim - https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q - 16

### RESEARCH
---
INTERFACE

STRINGER/ERROR STATE

p  = &Vertex{1, 2}

## COMPILE/RUN
---
TO COMILE AND RUN

go run script.go

TO JUST COMPILE. THE COMPILED FILE CAN BE RUN ON ANY OS.

go build script.go

### PACKAGE
---
EVERY GO PROGRAM IS MADE UP OF PACKAGES

PROGRAMS START RUNNING IN PACKAGE MAIN

### IMPORT
---
import "fmt"

-OR-

import (

	"fmt"
	
	"math"
	
	)

TO IMPORT SUBPACKAGES

import "math/rand"

TO USE

rand.Intn(100)

### MAIN FUNCTION
---
MAIN FUNCTION WILL ALWAYS RUN

func main() {
}

### EXPORTED NAME
---
A NAME IS EXPORTED IF IT STARTS WITH A CAPITAL NAME.

ALL THE FUNCTIONS FROM A PACKAGE START WITH CAPITAL LETTERS.

math.Pi, fmt.Println (), ETC

### COMMENTS
---
// FOR SINGLE LINE
/* */ FOR MULTIPLE LINES

### VARIABLE
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
---
SHORT VARIABLE DECLARATION (MUST BE WITHIN A FUNCTION)

k := 3

### IMPLICIT/EXPLICIT

IMPLICIT - IMPLICITLY ASK GO TO DEFINE WHAT THIS VARIABLE SHOULD BE

var number = 260

EXPLICIT - EXPLICITLY DEFINED THE TYPE OF VARIABLE

var number uint16 = 260

IT IS POSSIBLE THAT WHAT GO DIFINED IS NOT WHAT YOU NEED

### TYPE
---
bool

string

int  int8  int16  int32  int64

uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32, represents a Unicode code point

float32 float64

complex64 complex128

### CONSTANT
---
CONSTANT VARIABLES DO NOT CHANGE

const Pi = 3.14

### QUOTES
---
"" REPRESENTS A STRING (USER DOUBLE QUOTES)

'' REPRESENTS A CHARACTER

\`\` MULTILINE STRING

### PRINT
---
fmt.Println () - PRINT LINE

fmt.Sprint () - STORE FORMATTED STRING TO VARIABLE

fmt.Printf () - PRINT FORMATTED STRING OR WITHOUT NEWLINE

### FORMATTING
---
GENERAL

%v - VALUE

%T - TYPE

%% - PERCENT SIGN

BOOLEAN

%t - BOOLEARN

INTEGER

%b - BASE 2 (BINARY REPRESENTATION)

%o - BASE 8 (OCTAL DECIMAL)

%d - BASE 10 (BASE, DECIMAL)

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

### CONDITION & BOOLEAN EXPRESSION
---
COMPARING VARIABLES MUST BE THE SAME TYPE

x := 5

y := 5.5

z = x == int (y)

### CONDITIONAL STATEMENT (IF)
---
v = 5

if v < 10 {

}

IF CAN START WITH A DECLARATION AND THIS VARIABLE IS ALSO AVAILABLE IN ELSE BLOCKS

if v := 5; v < 10 {

} elses if v == 10 {

} else {
	
}

### SWITCH
---
SWITCH RUNS THE FIRST CAST THAT SATISFIES CONDITION

switch os := runtime.GOOS; os {

case "darwin", "another match":

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

### FOR LOOP / WHILE LOOP
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

### RANGE
---

ITERATE ARRAY, SLICE, AND MAP (YOU CAN USE _ AS VARIABLE TO UNASSIGN)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

for i, v := range pow {

	fmt.Printf("2**%d = %d\n", i, v)
	
}

### ARRAY
---
ARRAY CANNOT BE RESIZED

ARRAY/SLICE CAN ONLY STORE ONE TYPE OF DATA (STRING, INT, ETC)

var a [2]string

a[0] = "Hello"

a[1] = "World"

-OR-
	
primes := [6]int{2, 3, 5, 7, 11, 13}

-OR-

MULTIDIMENTIONAL ARRAY

arr2D := [2][2]int{{1,2},{3,4}}

### Slice
---
SLICE CAN BE DYNAMICALLY SIZED

ARRAY/SLICE CAN ONLY STORE ONE TYPE OF DATA (STRING, INT, ETC)

q := []int{2, 3, 5, 7, 11, 13}

-OR-

r := []bool{true, false, true, true, false, true}

-OR-

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

SLICE OPERATOR :, IF START INDEX IS NOT SET DEFAULT IS 0 IF END INDEX IS NOT SET ITS SIZE OF THE ARRAY

b := names[:]

b[0] = "XXX" (JOHN BECOMES XXX)

STARTING INDEX OF SLICE WILL COMPLETELY DROP PREVIOUS ELEMENTS (START INDEX IS NOT ADJUSTABLE AND BECOMES BASE WHEREAS END INDEX IS ADJUSTABLE)

s := []int{2, 3, 5, 7, 11, 13}

s = s[:0] (LENGTH 0, CAPACITY (6)

s = s[:4] (LENGTH 4, CAPACITY 6)

s = s[2:] (LENGTH 2, CAPACITY 4, DROPS THE FIRST TWO ELEMENTS)

s = s[0:4] (LENGTH 4, CAPACITY 4)

MULTIDIMENTIONAL SLICE

board := [][]string{

	[]string{"_", "_", "_"},

	[]string{"_", "_", "_"},

	[]string{"_", "_", "_"},

}

board[0][0] = "X"

MAKE CAN BE USED TO SET LENGTH AND CAPACITY

b := make([]int, 0, 5) (LENGTH 0, CAPACITY 5)

### ARRAY/SLICE FUNCTION
---
len (array) - LENGTH

cap (array) - CAPACITY

append (array, element) - APPEND

### HASHTABLE/MAP
---
type Vertex struct {

	Lat, Long float64
	
}

YOU NEED TO USE MAKE IF YOU ARE ONLY INITAILIZING AND THEN ASSIGN KEY VALUE LATER

m := make(map[string]Vertex)

m["Bell Labs"] = Vertex{

	40.68433, -74.39967,

}

-OR-

var m = map[string]Vertex{

	"Bell Labs": Vertex{
	
		40.68433, -74.39967,

	},
	"Google": Vertex{
	
		37.42202, -122.08408,

	},
}

-OR-

var m = map[string]Vertex{

	"Bell Labs": {40.68433, -74.39967},
	
	"Google":    {37.42202, -122.08408},
	
}

### MAP/HASHTABLE FUNCTION
---
map[key] = value - INSERT/UPDATE

map[key] - RETRIEVE

delete (map, key) - DELETE

value, ok = m[key] - CHECK IF A KEY EXISTS (ok = TRUE IF EXIST OTHERWISE FALSE, VALUE WILL BE DEFAULT TO 0 IF NOT EXIST)

len (map) - SIZE OF MAP

### FUNCTION
---
SINGLE RETURN

func add (x int, y int) int {
	return x + y, x
}

MULTIPULE RETURNS

func add (x, y int) (int, int) {

	return x + y, x
	
}

NAMED RETURN

func split(sum int) (x, y int) {

	x = sum * 4 / 9
	
	y = sum - x
	
	return
}

### DEFER
---
A DEFER STATEMENT DEFERS THE EXECUTION OF A FUNCTION UNTIL OTHERS ARE DONE/PANIC AND STACKS

fmt.Println("counting")

for i := 0; i < 10; i++ {

	defer fmt.Println(i)
	
}

fmt.Println("done")

### Pointer
---
i, j := 42, 2701

& IS POINT TO MEMORY ADDRESS. * IS READ THROUGH MEMORY ADDRESS

p := &i         // point to i

fmt.Println(*p) // read i through the pointer

*p = 21         // set i through the pointer

fmt.Println(i)  // see the new value of i

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

### CONCURRENCY / GOR OUTINES
---
CONCURRENCY (go) IS LIKE DEALING WITH MULTIPLE THINGS (THREADS) AT ANY ONE TIME

func say(s string) {

	for i:=0; i < 3; i++ {
	
		time.Sleep(100*time.Millisecond)
		
		fmt.Println(s)
		
	}	
	
}

func main() {

	go say("Hey")
	
	say("there")
	
}

### SYNCRONIZATION
---

SYNCRONIZATION IS USED TO MARK EACH THREAD AS START AND COMPLETE SO OS CAN WAIT FOR THEM TO FINISH

var wg sync.WaitGroup

func cleanup () {

  if r := recover (); r != nil {
  
    fmt.Println ("RECOVERED IN CLEANUP", r)
    
  }
  
}

func say (s string) {

MARK DONE
  defer wg.Done ()

CHECK FOR PANIC AND RECOVER FROM IT
  defer cleanup ()
  
  for i:= 0; i < 3; i++ {
  
    fmt.Println (s)
    
    time.Sleep (time.Millisecond * 100)
    
    if i == 2 {
    
      panic ("OH DEAR 2")
      
    }
    
  }
  
}

func main () {

MARK START
  wg.Add (1)
  
  go say ("HEY")
  
MARK START
  wg.Add (1)
  
  go say ("THERE")
  
WAIT FOR ALL THREADS TO BE COMPLETED
  wg.Wait ()
  
  go say ("HI")
  
  wg.Wait ()
  
}

### CHANNEL
---
CHANNEL <- IN GO IS TO USE THEM WITH GO ROUTINES TO SEND AND RECEIVE VALUES BETWEEN THEM

SEND AND RECEIVE PART OF CHANNEL IS GOING TO BLOCK (WITHOUTIT PROGRAM CAN FINISH BEFORE THE CONCURRENCY IS DONE) GO ROUTINES SO YOU DON'T HAVE TO WAIT FOR THEM

var wg sync.WaitGroup

func foo(c chan int, someValue int) {

    defer wg.Done()
    
    c <- someValue * 5
    
}

func main() {

    TYPE, BUFFER SIZE
    
    fooVal := make(chan int, 10)
    
    for i := 0; i < 10; i++ {
    
        wg.Add(1)
	
        go foo(fooVal, i)
	
    }
    
    wg.Wait()
    
    close(fooVal)
    
    for item := range fooVal {
    
        fmt.Println(item)
	
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
