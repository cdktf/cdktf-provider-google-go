package monitoringslo


type MonitoringSloWindowsBasedSliGoodTotalRatioThreshold struct {
	// basic_sli_performance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_slo#basic_sli_performance MonitoringSlo#basic_sli_performance}
	BasicSliPerformance *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance `field:"optional" json:"basicSliPerformance" yaml:"basicSliPerformance"`
	// performance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_slo#performance MonitoringSlo#performance}
	Performance *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance `field:"optional" json:"performance" yaml:"performance"`
	// If window performance >= threshold, the window is counted as good.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_slo#threshold MonitoringSlo#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

