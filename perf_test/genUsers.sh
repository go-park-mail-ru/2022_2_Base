mkdir -p apiRequests/users
rm ./apiRequests/users/*

for file in {0..1000}
do
    bash curlReg.sh
    export file
    export csrf=$(cat tempCurl.txt | grep csrf: | cut -b 7-)
    echo $csrf > ./apiRequests/users/user$file.txt
    export cookie=$(cat tempCurl.txt | grep cookie: | cut -b 13- | cut -f 1 -d";")
    echo $cookie >> ./apiRequests/users/user$file.txt
    export mail=$(cat tempCurl.txt | grep mail: | cut -b 9-)
    echo $mail >> ./apiRequests/users/user$file.txt
    export pwd=$(cat tempCurl.txt | grep pwd: | cut -b 6-)
    echo $pwd >> ./apiRequests/users/user$file.txt
    bash curlAddAddress.sh
    bash curlGetAddressID.sh
done

rm tempCurl.txt
rm tempAddress.txt
rm tempAddressID.txt
