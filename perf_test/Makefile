run-test:
	bash genTestCase.sh | vegeta attack -duration=10s -rate=infinity -max-workers=5000 -http2 | tee results.bin | vegeta report

run-test-static:
	vegeta attack -duration=1s -rate=5 -http2 -targets=targets.txt | tee results.bin | vegeta report

get-report:run-test
	vegeta report -type=json results.bin > metrics.json && \
	cat results.bin | vegeta plot > plot.html && \
	cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"

live-report:run-test
	jaggr @count=rps \
        hist\[100,200,300,400,500\]:code \
        p25,p50,p95:latency \
        sum:bytes_in \
        sum:bytes_out | \
    jplot rps+code.hist.100+code.hist.200+code.hist.300+code.hist.400+code.hist.500 \
        latency.p95+latency.p50+latency.p25 \
        bytes_in.sum+bytes_out.sum	
  	
	