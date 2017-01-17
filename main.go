package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	// Should be called at the very beginning of main().
	systray.Run(start)
}

func start() {
	systray.SetTitle("🤔")

	quit := systray.AddMenuItem("Quit", "IDK, stop?")
	fmt.Println("starting")
	name := os.Args[2]
	url := os.Args[1]
	systray.SetTooltip(url)
	systray.AddMenuItem(name+": "+url, url)
	go monitor(url, name)
	<-quit.ClickedCh
	fmt.Println("stopping")
	os.Exit(0)
}

func monitor(url string, name string) {
	for {
		resp, err := http.Get(url)
		code := resp.StatusCode
		resp.Body.Close()
		fmt.Println(code)
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch code {
		case 502:
			systray.SetTitle("🚂" + name) //it's not running
			time.Sleep(time.Second * 10)

		case 503:
			systray.SetTitle("🔥" + name) //it's broken
			time.Sleep(time.Second * 5)

		case 200:
			systray.SetTitle("👍" + name) //it's working
			time.Sleep(time.Second * 60)

		default:
			systray.SetTitle("🤔" + name)
			time.Sleep(time.Second * 10)

		}
	}

}
