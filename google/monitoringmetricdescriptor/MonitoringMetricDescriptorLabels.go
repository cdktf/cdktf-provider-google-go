package monitoringmetricdescriptor


type MonitoringMetricDescriptorLabels struct {
	// The key for this label.
	//
	// The key must not exceed 100 characters. The first character of the key must be an upper- or lower-case letter, the remaining characters must be letters, digits or underscores, and the key must match the regular expression [a-zA-Z][a-zA-Z0-9_]*
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/monitoring_metric_descriptor#key MonitoringMetricDescriptor#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A human-readable description for the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/monitoring_metric_descriptor#description MonitoringMetricDescriptor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of data that can be assigned to the label. Default value: "STRING" Possible values: ["STRING", "BOOL", "INT64"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/monitoring_metric_descriptor#value_type MonitoringMetricDescriptor#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

