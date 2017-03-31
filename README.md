# gateway-pusher

Simple utility to push a metric for a successful run of a batch job to a prometheus pushgateway.

# Usage
```
Usage of ./gateway-pusher:
  -datacenter string
    	Datacenter label to add (default "eu-west-1")
  -environment string
    	Environment label to add (default "test")
  -job-name string
    	Name of the job to report on (default "batch_job")
  -pushgateway-url string
    	Pushgateway to use
```
