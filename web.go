package main

import (
  "fmt"
  "net/http"
)

func index_handler (w http.ResponseWriter, r *http.Request) {
  // FPRINTF IS FORMAT WHAT YOU SPECIFY
  fmt.Fprintf (w, "<h1>Whoa, Go is neat!</h1>")
  fmt.Fprintf (w, "<p>%s and %s are nice</p>", "something", "<strong>some other thing</strong>")
  
  fmt.Fprintf (w, `<h1>Whoa, Go is multi!</h1>
<p>line add</p>`
}

func about_handler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf (w, "ABOUT PAGE!")
}

func main () {
  // TAKES URL PATH AND FIND CORRESPONDING FUNCTION. / IS THE HOMEPAGE. INDEX_HANDLER IS THE FUNCTION WE WANT TO RUN
  http.HandleFunc ("/", index_handler)
  http.HandleFunc ("/about/", about_handler)
  
  // CREATES A LOCAL SERVER AT PORT 8000
  http.ListenAndServe (":8000", nil)
}
