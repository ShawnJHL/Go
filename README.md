# Go
***

### TIPS
---
STATIC TYPE - TYPE OF A VARIABLE CANNOT CHANGE (MOST OF THE TIME)

### Resources
---
sentdex - https://www.youtube.com/playlist?list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX  
Tech With Tim - https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q  

### COMPILE/RUN
---
TO COMILE AND RUN
> go run script.go  

TO COMPILE ONLY. THE COMPILED FILE CAN BE RUN ON ANY OS
> go build script.go  

### PACKAGE
---
EVERY GO PROGRAM IS MADE UP OF PACKAGES
```Go
package main
```

### IMPORT
---
```Go
import "fmt"

-OR-

import (

	"fmt"
	
	"math"
	
	)
```

TO IMPORT SUBPACKAGES
```Go
import "math/rand"

rand.Intn(100)
```

### MAIN FUNCTION
---
MAIN FUNCTION WILL ALWAYS RUN
```Go
func main() {}
```

### EXPORTED NAME
---
A NAME IS EXPORTED IF IT STARTS WITH A CAPITAL NAME
```Go
math.Pi
fmt.Println
ETC
```

### COMMENTS
---
SINGLE LINE
```Go
//
```

MULTIPLE LINES
```Go
/* */
```

### VARIABLE
---
VARIABLES WITHOUT INITALIZERS
```Go
var c, python, java bool
```

VARIABLES WITH INITIALIZERS
```Go
var i, j int = 1, 2
```
-OR-
```Go
var (
	i int = 1
	
	j int = 2
)
```

### EXPLRESSION ASSIGNMENT OPERATOR/WALRUS OPERATOR
---
SHORT VARIABLE DECLARATION (MUST BE WITHIN A FUNCTION)
```Go
k := 3
```

### IMPLICIT/EXPLICIT
---
IMPLICIT - IMPLICITLY ASK GO TO DEFINE WHAT THIS VARIABLE SHOULD BE
```Go
var number = 260
```
EXPLICIT - EXPLICITLY DEFINED THE TYPE OF VARIABLE
```Go
var number uint16 = 260
```
IT IS POSSIBLE THAT WHAT GO DIFINED IS NOT WHAT YOU NEED

### TYPE
---
bool  
string  
int
- int
- int8
- int16
- int32
- int64

uint  
- uint
- uint8
- uint16
- uint32
- uint64
- uintptr

byte - ALIAS FOR UINT8  
rune - ALIAS FOR INT32, REPRESENTS A UNICODE POINT  
float
- float32
- float64

complex
- complex64
- complex128

func

### CONSTANT
---
```Go
const Pi = 3.14
```

### QUOTES
---
REPRESENTS A STRING (USER DOUBLE QUOTES)
```Go
""
```

REPRESENTS A CHARACTER
```Go
'' 
```

MULTILINE STRING
```Go
``
```

### PRINT
---
PRINT LINE
```Go
fmt.Println ()
```

PRINT FORMATTED STRING OR WITHOUT NEWLINE
```Go
fmt.Printf ()
```

STORE FORMATTED STRING TO VARIABLE
```Go
fmt.Sprint ()
```

### FORMATTING
---
GENERAL  
- %v - VALUE
- %T - TYPE  
- %% - PERCENT SIGN  

BOOLEAN  
- %t - BOOLEARN

INTEGER  
- %b - BASE 2 (BINARY REPRESENTATION)
- %o - BASE 8 (OCTAL DECIMAL)
- %d - BASE 10 (BASE, DECIMAL)
- %x - BASE 16 LOWER (HEXADECIMAL)
- %X - BASE 16 UPPER (HEXADECIMAL)

FLOATING POINTS  
- %e - SCIENTIFIC NOTATION OF e
- %f / %F - DECIMAL TO EXPONENT
- %g - LARGET DECIMAL TO EXPONENT

STRINGS
- %s - STRING
- %q - QUOTED STRING

WIDTH & PRECISION
- %f - DEFAULT WIDTH (PADDING LEFT, USE - FOR RIGHT), DEFAULT PRECISION (FLOATING POINTS)
- %9f - WIDTH 9, DEFAULT PRECISION
- %.2f - DEFAULT WIDTH, PRECISION 2
- %9.2f - WIDTH 9, PRECISION 2
- %9.F - WIDTH 9, PRECISION 0

PADDING
- %09D - PAD DIGITS TO LENGTH 9 WITH PRECEEDING 0'S
- %-4D - PAD WITH SPACES (WIDTH 4, LEFT JUSTIFIED)

