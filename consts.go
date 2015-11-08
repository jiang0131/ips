package main

const (
	SERVICE_NAME                           = "ips"
	LOG_FILE_PATH                          = "./ips.log"
	ITEM_PRICE_CACHE_REFRESH_INTERVAL_MINS = 10
)

const (
	DB_HOST        = "offerupchallenge.cgtzqpsohu0g.us-east-1.rds.amazonaws.com"
	DB_PORT        = "5432"
	DB_USERID      = "offerupchallenge"
	DB_PWD         = "ouchallenge"
	ITEMPRICES_DB  = "itemprices"
	ITEMSALE_TABLE = "itemPrices_itemsale"
)

const ITEM_PRICE_QUERY = `
	SELECT
	city, title, list_price, cnt,
	RANK() OVER(PARTITION BY city, title ORDER BY cnt desc, list_price desc) AS rank
	FROM
	(
	  SELECT city, title, list_price, count(*) AS cnt
	  FROM "itemPrices_itemsale"
	  GROUP BY city, title, list_price
	) AS X
	;
	`

const TITLE_PRICE_QUERY = `
	SELECT
	title, list_price, cnt,
	RANK() OVER(PARTITION BY title ORDER BY cnt desc, list_price desc) AS rank
	FROM
	(
	  SELECT title, list_price, count(*) AS cnt
	  FROM "itemPrices_itemsale"
	  GROUP BY title, list_price
	) AS X
	;
	`
