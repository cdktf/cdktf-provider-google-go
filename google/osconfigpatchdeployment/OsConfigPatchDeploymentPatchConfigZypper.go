package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfigZypper struct {
	// Install only patches with these categories. Common categories include security, recommended, and feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#categories OsConfigPatchDeployment#categories}
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// List of packages to exclude from update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#excludes OsConfigPatchDeployment#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// An exclusive list of patches to be updated.
	//
	// These are the only patches that will be installed using 'zypper patch patch:' command.
	// This field must not be used with any other patch configuration fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#exclusive_patches OsConfigPatchDeployment#exclusive_patches}
	ExclusivePatches *[]*string `field:"optional" json:"exclusivePatches" yaml:"exclusivePatches"`
	// Install only patches with these severities. Common severities include critical, important, moderate, and low.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#severities OsConfigPatchDeployment#severities}
	Severities *[]*string `field:"optional" json:"severities" yaml:"severities"`
	// Adds the --with-optional flag to zypper patch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#with_optional OsConfigPatchDeployment#with_optional}
	WithOptional interface{} `field:"optional" json:"withOptional" yaml:"withOptional"`
	// Adds the --with-update flag, to zypper patch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#with_update OsConfigPatchDeployment#with_update}
	WithUpdate interface{} `field:"optional" json:"withUpdate" yaml:"withUpdate"`
}

