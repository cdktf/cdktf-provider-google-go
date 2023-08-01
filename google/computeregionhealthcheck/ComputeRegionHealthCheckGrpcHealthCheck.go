package computeregionhealthcheck


type ComputeRegionHealthCheckGrpcHealthCheck struct {
	// The gRPC service name for the health check. The value of grpcServiceName has the following meanings by convention:.
	//
	// Empty serviceName means the overall status of all services at the backend.
	// Non-empty serviceName means the health of that gRPC service, as defined by the owner of the service.
	//
	// The grpcServiceName can only be ASCII.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#grpc_service_name ComputeRegionHealthCheck#grpc_service_name}
	GrpcServiceName *string `field:"optional" json:"grpcServiceName" yaml:"grpcServiceName"`
	// The port number for the health check request.
	//
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#port ComputeRegionHealthCheck#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#port_name ComputeRegionHealthCheck#port_name}
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
	// Specifies how port is selected for health checking, can be one of the following values:.
	//
	// 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//
	// 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//
	// 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	// network endpoint is used for health checking. For other backends, the
	// port or named port specified in the Backend Service is used for health
	// checking.
	//
	// If not specified, gRPC health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#port_specification ComputeRegionHealthCheck#port_specification}
	PortSpecification *string `field:"optional" json:"portSpecification" yaml:"portSpecification"`
}

