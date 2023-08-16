package appengineflexibleappversion


type AppEngineFlexibleAppVersionDeploymentContainer struct {
	// URI to the hosted container image in Google Container Registry.
	//
	// The URI must be fully qualified and include a tag or digest.
	// Examples: "gcr.io/my-project/image:tag" or "gcr.io/my-project/image@digest"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_flexible_app_version#image AppEngineFlexibleAppVersion#image}
	Image *string `field:"required" json:"image" yaml:"image"`
}

