package appenginestandardappversion


type AppEngineStandardAppVersionDeploymentFiles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_standard_app_version#name AppEngineStandardAppVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Source URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_standard_app_version#source_url AppEngineStandardAppVersion#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// SHA1 checksum of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_standard_app_version#sha1_sum AppEngineStandardAppVersion#sha1_sum}
	Sha1Sum *string `field:"optional" json:"sha1Sum" yaml:"sha1Sum"`
}

