package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var stored = 0

const bucket_size = 10

var last_stamped_t = 0

const period = 10

func LeakyBucket(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if stored < bucket_size {
			stored++
		} else {
			return c.JSON(http.StatusTooManyRequests, 429)
		}
		tm := time.Now()

		fmt.Println(tm)
		return next(c)
	}
}
