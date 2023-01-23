curl 'https://www.reazon.ru/api/v1/user/profile' --http1.1 \
  -H 'authority: www.reazon.ru' \
  -H 'accept: application/json' \
  -H 'accept-language: en-US,en;q=0.9,ru;q=0.8' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'cookie: '$cookie \
  -H 'csrf: '$csrf \
  -H 'dnt: 1' \
  -H 'pragma: no-cache' \
  -H 'referer: https://www.reazon.ru/user' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  --compressed \
  > tempAddressID.txt

cat tempAddressID.txt | jq -r '.address[0].id' >> ./apiRequests/users/user$file.txt
cat tempAddressID.txt | jq -r '.id' >> ./apiRequests/users/user$file.txt
