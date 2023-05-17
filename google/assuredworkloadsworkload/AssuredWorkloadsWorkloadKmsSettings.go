package assuredworkloadsworkload


type AssuredWorkloadsWorkloadKmsSettings struct {
	// Required.
	//
	// Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/assured_workloads_workload#next_rotation_time AssuredWorkloadsWorkload#next_rotation_time}
	NextRotationTime *string `field:"required" json:"nextRotationTime" yaml:"nextRotationTime"`
	// Required.
	//
	// Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/assured_workloads_workload#rotation_period AssuredWorkloadsWorkload#rotation_period}
	RotationPeriod *string `field:"required" json:"rotationPeriod" yaml:"rotationPeriod"`
}

