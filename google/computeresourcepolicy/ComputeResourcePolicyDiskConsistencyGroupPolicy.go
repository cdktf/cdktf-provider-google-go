package computeresourcepolicy


type ComputeResourcePolicyDiskConsistencyGroupPolicy struct {
	// Enable disk consistency on the resource policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_resource_policy#enabled ComputeResourcePolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

