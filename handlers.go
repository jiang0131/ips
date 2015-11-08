package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func GetPriceHandler(req *http.Request, r render.Render, c martini.Context) {
	title := req.URL.Query().Get("item")
	city := req.URL.Query().Get("city")

	if len(title) > 0 {
		if len(city) > 0 {
			context := map[string]interface{}{
				"item":             title,
				"item_count":       priceCache.itemTotalCnt[Item{city, title}],
				"price_suggestion": priceCache.itemModePrice[Item{city, title}],
				"city":             city,
			}
			response := map[string]interface{}{
				"status":  http.StatusOK,
				"context": context,
			}
			r.JSON(http.StatusOK, response)
		} else {
			context := map[string]interface{}{
				"item":             title,
				"item_count":       priceCache.titleTotalCnt[title],
				"price_suggestion": priceCache.titleModePrice[title],
			}
			response := map[string]interface{}{
				"status":  http.StatusOK,
				"context": context,
			}
			r.JSON(http.StatusOK, response)
		}
	} else {
		context := map[string]interface{}{
			"message": "Not found",
		}
		response := map[string]interface{}{
			"status":  http.StatusNotFound,
			"context": context,
		}
		r.JSON(http.StatusNotFound, response)
	}
}
