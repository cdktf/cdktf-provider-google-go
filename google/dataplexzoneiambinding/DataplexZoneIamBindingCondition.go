package dataplexzoneiambinding


type DataplexZoneIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataplex_zone_iam_binding#expression DataplexZoneIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataplex_zone_iam_binding#title DataplexZoneIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataplex_zone_iam_binding#description DataplexZoneIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

