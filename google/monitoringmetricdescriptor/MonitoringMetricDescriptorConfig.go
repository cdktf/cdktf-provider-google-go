package monitoringmetricdescriptor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitoringMetricDescriptorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A detailed description of the metric, which can be used in documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#description MonitoringMetricDescriptor#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// A concise name for the metric, which can be displayed in user interfaces.
	//
	// Use sentence case without an ending period, for example "Request count".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#display_name MonitoringMetricDescriptor#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Whether the metric records instantaneous values, changes to a value, etc.
	//
	// Some combinations of metricKind and valueType might not be supported. Possible values: ["METRIC_KIND_UNSPECIFIED", "GAUGE", "DELTA", "CUMULATIVE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#metric_kind MonitoringMetricDescriptor#metric_kind}
	MetricKind *string `field:"required" json:"metricKind" yaml:"metricKind"`
	// The metric type, including its DNS name prefix.
	//
	// The type is not URL-encoded. All service defined metrics must be prefixed with the service name, in the format of {service name}/{relative metric name}, such as cloudsql.googleapis.com/database/cpu/utilization. The relative metric name must have only upper and lower-case letters, digits, '/' and underscores '_' are allowed. Additionally, the maximum number of characters allowed for the relative_metric_name is 100. All user-defined metric types have the DNS name custom.googleapis.com, external.googleapis.com, or logging.googleapis.com/user/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#type MonitoringMetricDescriptor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether the measurement is an integer, a floating-point number, etc.
	//
	// Some combinations of metricKind and valueType might not be supported. Possible values: ["BOOL", "INT64", "DOUBLE", "STRING", "DISTRIBUTION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#value_type MonitoringMetricDescriptor#value_type}
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#id MonitoringMetricDescriptor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#labels MonitoringMetricDescriptor#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The launch stage of the metric definition. Possible values: ["LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#launch_stage MonitoringMetricDescriptor#launch_stage}
	LaunchStage *string `field:"optional" json:"launchStage" yaml:"launchStage"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#metadata MonitoringMetricDescriptor#metadata}
	Metadata *MonitoringMetricDescriptorMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#project MonitoringMetricDescriptor#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#timeouts MonitoringMetricDescriptor#timeouts}
	Timeouts *MonitoringMetricDescriptorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The units in which the metric value is reported.
	//
	// It is only applicable if the
	// valueType is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of
	// the stored metric values.
	//
	// Different systems may scale the values to be more easily displayed (so a value of
	// 0.02KBy might be displayed as 20By, and a value of 3523KBy might be displayed as
	// 3.5MBy). However, if the unit is KBy, then the value of the metric is always in
	// thousands of bytes, no matter how it may be displayed.
	//
	// If you want a custom metric to record the exact number of CPU-seconds used by a job,
	// you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently
	// 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as
	// 12005.
	//
	// Alternatively, if you want a custom metric to record data in a more granular way, you
	// can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value
	// 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).
	// The supported units are a subset of The Unified Code for Units of Measure standard.
	// More info can be found in the API documentation
	// (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/monitoring_metric_descriptor#unit MonitoringMetricDescriptor#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

