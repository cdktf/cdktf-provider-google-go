package osconfigpatchdeployment


type OsConfigPatchDeploymentInstanceFilterGroupLabels struct {
	// Compute Engine instance labels that must be present for a VM instance to be targeted by this filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/os_config_patch_deployment#labels OsConfigPatchDeployment#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
}

