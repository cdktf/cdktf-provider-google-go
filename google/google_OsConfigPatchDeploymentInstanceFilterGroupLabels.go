// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type OsConfigPatchDeploymentInstanceFilterGroupLabels struct {
	// Compute Engine instance labels that must be present for a VM instance to be targeted by this filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#labels OsConfigPatchDeployment#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
}

