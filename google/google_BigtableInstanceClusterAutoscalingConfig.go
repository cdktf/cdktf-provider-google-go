// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BigtableInstanceClusterAutoscalingConfig struct {
	// The target CPU utilization for autoscaling. Value must be between 10 and 80.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance#cpu_target BigtableInstance#cpu_target}
	CpuTarget *float64 `field:"required" json:"cpuTarget" yaml:"cpuTarget"`
	// The maximum number of nodes for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance#max_nodes BigtableInstance#max_nodes}
	MaxNodes *float64 `field:"required" json:"maxNodes" yaml:"maxNodes"`
	// The minimum number of nodes for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_instance#min_nodes BigtableInstance#min_nodes}
	MinNodes *float64 `field:"required" json:"minNodes" yaml:"minNodes"`
}

