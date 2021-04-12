package main

import ("fmt
        "net/http"
        "io/ioutil
        "encoding/xml")

/*
https://developer.mozilla.org/sitemap.xml
<sitemapindex>
<sitemap>
<loc>
https://developer.mozilla.org/sitemaps/ar/sitemap.xml.gz
</loc>
<lastmod>2021-04-12</lastmod>
</sitemap>
<sitemap>
<loc>
https://developer.mozilla.org/sitemaps/bg/sitemap.xml.gz
</loc>
<lastmod>2021-04-12</lastmod>
</sitemap>
*/

type SitemapIndex stsruct {
        # MUST USE CAPITALIZATION TO MAKE AS INTERNAL FUNCTION
        Locations []Location `xml:"sitemap"`
}

type Location struct {
        Loc string `xml:"loc"`
}

# STRING METHOD. TO CONVERT STRUCT LOCATION TO STRING WHENEVER YOU PRINT. LIKE OVERWRITING STRING METHOD FOR A PARTICULAR STRING
func (l location) String () string {
        return fmt.Sprintf (l.Loc)
}

func main () {
  # UNDERSCORE SPACE IS WHERE ERROR
  resp, _ := http.Get ("https://www.washingtonpost.com/news-sitemaps/index.xml")
  
  # READ CONTENTS
  bytes, _ := ioutil.ReadAll (resp.Body)
  # string_body := string (bytes)
        
  vas s SitemapIndex
  xml.Unmarshal (bytes, &s)
        
  fmt.Println (s.Location)
        
  resp.Body.Close ()
}
