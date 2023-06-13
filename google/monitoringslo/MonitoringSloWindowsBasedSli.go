package monitoringslo


type MonitoringSloWindowsBasedSli struct {
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) with ValueType = BOOL. The window is good if any true values appear in the window. One of 'good_bad_metric_filter', 'good_total_ratio_threshold', 'metric_mean_in_range', 'metric_sum_in_range' must be set for 'windows_based_sli'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/monitoring_slo#good_bad_metric_filter MonitoringSlo#good_bad_metric_filter}
	GoodBadMetricFilter *string `field:"optional" json:"goodBadMetricFilter" yaml:"goodBadMetricFilter"`
	// good_total_ratio_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/monitoring_slo#good_total_ratio_threshold MonitoringSlo#good_total_ratio_threshold}
	GoodTotalRatioThreshold *MonitoringSloWindowsBasedSliGoodTotalRatioThreshold `field:"optional" json:"goodTotalRatioThreshold" yaml:"goodTotalRatioThreshold"`
	// metric_mean_in_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/monitoring_slo#metric_mean_in_range MonitoringSlo#metric_mean_in_range}
	MetricMeanInRange *MonitoringSloWindowsBasedSliMetricMeanInRange `field:"optional" json:"metricMeanInRange" yaml:"metricMeanInRange"`
	// metric_sum_in_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/monitoring_slo#metric_sum_in_range MonitoringSlo#metric_sum_in_range}
	MetricSumInRange *MonitoringSloWindowsBasedSliMetricSumInRange `field:"optional" json:"metricSumInRange" yaml:"metricSumInRange"`
	// Duration over which window quality is evaluated, given as a duration string "{X}s" representing X seconds.
	//
	// Must be an
	// integer fraction of a day and at least 60s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/monitoring_slo#window_period MonitoringSlo#window_period}
	WindowPeriod *string `field:"optional" json:"windowPeriod" yaml:"windowPeriod"`
}

