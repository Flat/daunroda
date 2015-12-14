package main

import (
	"github.com/mitchellh/ioprogress"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func download(url string, md5String string, outDir string) {
	findex := len(strings.Split(url, "/")) - 1
	filename := strings.Split(url, "/")[findex]
	out, err := os.Create(outDir + filename)
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	progress := &ioprogress.Reader{
		Reader: resp.Body,
		Size:   resp.ContentLength,
	}

	_, err = io.Copy(out, progress)
	if err != nil {
		log.Fatal(err.Error())
	}

}
