package monitoringmetricdescriptor


type MonitoringMetricDescriptorMetadata struct {
	// The delay of data points caused by ingestion.
	//
	// Data points older than this age are guaranteed to be ingested and available to be read, excluding data loss due to errors. In '[duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration)'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/monitoring_metric_descriptor#ingest_delay MonitoringMetricDescriptor#ingest_delay}
	IngestDelay *string `field:"optional" json:"ingestDelay" yaml:"ingestDelay"`
	// The sampling period of metric data points.
	//
	// For metrics which are written periodically, consecutive data points are stored at this time interval, excluding data loss due to errors. Metrics with a higher granularity have a smaller sampling period. In '[duration format](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf?&_ga=2.264881487.1507873253.1593446723-935052455.1591817775#google.protobuf.Duration)'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/monitoring_metric_descriptor#sample_period MonitoringMetricDescriptor#sample_period}
	SamplePeriod *string `field:"optional" json:"samplePeriod" yaml:"samplePeriod"`
}

