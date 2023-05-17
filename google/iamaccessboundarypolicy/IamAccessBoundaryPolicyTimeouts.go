package iamaccessboundarypolicy


type IamAccessBoundaryPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/iam_access_boundary_policy#create IamAccessBoundaryPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/iam_access_boundary_policy#delete IamAccessBoundaryPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/iam_access_boundary_policy#update IamAccessBoundaryPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

