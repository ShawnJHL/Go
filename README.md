# Go
======


### Package
------
Every go program is made up of packages.
Programs start running in package main.


### Import
------
import "fmt"
import "math"


import (
	"fmt"
	"math"
)


### Exported Names
------
A name exported if it starts with a capital name like math.Pi


### Function
------
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
------
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
------
fmt.Println ()
fmt.Printf ()
fmt.Sprint ()
%b = binary
%t = boolean
%v = value
%T = type


### Types
------
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
------
Contstant variables do not change
const Pi = 3.14
