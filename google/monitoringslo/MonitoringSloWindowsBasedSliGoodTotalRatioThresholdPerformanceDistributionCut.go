package monitoringslo


type MonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCut struct {
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) aggregating values to quantify the good service provided.
	//
	// Must have ValueType = DISTRIBUTION and
	// MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/monitoring_slo#distribution_filter MonitoringSlo#distribution_filter}
	DistributionFilter *string `field:"required" json:"distributionFilter" yaml:"distributionFilter"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/monitoring_slo#range MonitoringSlo#range}
	Range *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCutRange `field:"required" json:"range" yaml:"range"`
}

