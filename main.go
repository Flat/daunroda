package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strconv"
	"strings"
)

func main() {
	var booru string
	var output string
	var rating string
	var apikey string
	var user string
	var page int
	var count int
	var id int
	cwd, _ := os.Getwd()
	app := cli.NewApp()
	app.Name = "daunroda"
	app.Usage = "A simple command line booru mass image downloader. Arguments accepted are tags for images to download."
	app.Version = "0.0.3"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ken Swenson (flat)",
			Email: "flat@imo.uto.moe",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "b, booru",
			Usage:       "booru to download from. Valid options: konachan (Required)",
			Destination: &booru,
		},
		cli.StringFlag{
			Name:        "o, output",
			Usage:       "output directory",
			Destination: &output,
			Value:       cwd,
		},
		cli.StringFlag{
			Name:        "r, rating",
			Usage:       "image rating(s) to include. Valid values: safe, questionable, explicit or any '+' delimited combination. (e.g. -r safe OR -r safe+questionable)",
			Value:       "safe",
			Destination: &rating,
		},
		cli.IntFlag{
			Name:        "p, page",
			Usage:       "page to download from.",
			Value:       0,
			Destination: &page,
		},
		cli.IntFlag{
			Name:        "c, count",
			Usage:       "number of images to download. (Max: 100)",
			Value:       20,
			Destination: &count,
		},
		cli.IntFlag{
			Name:        "i, id",
			Usage:       "single image id to download",
			Destination: &id,
		},
		cli.StringFlag{
			Name:        "u, user",
			Usage:       "username (Currently only yandere)",
			Value:       "",
			Destination: &user,
		},
		cli.StringFlag{
			Name:        "password, api-key",
			Usage:       "API key to auth with (Currently only yandere)",
			Value:       "",
			Destination: &apikey,
		},
	}
	app.Action = func(c *cli.Context) {
		var tags = make([]string, len(c.Args()))
		if len(c.Args()) > 0 {
			tags = c.Args()
		} else {
			cli.ShowAppHelp(c)
			return
		}
		if booru == "" {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		tagString := strings.Join(tags, "+")
		p := request(booru, tagString, rating, page, count, user, apikey)
		for _, image := range p.Post {
			download(image.FileURL, image.Md5, strconv.Itoa(image.Id), output)
		}

	}
	app.Run(os.Args)
}
