package routes

import (
	"github.com/vikassfteng/url-short/api/database"
)


func ResolveURLc *fiber.Ctx) error {
	url := c.Params("url")
	rdb := database.CreateClient(0)
	defer rdb.Close()
	value, err := r.Get(database.Ctx, url)
	if err ==redis.nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found in found database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to database",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}