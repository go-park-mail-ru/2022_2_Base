fileName="itemsToAdd"
cookie="session_id=527a65c1-57cf-403f-9093-be243079c80a"
csrf="2c7a9a27ffeb5897e8a6797953d5e455e16625d32cd517c92cff66caf3076008:1674192578"
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
    echo -e "${outString}@./items/${fileName}${currItem}.json\n" >> res.txt
done

echo -e "{
    \"items\": [${items:2}],
    \"address\": 1,
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