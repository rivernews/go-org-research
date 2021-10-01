package liscraper

import (
	"github.com/chromedp/chromedp"
)

func GetPopularBrowserOption(isHeadless bool, tmpDir string) []chromedp.ExecAllocatorOption {
	// configure options
	options := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
		chromedp.UserDataDir(tmpDir),
	}
	if !isHeadless {
		options = append(options,
			chromedp.Flag("headless", false),
			chromedp.Flag("hide-scrollbars", false),
			chromedp.Flag("mute-audio", false),
		)
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	return options
}
