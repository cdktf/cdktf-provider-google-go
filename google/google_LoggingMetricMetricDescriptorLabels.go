// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type LoggingMetricMetricDescriptorLabels struct {
	// The label key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#key LoggingMetric#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A human-readable description for the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#description LoggingMetric#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of data that can be assigned to the label. Default value: "STRING" Possible values: ["BOOL", "INT64", "STRING"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#value_type LoggingMetric#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

