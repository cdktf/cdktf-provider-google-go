package appenginestandardappversion


type AppEngineStandardAppVersionDeployment struct {
	// files block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_standard_app_version#files AppEngineStandardAppVersion#files}
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// zip block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_standard_app_version#zip AppEngineStandardAppVersion#zip}
	Zip *AppEngineStandardAppVersionDeploymentZip `field:"optional" json:"zip" yaml:"zip"`
}

