package dataplexasset


type DataplexAssetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#create DataplexAsset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#delete DataplexAsset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#update DataplexAsset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

