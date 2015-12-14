package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type XMLPosts struct {
	XMLName xml.Name  `xml:"posts"`
	Posts   []XMLPost `xml:"post"`
}
type XMLPost struct {
	XMLName xml.Name `xml:"post"`
	fileURL string   `xml:"file_url,attr"`
	id      int      `xml:"id,attr"`
	md5     string   `xml:"md5,attr"`
}

func ReadPosts(reader io.Reader) ([]XMLPost, error) {
	var xmlPosts XMLPosts
	if err := xml.NewDecoder(reader).Decode(&xmlPosts); err != nil {
		return nil, err
	}
	return xmlPosts.Posts, nil
}
func request(booru string, tags string, rating string, page int, count int) []XMLPost {
	var url string
	var params string
	switch booru {
	case "konachan":
		url = "http://konachan.com/post.xml?"
		params = fmt.Sprintf("tags=%s&rating=%s&page=%d&count=%d", tags, rating, page, count)
	}
	resp, err := http.Get(url + params)
	if err != nil {
		log.Fatalf("Failed to request posts, ERROR: %s", err.Error())
	}
	defer resp.Body.Close()
	p, err := ReadPosts(resp.Body)
	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response, ERROR: %s", err.Error())
	}

	return p
}
