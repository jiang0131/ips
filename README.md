Item Price Service (IPS)
===
A microservice to provide listing price recommendation given an item.


Build
-------------

    go build

Run
-------------
    ./ips

Endpoints
-------------
* GET /item-price-service/

  Query parameters:
  - item [string]
  - city [string]