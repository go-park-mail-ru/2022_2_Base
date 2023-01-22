fileName="itemsToAdd"

userNum=$(($RANDOM % $(ls ./apiRequests/users | wc -l)))
while [ $(head -n 1 ./apiRequests/users/user${userNum}.txt) == "" ]
do
  userNum=$(($RANDOM % $(ls ./apiRequests/users | wc -l)))
done

cookie=$(cat ./apiRequests/users/user${userNum}.txt | tail +2 | head -n 1)
csrf=$(cat ./apiRequests/users/user${userNum}.txt | tail +1 | head -n 1)
addressID=$(cat ./apiRequests/users/user${userNum}.txt | tail +5 | head -n 1)
itemsToAddCount=$(($RANDOM % 10 + 1))
itemsCount=200

echo "" > res.txt
items=""
for file in `seq 0 1 $itemsToAddCount`
do
    echo "${outString}POST https://www.reazon.ru/api/v1/cart/insertintocart" >> res.txt
    echo "${outString}Content-Type: application/json" >> res.txt
    echo "${outString}cookie:${cookie}" >> res.txt
    echo "${outString}csrf:${csrf}" >> res.txt
    currItem=$(($RANDOM % $itemsCount + 1))
    items="${items}, ${currItem}"
    echo -e "${outString}@./apiRequests/items/${fileName}${currItem}.json\n" >> res.txt
done

echo -e "{
    \"items\": [${items:2}],
    \"address\": ${addressID},
    \"deliverydate\": \"2023-01-20T10:00:00.000Z\",
    \"card\": 1,
    \"userid\": 1
}" > apiRequests/makeorder.json

echo "${outString}POST https://www.reazon.ru/api/v1/cart/makeorder" >> res.txt
echo "${outString}Content-Type: application/json" >> res.txt
echo "${outString}cookie:${cookie}" >> res.txt
echo "${outString}csrf:${csrf}" >> res.txt
echo "@./apiRequests/makeorder.json" >> res.txt

cat res.txt
rm res.txt