NEW LINE
- \n

TAB
- \t

### COMPARISON
---
ALL COMPARISON MUST BE THE SAME TYPE
```Go
x := 5
y := 5.5

z = x == int (y)
```

### CONDITIONAL STATEMENT (IF)
---
```Go
v = 5

if v < 10 {
}
```

IF CAN START WITH A DECLARATION AND THIS VARIABLE EXISTS WITHIN THE IF/ELSEIF/ELSE SCOPE
```Go
if v := 5; v < 10 {
} else if v == 10 {
} else {
}
```

### SWITCH
---
SWITCH RUNS THE FIRST CAST THAT SATISFIES CONDITION
```Go
switch os := runtime.GOOS; os {
	case "darwin", "another match":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
}
```

SWITCH WITHOUT A CONDITION IS SAME AS SWITCH TRUE
```Go
os := runtime.GOOS

switch {
	case os == "darwin":
		fmt.Println("OS X.")
	case os =="linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
}
```

### FOR LOOP / WHILE LOOP
---
INIT; CONDITION; POST
```Go
for i := 0; i < 10; i++ {
	fmt.Println (i)
}
```

WHILE LOOP
```Go
sum := 1

for sum < 1000 {
	sum += sum
}
```
-OR-
```Go
x := 5

for {
	x += 3
	
	if x > 5 {
		break
	}
}
```

### RANGE
---
ITERATE ARRAY, SLICE, AND MAP (YOU CAN USE _ AS VARIABLE TO UNASSIGN)
```Go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

for i, v := range pow {
	fmt.Printf("2**%d = %d\n", i, v)
}
```

### ARRAY
---
ARRAY CANNOT BE RESIZED AND CAN ONLY STORE ONE TYPE OF DATA
```Go
var a [2]string

a[0] = "Hello"
a[1] = "World"
```
-OR-
```Go
primes := [6]int{2, 3, 5, 7, 11, 13}
```
-OR-  
MULTIDIMENTIONAL ARRAY
```Go
arr2D := [2][2]int{{1,2},{3,4}}
```

### Slice
---
SLICE CAN BE DYNAMICALLY SIZED AND CAN ONLY STORE ONE TYPE OF DATA
```Go
q := []int{2, 3, 5, 7, 11, 13}
```
-OR-
```Go
r := []bool{true, false, true, true, false, true}
```
-OR-
```Go
s := []struct {
	i int
	b bool
	
}{
	{2, true},
	{3, false}
}
```

SLICE IS LIKE REFERENCES TO ARRAYS (IF YOU CHANGE IT, ORIGINAL ARRAY/SLICE CHANGES)
```Go
names := [2]string{
	"John",
	"Paul",
}
```

SLICE OPERATOR : IF START INDEX IS NOT SET DEFAULT IS 0 IF END INDEX IS NOT SET ITS SIZE OF THE ARRAY
```Go
b := names[:]

b[0] = "XXX" (JOHN BECOMES XXX)
```

STARTING INDEX OF SLICE WILL COMPLETELY DROP PREVIOUS ELEMENTS (START INDEX IS NOT ADJUSTABLE AND BECOMES BASE WHEREAS END INDEX IS ADJUSTABLE)
```Go
s := []int{2, 3, 5, 7, 11, 13}
s = s[:0] (LENGTH 0, CAPACITY (6)
s = s[:4] (LENGTH 4, CAPACITY 6)
s = s[2:] (LENGTH 2, CAPACITY 4, DROPS THE FIRST TWO ELEMENTS)
s = s[0:4] (LENGTH 4, CAPACITY 4)
```

MULTIDIMENTIONAL SLICE
```Go
board := [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}

board[0][0] = "X"
```

MAKE CAN BE USED TO SET LENGTH AND CAPACITY
```Go
b := make([]int, 0, 5) (LENGTH 0, CAPACITY 5)
```

### ARRAY/SLICE FUNCTION
---
LENGTH
```Go
len (array)
```

CAPACITY
```Go
cap (array)
```

APPEND
```Go
append (array, element)
```

### STRUCT/HASHTABLE/MAP
---
```Go
type Vertex struct {
	Lat, Long float64
}
```

YOU NEED TO USE MAKE IF YOU ARE ONLY INITAILIZING AND THEN ASSIGN KEY VALUE LATER
```Go
m := make(map[string]Vertex)

m["Bell Labs"] = Vertex{
	40.68433, -74.39967,
}
```
-OR-
```Go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```
-OR-
```Go
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```

