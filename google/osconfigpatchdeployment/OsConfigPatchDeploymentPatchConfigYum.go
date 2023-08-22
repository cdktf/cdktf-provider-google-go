package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfigYum struct {
	// List of packages to exclude from update. These packages will be excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_patch_deployment#excludes OsConfigPatchDeployment#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// An exclusive list of packages to be updated.
	//
	// These are the only packages that will be updated.
	// If these packages are not installed, they will be ignored. This field cannot be specified with
	// any other patch configuration fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_patch_deployment#exclusive_packages OsConfigPatchDeployment#exclusive_packages}
	ExclusivePackages *[]*string `field:"optional" json:"exclusivePackages" yaml:"exclusivePackages"`
	// Will cause patch to run yum update-minimal instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_patch_deployment#minimal OsConfigPatchDeployment#minimal}
	Minimal interface{} `field:"optional" json:"minimal" yaml:"minimal"`
	// Adds the --security flag to yum update. Not supported on all platforms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_patch_deployment#security OsConfigPatchDeployment#security}
	Security interface{} `field:"optional" json:"security" yaml:"security"`
}

