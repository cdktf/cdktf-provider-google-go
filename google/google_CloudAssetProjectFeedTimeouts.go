// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type CloudAssetProjectFeedTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_project_feed#create CloudAssetProjectFeed#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_project_feed#delete CloudAssetProjectFeed#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_project_feed#update CloudAssetProjectFeed#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

