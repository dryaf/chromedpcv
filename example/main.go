package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/rand99/chromedpcv"
)

func main() {

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	exampleDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	c, err := chromedp.New(ctxt,
		chromedp.WithRunnerOptions(
			runner.URL("https://jsfiddle.net/jchoL7ge/16/"),
			runner.Path(exampleDir+"/chrome-mac/Chromium.app/Contents/MacOS/Chromium"),
			runner.Flag("disable-client-side-phishing-detection", true),
			//runner.Flag("headless", true),
			runner.Flag("disable-gpu", true),
			runner.Flag("no-first-run", true),
			runner.Flag("disable-extensions", true),
			runner.Flag("disable-translate", true),
			runner.Flag("disable-web-security", true),
			runner.Flag("no-proxy-server", true),
			runner.Flag("timeout", 30000),
			runner.Flag("remote-debugging-port", 9222),
			runner.Flag("no-sandbox", true),
			runner.Flag("enable-logging", true),         // run faster
			runner.Flag("disable-setuid-sandbox", true), // run faster
			runner.Flag("no-default-browser-check", true)),
		chromedp.WithLog(func(s string, i ...interface{}) {
			if strings.Contains(fmt.Sprint(s, i), "chromedpcv") {
				fmt.Println(s, i)
			}
		}))
	if err != nil {
		log.Fatal(err)
	}
	// create new chromedpcv instance (default config will be set)
	chromecv := chromedpcv.New()

	// save screenshot and mark the image that has been found within
	chromecv.TemplateMatchMarkedScreenShotFilePath = "match.png"

	var imageSearchPosition chromedpcv.BrowserWindowPosition
	_ = imageSearchPosition

	// run task list
	fmt.Println("Run Tasks")
	err = c.Run(ctxt, chromedp.Tasks{
		WaitSeconds(10),
		chromecv.PositionWhereScreenLooksLike("./search_right_black_rect.png", &imageSearchPosition),

		// Don't make this mistake both values will be 0
		DebugPosition(imageSearchPosition.X, imageSearchPosition.Y),
		// MouseClickXY won't work
		// Use pointers
		DebugPositionPointer(&imageSearchPosition),
		chromecv.MouseClickAtPosition(&imageSearchPosition),

		// or search for lookalike region and click it immediately
		chromecv.MouseClickWhereScreenLooksLike("./search_right_black_rect.png"),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("let the amazing mouse click breath a little")
	time.Sleep(5 * time.Second)
	fmt.Println("take a look at match.png")
	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Good night!")
}

func DebugPosition(x, y int64) chromedp.Action {
	return chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
		fmt.Println("values: x: ", x, " y: ", y)
		return nil
	})
}

func DebugPositionPointer(position *chromedpcv.BrowserWindowPosition) chromedp.Action {
	return chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
		fmt.Println("pointer: x: ", position.X, " y: ", position.Y)
		return nil
	})
}

func WaitSeconds(num int64) chromedp.Action {
	return chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
		fmt.Print("Waiting ", num, " seconds ...")
		time.Sleep(time.Duration(num) * time.Second)
		fmt.Println("Done.")
		return nil
	})
}
