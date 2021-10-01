package scraperlog

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
	"github.com/rivernews/GoTools"
)

func LogTitle() chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var title string
		err := chromedp.Title(&title).Do(ctx)
		if err != nil {
			GoTools.SimpleLogger("ERROR", err.Error())
		}
		GoTools.SimpleLogger("INFO", fmt.Sprintf(`Title is "%s"`, title))
		return nil
	})
}
