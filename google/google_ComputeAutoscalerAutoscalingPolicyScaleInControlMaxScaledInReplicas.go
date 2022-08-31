// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas struct {
	// Specifies a fixed number of VM instances. This must be a positive integer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#fixed ComputeAutoscaler#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies a percentage of instances between 0 to 100%, inclusive. For example, specify 80 for 80%.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#percent ComputeAutoscaler#percent}
	Percent *float64 `field:"optional" json:"percent" yaml:"percent"`
}

