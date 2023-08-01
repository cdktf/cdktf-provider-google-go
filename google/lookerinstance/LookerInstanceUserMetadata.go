package lookerinstance


type LookerInstanceUserMetadata struct {
	// Number of additional Developer Users to allocate to the Looker Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/looker_instance#additional_developer_user_count LookerInstance#additional_developer_user_count}
	AdditionalDeveloperUserCount *float64 `field:"optional" json:"additionalDeveloperUserCount" yaml:"additionalDeveloperUserCount"`
	// Number of additional Standard Users to allocate to the Looker Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/looker_instance#additional_standard_user_count LookerInstance#additional_standard_user_count}
	AdditionalStandardUserCount *float64 `field:"optional" json:"additionalStandardUserCount" yaml:"additionalStandardUserCount"`
	// Number of additional Viewer Users to allocate to the Looker Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/looker_instance#additional_viewer_user_count LookerInstance#additional_viewer_user_count}
	AdditionalViewerUserCount *float64 `field:"optional" json:"additionalViewerUserCount" yaml:"additionalViewerUserCount"`
}

