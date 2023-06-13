package appengineflexibleappversion


type AppEngineFlexibleAppVersionHandlersScript struct {
	// Path to the script from the application root directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_flexible_app_version#script_path AppEngineFlexibleAppVersion#script_path}
	ScriptPath *string `field:"required" json:"scriptPath" yaml:"scriptPath"`
}

