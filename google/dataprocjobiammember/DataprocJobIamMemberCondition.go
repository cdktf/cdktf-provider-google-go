package dataprocjobiammember


type DataprocJobIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_job_iam_member#expression DataprocJobIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_job_iam_member#title DataprocJobIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_job_iam_member#description DataprocJobIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

