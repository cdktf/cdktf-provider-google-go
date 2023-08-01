package datagoogleiampolicy


type DataGoogleIamPolicyBinding struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/data-sources/iam_policy#members DataGoogleIamPolicy#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/data-sources/iam_policy#role DataGoogleIamPolicy#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/data-sources/iam_policy#condition DataGoogleIamPolicy#condition}
	Condition *DataGoogleIamPolicyBindingCondition `field:"optional" json:"condition" yaml:"condition"`
}

