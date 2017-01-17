package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	// Should be called at the very beginning of main().
	systray.Run(onReady)
}

func onReady() {
	systray.SetTitle("🤔")
	url := "https://sub.thegulocal.com/healthcheck"

	systray.AddMenuItem("Quit", "IDK, stop?")

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
			systray.SetTitle("🚂") //it's not running
			time.Sleep(time.Second * 10)

		case 503:
			systray.SetTitle("🔥") //it's broken
			time.Sleep(time.Second * 5)

		case 200:
			systray.SetTitle("🗞") //it's working
			time.Sleep(time.Second * 60)

		default:
			systray.SetTitle("🤔")
			time.Sleep(time.Second * 10)

		}
	}

}
