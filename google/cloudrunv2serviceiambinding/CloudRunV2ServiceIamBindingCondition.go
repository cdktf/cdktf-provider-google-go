package cloudrunv2serviceiambinding


type CloudRunV2ServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloud_run_v2_service_iam_binding#expression CloudRunV2ServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloud_run_v2_service_iam_binding#title CloudRunV2ServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloud_run_v2_service_iam_binding#description CloudRunV2ServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

