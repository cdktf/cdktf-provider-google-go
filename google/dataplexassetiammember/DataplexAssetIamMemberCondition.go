package dataplexassetiammember


type DataplexAssetIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset_iam_member#expression DataplexAssetIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset_iam_member#title DataplexAssetIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset_iam_member#description DataplexAssetIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

