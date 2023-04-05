package cloudrunservice


type CloudRunServiceTemplateSpecContainersEnvFromSecretRefLocalObjectReference struct {
	// Name of the referent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#name CloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

