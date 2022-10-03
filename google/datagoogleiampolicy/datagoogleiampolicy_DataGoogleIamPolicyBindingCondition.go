package datagoogleiampolicy


type DataGoogleIamPolicyBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/iam_policy#expression DataGoogleIamPolicy#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/iam_policy#title DataGoogleIamPolicy#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/iam_policy#description DataGoogleIamPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

