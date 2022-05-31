package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	var url string
	var err error

	flag.StringVar(&url, "url", "", "open url")
	flag.Parse()

	if len(strings.TrimSpace(url)) == 0 {
		url = "https://www.youtube.com/watch?v=dQw4w9WgXcQ" //default url
	}

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
		if err != nil {
			log.Fatal("please install xdg-open, run: sudo apt install xdg-utils")
		}
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupport os")
	}
	if err != nil {
		log.Fatal(err)
	}
}
