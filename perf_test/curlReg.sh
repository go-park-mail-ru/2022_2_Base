mail="$(echo $RANDOM | md5sum | head -c 15; echo)@m.ru"
pwd="$(echo $RANDOM | md5sum | head -c 15; echo)"

curl -sSL -D - 'https://www.reazon.ru/api/v1/signup' \
  -H 'authority: www.reazon.ru' \
  -H 'accept: application/json' \
  -H 'accept-language: en-US,en;q=0.9,ru-RU;q=0.8,ru;q=0.7' \
  -H 'content-type: application/json' \
  -H 'csrf: null' \
  -H 'dnt: 1' \
  -H 'origin: https://www.reazon.ru' \
  -H 'referer: https://www.reazon.ru/signup' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-gpc: 1' \
  --data-raw '{"password":"'"${pwd}"'","email":"'"${mail}"'","username":"kjn"}' \
  --compressed \
  > tempCurl.txt

  echo "mail: ${mail}" >> tempCurl.txt
  echo "pwd: ${pwd}" >> tempCurl.txt