package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func pushSuccessfulCompletion(pushGatewayURL string, jobName string, environment string, datacenter string) error {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: jobName + "_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a batch job.",
		ConstLabels: map[string]string{
			"environment": environment,
			"datacenter":  datacenter},
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
	pushGatewayURL := flag.String("pushgateway-url", "", "Pushgateway to use")
	jobName := flag.String("job-name", "batch_job", "Name of the job to report on")
	environment := flag.String("environment", "test", "Environment label to add")
	datacenter := flag.String("datacenter", "eu-west-1", "Datacenter label to add")
	flag.Parse()

	if *pushGatewayURL == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := pushSuccessfulCompletion(*pushGatewayURL, *jobName, *environment, *datacenter)
	if err != nil {
		fmt.Println(err)
	}
}
