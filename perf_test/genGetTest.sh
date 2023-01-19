userNum=$(($RANDOM % $(ls ./apiRequests/users | wc -l)))
cookie=$(cat ./apiRequests/users/user${userNum}.txt | tail +2 | head -n 1)  #"session_id=527a65c1-57cf-403f-9093-be243079c80a"
csrf=$(cat ./apiRequests/users/user${userNum}.txt | tail +1 | head -n 1)

echo "" > res.txt

echo "${outString}GET https://www.reazon.ru/api/v1/cart/orders" >> res.txt
echo "${outString}Content-Type: application/json" >> res.txt
echo "${outString}cookie:${cookie}" >> res.txt
echo "${outString}csrf:${csrf}" >> res.txt

cat res.txt
rm res.txt
