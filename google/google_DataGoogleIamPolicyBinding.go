// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataGoogleIamPolicyBinding struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/iam_policy#members DataGoogleIamPolicy#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/iam_policy#role DataGoogleIamPolicy#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/iam_policy#condition DataGoogleIamPolicy#condition}
	Condition *DataGoogleIamPolicyBindingCondition `field:"optional" json:"condition" yaml:"condition"`
}

