package dataplexzoneiambinding


type DataplexZoneIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_zone_iam_binding#expression DataplexZoneIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_zone_iam_binding#title DataplexZoneIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_zone_iam_binding#description DataplexZoneIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

