run-test:
	vegeta attack -duration=1s -rate=5 -http2 -targets=targets.txt | tee results.bin | vegeta report

get-report:run-test
	vegeta report -type=json results.bin > metrics.json && \
	cat results.bin | vegeta plot > plot.html && \
	cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
  	
gen-user:
	jq -ncM 'while(true; .+1) | {method: "POST", url: "https://reazon.ru/api/v1/signup", body: {"email": [.]} | @base64, header: {"Content-Type": ["application/json"]} }'	

# echo "POST https://reazon.ru/api/v1/signup" | vegeta attack -duration=10s -rate=500 -http2 | tee results.bin | vegeta report
