package cloudrunservice


type CloudRunServiceTemplateSpecContainersPorts struct {
	// Port number the container listens on.
	//
	// This must be a valid port number (between 1 and 65535). Defaults to "8080".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_run_service#container_port CloudRunService#container_port}
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// If specified, used to specify which protocol to use.
	//
	// Allowed values are "http1" (HTTP/1) and "h2c" (HTTP/2 end-to-end). Defaults to "http1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Protocol for port. Must be "TCP". Defaults to "TCP".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_run_service#protocol CloudRunService#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

