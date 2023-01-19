mkdir -p apiRequests/items
rm ./apiRequests/items/*
cd apiRequests/items

fileName="itemsToAdd"

for file in {0..200}
do
    touch $fileName$file".json"
    echo "{" > $fileName$file".json"
    echo "    \"itemid\": "$file >> $fileName$file".json"
    echo "}" >> $fileName$file".json"
done
