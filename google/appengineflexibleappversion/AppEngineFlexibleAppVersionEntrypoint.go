package appengineflexibleappversion


type AppEngineFlexibleAppVersionEntrypoint struct {
	// The format should be a shell command that can be fed to bash -c.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/app_engine_flexible_app_version#shell AppEngineFlexibleAppVersion#shell}
	Shell *string `field:"required" json:"shell" yaml:"shell"`
}