### MAP/HASHTABLE FUNCTION
---
INSERT/UPDATE
```Go
map[key] = value
```

RETRIEVE
```Go
map[key]
```

DELETE
```Go
delete (map, key)
```

CHECK IF A KEY EXISTS (ok = TRUE IF EXIST OTHERWISE FALSE, VALUE WILL BE DEFAULT TO 0 IF NOT EXIST)
```Go
value, ok = m[key]
```

SIZE OF MAP
```Go
len (map)
```

### DEFER
---
A DEFER STATEMENT DEFERS THE EXECUTION OF A FUNCTION UNTIL OTHERS ARE DONE/PANIC AND STACKS

```Go
fmt.Println("counting")

for i := 0; i < 10; i++ {
	defer fmt.Println(i)
}

fmt.Println("done")
```

### FUNCTION
---
WHEN A FUNCTION IS DEFINED, INPUT AND OUTPUT FOR FUNCTION MUST BE DEFINED ALSO (IF THERE IS ONE)

SINGLE RETURN
```Go
func add (x int, y int) int {
	return x + y, x
}
```

MULTIPULE RETURNS
```Go
func add (x, y int) (int, int) {
	return x + y, x
}
```

NAMED RETURN
```Go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	
	return
}
```

### FUNCTION AS A TYPE/VALUE
---
FUNCTION CAN BE ASSIGNED TO A VARIABLE LIKE OTHER TYPES
```Go
func testF () {
}

test := testF ()
```
-OR-
```Go
test := func testF(){
}
```

EXCUTE THE FUNCTION RIGHT AWAY AND ASSIGN THE RETURN TO A VARIABLE
```Go
testF := func (x int) int {
	return x
} (5)
```

### FUNCTION CLOSURE
---
A NESTED FUNCTION THAT REFERENCES VARIABLES FROM OUTSIDE AND IS BEING RETURNED (EX. sum)  

A FUNCTION CLOSURE IS USED WHEN:  
- A STATE/ENVIRONMENT MUST BE REMEMBERED (EX. sum)  
- CLASS WITH SINGLE METHOD

```Go
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
```

POSSIBLE TO RETURN MAP OF FUNCTIONS

### MUTTABLE & IMMUTABLE
---
MUTABLE - CANNOT CHANGE (COPIES THE CONTENT IN RAM)
- VARIABLE
- ARRAY (FIXED SIZE)

IMMUTABLE - CAN CHANGE (COPIES THE POINTER IN RAM)
- SLICE
- MAP

### POINTER/REFERENCE & DEREFERENCE
---
& - POINT TO MEMORY ADDRESS/REFERENCE/POINTER  
* - READ THROUGH MEMORY ADDRESS/DEREFERENCE  

```Go
i, j := 42, 2701

p := &i

*p = 21
```

-OR-

```Go
toChange := "hello"

var pointer *string = &toChange
```

FUNCTION CAN ALSO RECEIVE POINTER(s) AS INPUT AND DEREFERENCE (SO THE VARIABLE CAN BE CHANGED WITHIN THE FUNCTION)
```Go
func changeValue (str *string) {
	*str = "NEW STRING"
}
```

### STRUCT (CUSTOM STRUCTURE)
---
STRUCT IS A TYPE SO ALWAYS START WITH A CAPITAL LETTER

```Go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex {1, 2}
	
	v := Vertex {X:1}

	v Vertex = Vertex {1, 2}
	v.X = 4
}
```

TO PASS A STRUC TYPE AS POINTER/REFERENCE (SO YOU CAN CHANGE VALUES OF STRUCT)  
YOU DO NOT HAVE TO USE * (DEREFERENCE) ON STRUCT TYPE
```Go
func changeX (pt *Vertex) {
	pt.X = 0
}

v := &Vertex {100, 100}

changeX (v)
```

### EMBEDDED STRUCT
---
STRUCT IN ANOTHER STRUCT

```Go
type Circle struct {
	radius float64
	center *Point
}

c1 := Circle {1.11, &Point (1, 1)

c1.center.x = 5
```

-OR-

NAMING A STRUCT IS NOT NECESSARY AS LONG AS THERE ARE NO DUPLICATE ATTRIBUTE NAMES
```Go
type Circle struct {
	radisu float64
	*Point
}

c1 := Circle {1.11, &Point (1, 1)

c1.x = 5
```

PASSING A STRUCT AS POINTER IS NOT A MUST

