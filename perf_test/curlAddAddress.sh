curl 'https://www.reazon.ru/api/v1/user/profile' --http1.1 \
  -H 'authority: www.reazon.ru' \
  -H 'accept: application/json' \
  -H 'accept-language: en-US,en;q=0.9,ru-RU;q=0.8,ru;q=0.7' \
  -H 'content-type: application/json' \
  -H 'cookie: '$cookie \
  -H 'csrf: '$csrf \
  -H 'dnt: 1' \
  -H 'origin: https://www.reazon.ru' \
  -H 'referer: https://www.reazon.ru/user' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-gpc: 1' \
  --data-raw '{"username":"kjn","email":"'"${mail}"'","paymentmethods":[],"address":[{"city":"1","street":"2","house":"3","flat":"4","id":-1,"priority":true}]}' \
  --compressed \
    > /dev/null
