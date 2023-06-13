package dataplexlakeiambinding


type DataplexLakeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_lake_iam_binding#expression DataplexLakeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_lake_iam_binding#title DataplexLakeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_lake_iam_binding#description DataplexLakeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

