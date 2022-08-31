// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataprocMetastoreServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_metastore_service_iam_binding#expression DataprocMetastoreServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_metastore_service_iam_binding#title DataprocMetastoreServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_metastore_service_iam_binding#description DataprocMetastoreServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

