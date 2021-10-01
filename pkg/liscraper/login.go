package liscraper

import (
	"github.com/chromedp/chromedp"
)

func Login(email string, psw string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.linkedin.com/login`),
		chromedp.SendKeys(`#username`, email),
		chromedp.SendKeys(`#password`, psw),
		chromedp.Click(`button[type=submit][aria-label="Sign in"]`),
	}
}
