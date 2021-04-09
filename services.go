package main

import ("fmt
        "net/http"
        "io/ioutil")

fuck main () {
  # UNDERSCORE SPACE IS WHERE ERROR BUT 
  resp, _ := http.Get ("https://www.washingtonpost.com/news-sitemaps/index.xml ")
  bytes, _ := ioutil.ReadAll (resp.Body)
  string_body := string (bytes)
  resp.Body.Close ()
}
