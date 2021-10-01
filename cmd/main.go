package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/rivernews/GoTools"
	"github.com/rivernews/go-org-research/pkg/liscraper"
	"github.com/rivernews/go-org-research/pkg/scraperlog"

	"github.com/chromedp/chromedp"
)

func main() {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		GoTools.SimpleLogger("ERROR", err.Error())
	}
	defer os.RemoveAll(dir)

	// configure options
	options := liscraper.GetPopularBrowserOption(false, dir)

	// create chrome instance
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
		// chromedp.WithDebugf(log.Printf),
		chromedp.WithErrorf(log.Fatalf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// start the browser
	// see https://github.com/chromedp/chromedp/issues/513#issuecomment-558122963
	if err := chromedp.Run(ctx); err != nil {
		GoTools.SimpleLogger("ERROR", err.Error())
	}

	// navigate to a page, wait for an element, click
	var unblockCode string
	email := `user@example.com`
	psw := `*****`
	if err = chromedp.Run(ctx,
		liscraper.Login(email, psw),
		scraperlog.LogTitle(),
		liscraper.CheckBlocked(&unblockCode),
		liscraper.GlobalSearch(`bbc`),
		scraperlog.LogTitle(),
	); err != nil {
		GoTools.SimpleLogger("ERROR", err.Error())
	}
}
