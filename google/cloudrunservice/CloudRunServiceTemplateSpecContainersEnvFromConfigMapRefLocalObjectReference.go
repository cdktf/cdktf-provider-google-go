package cloudrunservice


type CloudRunServiceTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference struct {
	// Name of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

