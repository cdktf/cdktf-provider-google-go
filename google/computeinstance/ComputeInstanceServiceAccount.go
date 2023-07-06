package computeinstance


type ComputeInstanceServiceAccount struct {
	// A list of service scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/compute_instance#scopes ComputeInstance#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The service account e-mail address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/compute_instance#email ComputeInstance#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

