package osconfigpatchdeployment


type OsConfigPatchDeploymentPatchConfigApt struct {
	// List of packages to exclude from update. These packages will be excluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#excludes OsConfigPatchDeployment#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// An exclusive list of packages to be updated.
	//
	// These are the only packages that will be updated.
	// If these packages are not installed, they will be ignored. This field cannot be specified with
	// any other patch configuration fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#exclusive_packages OsConfigPatchDeployment#exclusive_packages}
	ExclusivePackages *[]*string `field:"optional" json:"exclusivePackages" yaml:"exclusivePackages"`
	// By changing the type to DIST, the patching is performed using apt-get dist-upgrade instead. Possible values: ["DIST", "UPGRADE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#type OsConfigPatchDeployment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

