package monitoringslo


type MonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio struct {
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) quantifying bad service provided, either demanded service that was not provided or demanded service that was of inadequate quality. Exactly two of good, bad, or total service filter must be defined (where good + bad = total is assumed).
	//
	// Must have ValueType = DOUBLE or ValueType = INT64 and
	// must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/monitoring_slo#bad_service_filter MonitoringSlo#bad_service_filter}
	BadServiceFilter *string `field:"optional" json:"badServiceFilter" yaml:"badServiceFilter"`
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) quantifying good service provided. Exactly two of good, bad, or total service filter must be defined (where good + bad = total is assumed).
	//
	// Must have ValueType = DOUBLE or ValueType = INT64 and
	// must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/monitoring_slo#good_service_filter MonitoringSlo#good_service_filter}
	GoodServiceFilter *string `field:"optional" json:"goodServiceFilter" yaml:"goodServiceFilter"`
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) quantifying total demanded service. Exactly two of good, bad, or total service filter must be defined (where good + bad = total is assumed).
	//
	// Must have ValueType = DOUBLE or ValueType = INT64 and
	// must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/monitoring_slo#total_service_filter MonitoringSlo#total_service_filter}
	TotalServiceFilter *string `field:"optional" json:"totalServiceFilter" yaml:"totalServiceFilter"`
}

