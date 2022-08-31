// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type CloudRunServiceTemplateSpecContainersEnvFromSecretRefLocalObjectReference struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#name CloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

