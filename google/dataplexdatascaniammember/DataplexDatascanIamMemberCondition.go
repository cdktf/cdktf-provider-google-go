package dataplexdatascaniammember


type DataplexDatascanIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_datascan_iam_member#expression DataplexDatascanIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_datascan_iam_member#title DataplexDatascanIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_datascan_iam_member#description DataplexDatascanIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

