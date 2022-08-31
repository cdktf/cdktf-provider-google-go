// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type CloudAssetFolderFeedTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_folder_feed#create CloudAssetFolderFeed#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_folder_feed#delete CloudAssetFolderFeed#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_folder_feed#update CloudAssetFolderFeed#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

