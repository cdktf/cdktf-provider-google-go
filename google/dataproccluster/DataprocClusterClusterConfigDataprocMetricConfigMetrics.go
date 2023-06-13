package dataproccluster


type DataprocClusterClusterConfigDataprocMetricConfigMetrics struct {
	// A source for the collection of Dataproc OSS metrics (see [available OSS metrics] (https://cloud.google.com//dataproc/docs/guides/monitoring#available_oss_metrics)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#metric_source DataprocCluster#metric_source}
	MetricSource *string `field:"required" json:"metricSource" yaml:"metricSource"`
	// Specify one or more [available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics) to collect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#metric_overrides DataprocCluster#metric_overrides}
	MetricOverrides *[]*string `field:"optional" json:"metricOverrides" yaml:"metricOverrides"`
}

