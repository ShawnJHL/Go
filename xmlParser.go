package main

import ("fmt
        "net/http"
        "io/ioutil
        "encoding/xml")

/*
EXAMPLE 1
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

# SitemapIndex, Location, AND STRING TYPES CAN BE COMBINED INTO ONE STRUCT TO SAVE SPACE
# SITEMAP>LOC MEANS LOC TAG UNDER SITEMAP TAG
type SitemapIndex stsruct {
        # MUST USE CAPITALIZATION TO MAKE AS INTERNAL FUNCTION
        Locations []string `xml:"sitemap>loc"`
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

/*
EXAMPLE 2
https://www.washingtonpost.com/news-sitemap-index.xml
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
	xmlns:n="http://www.google.com/schemas/sitemap-news/0.9" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
       http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd
       http://www.google.com/schemas/sitemap-news/0.9
       http://www.google.com/schemas/sitemap-news/0.9/sitemap-news.xsd">
	<url>
			<loc>https://www.washingtonpost.com/business/technology/un-adds-32-items-to-list-of-prohibited-goods-for-north-korea/2017/10/23/5f112818-b812-11e7-9b93-b97043e57a22_story.html</loc>
			<changefreq>hourly</changefreq>
			<n:news>
				<n:publication>
					<n:name>Washington Post</n:name>
					<n:language>en</n:language>
				</n:publication>
				<n:publication_date>2017-10-23T22:12:20Z</n:publication_date>
				<n:title>UN adds 32 items to list of prohibited goods for North Korea</n:title>
				<n:keywords>
					UN-United Nations-North Korea-Sanctions,North Korea,East Asia,Asia,United Nations Security Council,United Nations,Business,General news,Sanctions and embargoes,Foreign policy,International relations,Government and politics,Government policy,Military technology,Technology</n:keywords>
			</n:news>
		</url>
	<url>
			<loc>https://www.washingtonpost.com/business/technology/cisco-systems-buying-broadsoft-for-19-billion-cash/2017/10/23/ae024774-b7f2-11e7-9b93-b97043e57a22_story.html</loc>
			<changefreq>hourly</changefreq>
			<n:news>
				<n:publication>
					<n:name>Washington Post</n:name>
					<n:language>en</n:language>
				</n:publication>
				<n:publication_date>2017-10-23T21:42:14Z</n:publication_date>
				<n:title>Cisco Systems buying BroadSoft for $1.9 billion cash</n:title>
				<n:keywords>
					US-Cisco-BroadSoft-Acquisition,Cisco Systems Inc,Business,Technology,Communication technology</n:keywords>
			</n:news>
		</url>
	</urlset>
*/

type News struct {
        Titles []string `xml:"url>news>title"`
        Keywords []string `xml:"url>news>keywords"`
        Locations []string `xml:"url>loc"
}

func main() {
	var s Sitemapindex
        var n News
        
        resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
        	bytes, _ := ioutil.ReadAll(resp.Body)
                xml.Unmarshal(bytes, &n)
	}
}
