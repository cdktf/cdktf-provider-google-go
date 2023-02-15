package dataprocjobiambinding


type DataprocJobIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_job_iam_binding#expression DataprocJobIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_job_iam_binding#title DataprocJobIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_job_iam_binding#description DataprocJobIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

