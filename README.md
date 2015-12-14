# daunroda
Mass download images from various booru sites.


NAME:
   daunroda - A simple command line booru mass image downloader. Arguments accepted are tags for images to download.

USAGE:
   daunroda.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR(S):
   Ken Swenson (flat) <flat@imo.uto.moe>

COMMANDS:
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS: | Description
------------ | -------------
   -b, --booru                                                                  |booru to download from. (Required)
   -o, --output "./" |output directory
   -r, --rating "safe"                                                          |image rating(s) to include. Valid values: safe, questionable, explicit or any '+' delimited combination. (e.g. -r safe OR -r safe+questionable)
   -p, --page "0"                                                               |page to download from.
   -c, --count "20"                                                             |number of images to download. (Max: 100)
   -i, --id "0"                                                                 |single image id to download
   --help, -h                                                                   |show help
   --version, -v                                                                |print the version