### FUNCTION VS METHOD
---
FUNCTION CAN BE USED ANYWHERE
> EX. func getAge () int {}

METHODS ARE ASSOCIATED WITH A CLASS (TYPE). TYPE MUST BE PROVIDED.
> EX. func (s Student) getAge () int{}

### VALUE RECEIVE METHOD VS POINTER RECEIVE METHOD
---
VALUE RECEIVE METHOD CAN ONLY RECEIVE TYPET BUT CANNOT CHANGE THEM OUTSIDE OF ITS SCOPE  

POINTER RECEIVE METHOD CAN MODIFY VALUES/ATTRIBUTES OF THE TYPE (PROVIDE * ON INPUT TYPE TO SPECIFY POINTER)  

(IT MAY BE COMMON TO USE POINTER RECEIVE ALL THE TIME ON METHODS)

### VALUE RECEIVE METHOD
---
```Go
type Student struct {
	name string
	grades []int
	age int
}

func (s Student) getAverageGrade() int {
	sum := 0
	
	for _, v := range s.grades {
		sum += v
	}
	
	return float32 (sum) / float32 (len (s.grades))
}

func main() {
	s1 := Student {"Tim", []int{10, 20 ,30}, 10}
	
	s1.getAverageGrade ()
}
```

### POINTER RECEIVER METHOD
---
```Go
func (s *Student) setAge (age int) {
	s.age = age
}

func main() {
	s1 := Student {"Tim", []int{10, 20 ,30}, 10}
	
	s1.setAge (20)
}
```

### NON-STRUCT TYPE METHOD
---
```Go
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	
	return float64(f)
	
}
```

### CONCURRENCY / GO ROUTINES
---
CONCURRENCY (go) IS LIKE DEALING WITH MULTIPLE THINGS (THREADS) AT ANY ONE TIME

```Go
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
```

### SYNCRONIZATION
---
SYNCRONIZATION IS USED TO MARK EACH THREAD AS START AND COMPLETE SO OS CAN WAIT FOR THEM TO FINISH

```Go
var wg sync.WaitGroup

func cleanup () {
	if r := recover (); r != nil {
  
   	fmt.Println ("RECOVERED IN CLEANUP", r)
    
 	}
}

func say (s string) {
	# MARK DONE
	defer wg.Done ()

	# CHECK FOR PANIC AND RECOVER FROM IT
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
	# MARK START
	wg.Add (1)
  
 	go say ("HEY")
  
	# MARK START
  	wg.Add (1)
  
 	go say ("THERE")
  
	# WAIT FOR ALL THREADS TO BE COMPLETED
  	wg.Wait ()
  
  	go say ("HI")
  
 	wg.Wait ()
}
```

### CHANNEL
---
CHANNEL <- IN GO IS TO USE THEM WITH GO ROUTINES TO SEND AND RECEIVE VALUES BETWEEN THEM

SEND AND RECEIVE PART OF CHANNEL IS GOING TO BLOCK (WITHOUTIT PROGRAM CAN FINISH BEFORE THE CONCURRENCY IS DONE) GO ROUTINES SO YOU DON'T HAVE TO WAIT FOR THEM

```Go
var wg sync.WaitGroup

func foo(c chan int, someValue int) {
    defer wg.Done()
    
    c <- someValue * 5
}

func main() {
    # TYPE, BUFFER SIZE
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
```

### INTERFACE
---
INTERFACE DEFINES SIMILAR BEHAVIORS BETWEEN TYPES/OBJECTES  
ANY TYPE/STRUCT THAT HAS ALL THE SAME STRUCT METHODS (NAME, INPUT) & RETURN TYPES DEFINED IN INTERFACE CAN BE PART OF THE INTERFACE  
HOWEVER, BECAUSE OF THAT YOU CANNOT ACCESS INDIVIDUAL STRUCT'S ATTRIBUTES (WIDTH, HEIGHT, RADIUS), BUT ONLY WHAT'S DEFINED WITHIN THE INTERFACE (AREA)
```Go
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width float64
	height float64
}

# EXAMPLE WITH POINTER AND VALUE RECEIVE METHODS

type (r *rect) area() float64 {
	return r.width* r.height
}

type (c circle) area () float64 {
	return 
}

func getArea (s shape) float64 {
	return s.area()
}

func main () {
	c1 := circle{4.5}
	
	r1 := rect {5, 7}
	
	shapes  := []shape{c1, &r1}
	shapes [0].area ()
	
	# -OR-
	
	getArea (shapes)
}
```
