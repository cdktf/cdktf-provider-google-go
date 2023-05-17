package appengineflexibleappversion


type AppEngineFlexibleAppVersionDeployment struct {
	// cloud_build_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#cloud_build_options AppEngineFlexibleAppVersion#cloud_build_options}
	CloudBuildOptions *AppEngineFlexibleAppVersionDeploymentCloudBuildOptions `field:"optional" json:"cloudBuildOptions" yaml:"cloudBuildOptions"`
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#container AppEngineFlexibleAppVersion#container}
	Container *AppEngineFlexibleAppVersionDeploymentContainer `field:"optional" json:"container" yaml:"container"`
	// files block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#files AppEngineFlexibleAppVersion#files}
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// zip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_flexible_app_version#zip AppEngineFlexibleAppVersion#zip}
	Zip *AppEngineFlexibleAppVersionDeploymentZip `field:"optional" json:"zip" yaml:"zip"`
}

