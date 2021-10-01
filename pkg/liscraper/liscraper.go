package liscraper

import (
	"context"
	"fmt"
	"strings"

	"github.com/rivernews/GoTools"

	"github.com/chromedp/chromedp"
)

func GlobalSearch(keyword string) chromedp.Tasks {
	searchBarSel := `body > div.application-outlet > #global-nav > div.global-nav__content > div#global-nav-search > div.search-global-typeahead > div#global-nav-typeahead > input.search-global-typeahead__input`
	return chromedp.Tasks{
		chromedp.SendKeys(searchBarSel, keyword, chromedp.ByQuery),
	}
}

func CheckBlocked(codePtr *string) chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		title := GetTitle(ctx)
		if strings.Contains(title, "Security") {
			GoTools.SimpleLogger("WARN", fmt.Sprintf(`Page security check - requires human engagement. Page title: %s`, title))
			fmt.Printf(`Please enter code that would help unblock:`)
			fmt.Scanln("%s", codePtr)
			GoTools.SimpleLogger("INFO", fmt.Sprintf(`You typed in %s`, *codePtr))
		}
		return nil
	})
}

func GetTitle(ctx context.Context) string {
	var title string
	err := chromedp.Run(ctx, chromedp.Title(&title))
	if err != nil {
		GoTools.SimpleLogger("ERROR", err.Error())
	}
	return title
}
