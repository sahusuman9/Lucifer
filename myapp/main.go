package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cavaliergopher/grab/v3"
)

//               scheme://host/path
const webpage = "https://www.google.com/+"

func main() {
	resp, err := http.Get(webpage)
	if err != nil {
		fmt.Println("Error getting webpage", err)
		return
	}

	/*
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body", err)
			return
		}
		resp.Body.Close()
	*/
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Println("Error creating go query document", err)
		return
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		imgPath, _ := s.Attr("src")
		if strings.HasPrefix(imgPath, "http") {
			downloadImage(imgPath)
		} else {
			if strings.HasPrefix(imgPath, "//") {
				scheme := getWebpageScheme(webpage)
				if scheme != "" {
					downloadImage(scheme + ":" + imgPath)
				}
			} else {
				prefix := getWebpageHostAndScheme(webpage)
				downloadImage(prefix + imgPath)
			}
		}
	})
}

func downloadImage(path string) {
	r, err := grab.Get(".", path)
	if err != nil {
		fmt.Println("Error downloading ", err)
		return
	}
	fmt.Println("Downloaded image ", r.Filename)
}

func getWebpageScheme(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		fmt.Println("Could not parse webpage link", err)
		return ""
	}
	return u.Scheme
}

func getWebpageHostAndScheme(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		fmt.Println("Error parsing url", err)
		return ""
	}
	return u.Scheme + "://" + u.Host
}
