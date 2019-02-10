package chromedpcv

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"

	"github.com/pkg/errors"
)

func logOnlyIfErr(err error) {
	if err != nil {
		log.Printf("%+v\n", errors.Wrap(err, "WARN:"))
	}
}

func fileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func screenShotToTempFile(ctxt context.Context, h cdp.Executor) (tempFilePath string, err error) {
	var screenShotBytes []byte
	err = chromedp.CaptureScreenshot(&screenShotBytes).Do(ctxt, h)
	if err != nil {
		return tempFilePath, errors.Wrap(err, `chromedp.CaptureScreenshot`)
	}
	screenShotFile, err := ioutil.TempFile("", "screenshot_*.png")
	if err != nil {
		return tempFilePath, errors.Wrap(err, `create tempfile`)
	}
	tempFilePath = screenShotFile.Name()
	if _, err = screenShotFile.Write(screenShotBytes); err != nil {
		return tempFilePath, errors.Wrap(err, `write data to tempfile:`+tempFilePath)
	}
	err = screenShotFile.Close()
	if err != nil {
		return tempFilePath, errors.Wrap(err, `closing tempfile:`+tempFilePath)
	}
	return tempFilePath, err
}
