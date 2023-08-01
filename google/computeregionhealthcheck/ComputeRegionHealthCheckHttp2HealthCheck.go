package computeregionhealthcheck


type ComputeRegionHealthCheckHttp2HealthCheck struct {
	// The value of the host header in the HTTP2 health check request.
	//
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#host ComputeRegionHealthCheck#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The TCP port number for the HTTP2 health check request. The default value is 443.
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
	// If not specified, HTTP2 health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#port_specification ComputeRegionHealthCheck#port_specification}
	PortSpecification *string `field:"optional" json:"portSpecification" yaml:"portSpecification"`
	// Specifies the type of proxy header to append before sending data to the backend.
	//
	// Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#proxy_header ComputeRegionHealthCheck#proxy_header}
	ProxyHeader *string `field:"optional" json:"proxyHeader" yaml:"proxyHeader"`
	// The request path of the HTTP2 health check request. The default value is /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#request_path ComputeRegionHealthCheck#request_path}
	RequestPath *string `field:"optional" json:"requestPath" yaml:"requestPath"`
	// The bytes to match against the beginning of the response data.
	//
	// If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_health_check#response ComputeRegionHealthCheck#response}
	Response *string `field:"optional" json:"response" yaml:"response"`
}

