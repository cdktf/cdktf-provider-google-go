package appengineflexibleappversion


type AppEngineFlexibleAppVersionDeploymentZip struct {
	// Source URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#source_url AppEngineFlexibleAppVersion#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// files count.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#files_count AppEngineFlexibleAppVersion#files_count}
	FilesCount *float64 `field:"optional" json:"filesCount" yaml:"filesCount"`
}

