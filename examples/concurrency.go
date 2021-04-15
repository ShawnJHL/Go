package main

import (
  "fmt"
  "time"
  "sync"
)

var wg sync.WaitGroup

func cleanup () {
  if r := recover (); r != nil {
    fmt.Println ("RECOVERED IN CLEANUP", r)
  }
}

func say (s string) {
  defer wg.Done ()
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
  wg.Add (1)
  go say ("HEY")
  wg.Add (1)
  go say ("THERE")
  
  wg.Wait ()
  go say ("HI")
  wg.Wait ()
}
