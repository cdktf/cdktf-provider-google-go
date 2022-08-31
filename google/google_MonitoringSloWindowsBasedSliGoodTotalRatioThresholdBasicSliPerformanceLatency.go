// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency struct {
	// A duration string, e.g. 10s. Good service is defined to be the count of requests made to this service that return in no more than threshold.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#threshold MonitoringSlo#threshold}
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

