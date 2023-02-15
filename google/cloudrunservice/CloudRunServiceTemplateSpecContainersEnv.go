package cloudrunservice


type CloudRunServiceTemplateSpecContainersEnv struct {
	// Name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#name CloudRunService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any route environment variables.
	//
	// If a variable cannot be resolved,
	// the reference in the input string will be unchanged. The $(VAR_NAME)
	// syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	// references will never be expanded, regardless of whether the variable
	// exists or not.
	// Defaults to "".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#value CloudRunService#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_from block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#value_from CloudRunService#value_from}
	ValueFrom *CloudRunServiceTemplateSpecContainersEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

