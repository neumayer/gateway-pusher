package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func pushSuccessfulCompletion(pushGatewayURL string, jobName string) error {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: jobName + "_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a batch job.",
	})
	completionTime.Set(float64(time.Now().Unix()))
	return push.Collectors(
		jobName,
		push.HostnameGroupingKey(),
		pushGatewayURL,
		completionTime,
	)
}

func main() {
	pushGatewayURL := flag.String("pushgateway-url", "http://localhost:9091", "Pushgateway to use")
	jobName := flag.String("job-name", "batch_job", "Name of the job to report on")
	flag.Parse()

	err := pushSuccessfulCompletion(*pushGatewayURL, *jobName)
	if err != nil {
		fmt.Println(err)
	}
}
