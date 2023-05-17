package cloudrunservice


type CloudRunServiceTemplateSpecContainersEnv struct {
	// Name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Defaults to "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#value CloudRunService#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#value_from CloudRunService#value_from}
	ValueFrom *CloudRunServiceTemplateSpecContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

