package monitoringslo


type MonitoringSloRequestBasedSliDistributionCut struct {
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) aggregating values to quantify the good service provided.
	//
	// Must have ValueType = DISTRIBUTION and
	// MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#distribution_filter MonitoringSlo#distribution_filter}
	DistributionFilter *string `field:"required" json:"distributionFilter" yaml:"distributionFilter"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#range MonitoringSlo#range}
	Range *MonitoringSloRequestBasedSliDistributionCutRange `field:"required" json:"range" yaml:"range"`
}

