package helpers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/jklq/bug-tracker/view"
)

func HtmxTemplate(c *fiber.Ctx) string {
	if IsHtmxRequest(c) {
		return ""
	}
	return "layouts/main"
}

func HtmxLayoutComponent(c *fiber.Ctx) templ.Component {
	if IsHtmxRequest(c) {
		return view.BasicEmpty()
	}
	return view.AppLayout(nil)
}
func HtmxLayoutComponentBasic(c *fiber.Ctx) templ.Component {
	if IsHtmxRequest(c) {
		return view.BasicEmpty()
	}
	return view.BasicLayout()
}
