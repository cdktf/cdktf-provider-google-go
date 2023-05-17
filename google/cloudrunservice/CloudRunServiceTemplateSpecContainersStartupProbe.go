package cloudrunservice


type CloudRunServiceTemplateSpecContainersStartupProbe struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#failure_threshold CloudRunService#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#grpc CloudRunService#grpc}
	Grpc *CloudRunServiceTemplateSpecContainersStartupProbeGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#http_get CloudRunService#http_get}
	HttpGet *CloudRunServiceTemplateSpecContainersStartupProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Number of seconds after the container has started before the probe is initiated.
	//
	// Defaults to 0 seconds. Minimum value is 0. Maximum value is 240.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#initial_delay_seconds CloudRunService#initial_delay_seconds}
	InitialDelaySeconds *float64 `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. Maximum value is 240.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#period_seconds CloudRunService#period_seconds}
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#tcp_socket CloudRunService#tcp_socket}
	TcpSocket *CloudRunServiceTemplateSpecContainersStartupProbeTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1. Maximum value is 3600.
	// Must be smaller than periodSeconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_run_service#timeout_seconds CloudRunService#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

