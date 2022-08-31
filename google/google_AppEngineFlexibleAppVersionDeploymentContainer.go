// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type AppEngineFlexibleAppVersionDeploymentContainer struct {
	// URI to the hosted container image in Google Container Registry.
	//
	// The URI must be fully qualified and include a tag or digest.
	// Examples: "gcr.io/my-project/image:tag" or "gcr.io/my-project/image@digest"
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#image AppEngineFlexibleAppVersion#image}
	Image *string `field:"required" json:"image" yaml:"image"`
}

