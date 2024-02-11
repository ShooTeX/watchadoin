package checker

import (
	"github.com/playwright-community/playwright-go"
)

type Checker struct {
	*playwright.Playwright
	currentValue *string
	url          string
	selector     string
}

func New(pw *playwright.Playwright, url, selector string) *Checker {
	return &Checker{
		Playwright:   pw,
		currentValue: nil,
		url:          url,
		selector:     selector,
	}
}

type IsSameResponse struct {
	OldValue *string
	NewValue *string
	IsSame   bool
}

func (c *Checker) IsSame() (IsSameResponse, error) {
	response := IsSameResponse{}

	browser, err := c.Chromium.Launch()
	if err != nil {
		return response, err
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		return response, err
	}
	defer page.Close()

	if _, err = page.Goto(c.url); err != nil {
		return response, err
	}

	value, err := page.Locator(c.selector).First().TextContent()
	if err != nil {
		return response, err
	}

	if c.currentValue == nil {
		response.OldValue = new(string)
		response.NewValue = &value

		c.currentValue = &value

		return response, err
	}

	if *c.currentValue != value {
		response.OldValue = c.currentValue
		response.NewValue = &value

		c.currentValue = &value

		return response, err
	}

	response.IsSame = true
	response.OldValue = c.currentValue
	response.NewValue = &value

	return response, nil
}
