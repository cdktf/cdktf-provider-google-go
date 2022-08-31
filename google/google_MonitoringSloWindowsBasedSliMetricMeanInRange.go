// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type MonitoringSloWindowsBasedSliMetricMeanInRange struct {
	// range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#range MonitoringSlo#range}
	Range *MonitoringSloWindowsBasedSliMetricMeanInRangeRange `field:"required" json:"range" yaml:"range"`
	// A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) specifying the TimeSeries to use for evaluating window The provided TimeSeries must have ValueType = INT64 or ValueType = DOUBLE and MetricKind = GAUGE. Mean value 'X' should satisfy 'range.min <= X <= range.max' under good service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#time_series MonitoringSlo#time_series}
	TimeSeries *string `field:"required" json:"timeSeries" yaml:"timeSeries"`
}

