package appenginestandardappversion


type AppEngineStandardAppVersionHandlersScript struct {
	// Path to the script from the application root directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_standard_app_version#script_path AppEngineStandardAppVersion#script_path}
	ScriptPath *string `field:"required" json:"scriptPath" yaml:"scriptPath"`
}

