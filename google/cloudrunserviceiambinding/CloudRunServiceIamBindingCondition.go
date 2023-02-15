package cloudrunserviceiambinding


type CloudRunServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service_iam_binding#expression CloudRunServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service_iam_binding#title CloudRunServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service_iam_binding#description CloudRunServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

