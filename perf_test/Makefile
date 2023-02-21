run-post-test:
	bash genTestCase.sh | vegeta attack -duration=40s -rate=250 -max-workers=1000 -http2 | tee results.bin | vegeta report

run-get-test:
	bash genGetTest.sh | vegeta attack -duration=2223s -rate=450 -max-workers=1000 -http2 | tee results.bin | vegeta report

run-test-static:
	vegeta attack -duration=1s -rate=5 -http2 -targets=targets.txt | tee results.bin | vegeta report

get-report:run-post-test
	vegeta report -type=json results.bin > metrics.json && \
	cat results.bin | vegeta plot > plot.html && \
	cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms, 500ms, 800ms, 1000ms]"
