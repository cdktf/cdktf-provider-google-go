package computebackendservice


type ComputeBackendServiceLocalityLbPolicies struct {
	// custom_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service#custom_policy ComputeBackendService#custom_policy}
	CustomPolicy *ComputeBackendServiceLocalityLbPoliciesCustomPolicy `field:"optional" json:"customPolicy" yaml:"customPolicy"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service#policy ComputeBackendService#policy}
	Policy *ComputeBackendServiceLocalityLbPoliciesPolicy `field:"optional" json:"policy" yaml:"policy"`
}
