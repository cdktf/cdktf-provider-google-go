package cloudrunservice


type CloudRunServiceTemplateSpecContainersStartupProbeGrpc struct {
	// Port number to access on the container.
	//
	// Number must be in the range 1 to 65535.
	// If not specified, defaults to the same value as container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#port CloudRunService#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#service CloudRunService#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

