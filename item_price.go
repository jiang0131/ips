package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Item struct {
	city  string
	title string
}

type ItemPriceCache struct {
	itemModePrice  map[Item]int
	itemTotalCnt   map[Item]int
	titleModePrice map[string]int
	titleTotalCnt  map[string]int
}

var priceCache *ItemPriceCache

func ItemPriceCacheInit() error {

	priceCache = &ItemPriceCache{itemModePrice: map[Item]int{}, itemTotalCnt: map[Item]int{}, titleModePrice: map[string]int{}, titleTotalCnt: map[string]int{}}

	if err := refreshItemPrice(); err != nil {
		return err
	}

	go startPollingPriceCache()

	return nil
}

func refreshItemPrice() error {
	options := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", DB_HOST, DB_PORT, DB_USERID, DB_PWD, ITEMPRICES_DB)
	db, err := sql.Open("postgres", options)
	defer db.Close()

	if err != nil {
		return err
	}

	priceCacheNew := ItemPriceCache{itemModePrice: map[Item]int{}, itemTotalCnt: map[Item]int{}, titleModePrice: map[string]int{}, titleTotalCnt: map[string]int{}}

	rows, err := db.Query(ITEM_PRICE_QUERY)
	defer rows.Close()

	for rows.Next() {
		var city, title string
		var modePrice, cnt, rank int
		if err := rows.Scan(&city, &title, &modePrice, &cnt, &rank); err != nil {
			fmt.Println("scan err=", err)
		}
		if rank == 1 {
			priceCacheNew.itemModePrice[Item{city, title}] = modePrice
		}
		if _, ok := priceCacheNew.itemTotalCnt[Item{city, title}]; ok {
			priceCacheNew.itemTotalCnt[Item{city, title}] += cnt
		} else {
			priceCacheNew.itemTotalCnt[Item{city, title}] = cnt
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	rows, err = db.Query(TITLE_PRICE_QUERY)

	for rows.Next() {
		var title string
		var listPrice, cnt, rank int
		if err := rows.Scan(&title, &listPrice, &cnt, &rank); err != nil {
			fmt.Println("scan err=", err)
		}
		if rank == 1 {
			priceCacheNew.titleModePrice[title] = listPrice
		}
		if _, ok := priceCacheNew.titleTotalCnt[title]; ok {
			priceCacheNew.titleTotalCnt[title] += cnt
		} else {
			priceCacheNew.titleTotalCnt[title] = cnt
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	priceCache = &priceCacheNew

	logData := make(map[string]interface{})
	logData["name"] = "ItemPriceCache"
	logData["message"] = "ItemPriceCache Refreshed"
	logger.info(logData)
	fmt.Println("ItemPriceCache Refreshed at", time.Now())

	return nil
}

func startPollingPriceCache() {
	for _ = range time.Tick(ITEM_PRICE_CACHE_REFRESH_INTERVAL_MINS * time.Minute) {
		if err := refreshItemPrice(); err != nil {
			logData := make(map[string]interface{})
			logData["name"] = "ItemPriceCache"
			logData["message"] = err.Error()
			logger.error(logData)
		}
	}
}
