package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	switch booru {
	case "konachan":
		url = "http://konachan.com/post.xml?"
		params = fmt.Sprintf("tags=%s&rating=%s&page=%d&limit=%d", tags, rating, page, count)
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
