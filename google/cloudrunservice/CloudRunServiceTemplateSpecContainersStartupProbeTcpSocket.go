package cloudrunservice


type CloudRunServiceTemplateSpecContainersStartupProbeTcpSocket struct {
	// Port number to access on the container.
	//
	// Number must be in the range 1 to 65535.
	// If not specified, defaults to the same value as container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_run_service#port CloudRunService#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

