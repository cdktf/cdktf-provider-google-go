// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataprocAutoscalingPolicyBasicAlgorithm struct {
	// yarn_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_autoscaling_policy#yarn_config DataprocAutoscalingPolicy#yarn_config}
	YarnConfig *DataprocAutoscalingPolicyBasicAlgorithmYarnConfig `field:"required" json:"yarnConfig" yaml:"yarnConfig"`
	// Duration between scaling events. A scaling period starts after the update operation from the previous event has completed.
	//
	// Bounds: [2m, 1d]. Default: 2m.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_autoscaling_policy#cooldown_period DataprocAutoscalingPolicy#cooldown_period}
	CooldownPeriod *string `field:"optional" json:"cooldownPeriod" yaml:"cooldownPeriod"`
}

