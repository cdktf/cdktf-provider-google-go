package cloudrunv2jobiambinding


type CloudRunV2JobIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_job_iam_binding#expression CloudRunV2JobIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_job_iam_binding#title CloudRunV2JobIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_job_iam_binding#description CloudRunV2JobIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

