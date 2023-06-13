package dataplexdatascaniambinding


type DataplexDatascanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan_iam_binding#expression DataplexDatascanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan_iam_binding#title DataplexDatascanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan_iam_binding#description DataplexDatascanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

