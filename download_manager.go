package main

import (
	"fmt"
	"github.com/mitchellh/ioprogress"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func download(url string, md5String string, id string, outDir string) {
	filename := id + "." + strings.Split(url, ".")[len(strings.Split(url, "."))-1]
	out, err := os.Create(outDir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	bar := ioprogress.DrawTextFormatBar(50)
	myDrawFunc := ioprogress.DrawTerminalf(os.Stdout, func(progress, total int64) string {
		return fmt.Sprintf("%s %s %30s", filename, bar(progress, total), ioprogress.DrawTextFormatBytes(progress, total))
	})

	progress := &ioprogress.Reader{
		Reader:       resp.Body,
		Size:         resp.ContentLength,
		DrawInterval: time.Millisecond,
		DrawFunc:     myDrawFunc,
	}

	_, err = io.Copy(out, progress)
	if err != nil {
		log.Fatal(err)
	}

}
