package cloudrunv2serviceiammember


type CloudRunV2ServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_service_iam_member#expression CloudRunV2ServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_service_iam_member#title CloudRunV2ServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_service_iam_member#description CloudRunV2ServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

