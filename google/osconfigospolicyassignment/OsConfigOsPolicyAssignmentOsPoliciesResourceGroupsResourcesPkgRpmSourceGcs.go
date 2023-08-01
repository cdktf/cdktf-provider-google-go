package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgRpmSourceGcs struct {
	// Bucket of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_os_policy_assignment#bucket OsConfigOsPolicyAssignment#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Name of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_os_policy_assignment#object OsConfigOsPolicyAssignment#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Generation number of the Cloud Storage object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_os_policy_assignment#generation OsConfigOsPolicyAssignment#generation}
	Generation *float64 `field:"optional" json:"generation" yaml:"generation"`
}

