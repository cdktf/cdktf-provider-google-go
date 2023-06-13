package dataprocmetastoreserviceiambinding


type DataprocMetastoreServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_metastore_service_iam_binding#expression DataprocMetastoreServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_metastore_service_iam_binding#title DataprocMetastoreServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_metastore_service_iam_binding#description DataprocMetastoreServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

