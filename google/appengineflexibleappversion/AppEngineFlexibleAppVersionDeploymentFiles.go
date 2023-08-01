package appengineflexibleappversion


type AppEngineFlexibleAppVersionDeploymentFiles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_flexible_app_version#name AppEngineFlexibleAppVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Source URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_flexible_app_version#source_url AppEngineFlexibleAppVersion#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// SHA1 checksum of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_flexible_app_version#sha1_sum AppEngineFlexibleAppVersion#sha1_sum}
	Sha1Sum *string `field:"optional" json:"sha1Sum" yaml:"sha1Sum"`
}

