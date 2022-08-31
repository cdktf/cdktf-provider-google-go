// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataprocMetastoreServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_metastore_service_iam_member#expression DataprocMetastoreServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_metastore_service_iam_member#title DataprocMetastoreServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_metastore_service_iam_member#description DataprocMetastoreServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

