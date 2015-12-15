package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type XMLPost struct {
	FileURL string `xml:"file_url,attr"`
	Id      int    `xml:"id,attr"`
	Md5     string `xml:"md5,attr"`
}

type XMLPosts struct {
	XMLName xml.Name  `xml:"posts"`
	Post    []XMLPost `xml:"post"`
}

func request(booru string, tags string, rating string, page int, count int) XMLPosts {
	var url string
	var params string
	ratings := strings.Split(rating, "+")
	switch booru {
	case "konachan":
		safe, questionable, explicit := false, false, false
		for _, x := range ratings {
			if x == "safe" {
				safe = true
			} else if x == "questionable" {
				questionable = true
			} else if x == "explicit" {
				explicit = true
			}
		}
		if safe && questionable && explicit {
			rating = ""
		} else if safe && questionable && !explicit {
			rating = "%20rating%3aquestionableless"
		} else if safe && !questionable && !explicit {
			rating = "%20rating%3asafe"
		} else if !safe && questionable && !explicit {
			rating = "%20rating%3aquestionable"
		} else if !safe && !questionable && explicit {
			rating = "%20rating%3aexplicit"
		} else if !safe && questionable && explicit {
			rating = "%20rating%3aquestionableplus"
		} else if safe && !questionable && explicit {
			log.Fatal("safe+explicit not a valid combination.")
		}
		url = "http://konachan.com/post.xml?"
		params = fmt.Sprintf("tags=%s%s&page=%d&limit=%d", tags, rating, page, count)
	default:
		log.Fatal("Unsupported booru selected.")
	}
	resp, err := http.Get(url + params)
	if err != nil {
		log.Fatalf("Failed to request posts, ERROR: %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response, ERROR: %s", err.Error())
	}
	p := XMLPosts{}
	err = xml.Unmarshal(body, &p)
	if err != nil {
		log.Fatalf("Failed to decode response. ERROR: %s", err.Error())
	}

	return p
}
