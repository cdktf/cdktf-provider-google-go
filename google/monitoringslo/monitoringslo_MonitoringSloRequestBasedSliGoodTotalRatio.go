package monitoringslo


type MonitoringSloRequestBasedSliGoodTotalRatio struct {
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) quantifying bad service provided, either demanded service that was not provided or demanded service that was of inadequate quality.
	//
	// Must have ValueType = DOUBLE or ValueType = INT64 and
	// must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Exactly two of 'good_service_filter','bad_service_filter','total_service_filter'
	// must be set (good + bad = total is assumed).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#bad_service_filter MonitoringSlo#bad_service_filter}
	BadServiceFilter *string `field:"optional" json:"badServiceFilter" yaml:"badServiceFilter"`
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) quantifying good service provided. Must have ValueType = DOUBLE or ValueType = INT64 and must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Exactly two of 'good_service_filter','bad_service_filter','total_service_filter'
	// must be set (good + bad = total is assumed).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#good_service_filter MonitoringSlo#good_service_filter}
	GoodServiceFilter *string `field:"optional" json:"goodServiceFilter" yaml:"goodServiceFilter"`
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) quantifying total demanded service.
	//
	// Must have ValueType = DOUBLE or ValueType = INT64 and
	// must have MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Exactly two of 'good_service_filter','bad_service_filter','total_service_filter'
	// must be set (good + bad = total is assumed).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#total_service_filter MonitoringSlo#total_service_filter}
	TotalServiceFilter *string `field:"optional" json:"totalServiceFilter" yaml:"totalServiceFilter"`
}

