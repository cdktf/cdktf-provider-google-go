package computeregionhealthcheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionHealthCheckConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#name ComputeRegionHealthCheck#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#check_interval_sec ComputeRegionHealthCheck#check_interval_sec}
	CheckIntervalSec *float64 `field:"optional" json:"checkIntervalSec" yaml:"checkIntervalSec"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#description ComputeRegionHealthCheck#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// grpc_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#grpc_health_check ComputeRegionHealthCheck#grpc_health_check}
	GrpcHealthCheck *ComputeRegionHealthCheckGrpcHealthCheck `field:"optional" json:"grpcHealthCheck" yaml:"grpcHealthCheck"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#healthy_threshold ComputeRegionHealthCheck#healthy_threshold}
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// http2_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#http2_health_check ComputeRegionHealthCheck#http2_health_check}
	Http2HealthCheck *ComputeRegionHealthCheckHttp2HealthCheck `field:"optional" json:"http2HealthCheck" yaml:"http2HealthCheck"`
	// http_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#http_health_check ComputeRegionHealthCheck#http_health_check}
	HttpHealthCheck *ComputeRegionHealthCheckHttpHealthCheck `field:"optional" json:"httpHealthCheck" yaml:"httpHealthCheck"`
	// https_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#https_health_check ComputeRegionHealthCheck#https_health_check}
	HttpsHealthCheck *ComputeRegionHealthCheckHttpsHealthCheck `field:"optional" json:"httpsHealthCheck" yaml:"httpsHealthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#id ComputeRegionHealthCheck#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#log_config ComputeRegionHealthCheck#log_config}
	LogConfig *ComputeRegionHealthCheckLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#project ComputeRegionHealthCheck#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Region in which the created health check should reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#region ComputeRegionHealthCheck#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ssl_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#ssl_health_check ComputeRegionHealthCheck#ssl_health_check}
	SslHealthCheck *ComputeRegionHealthCheckSslHealthCheck `field:"optional" json:"sslHealthCheck" yaml:"sslHealthCheck"`
	// tcp_health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#tcp_health_check ComputeRegionHealthCheck#tcp_health_check}
	TcpHealthCheck *ComputeRegionHealthCheckTcpHealthCheck `field:"optional" json:"tcpHealthCheck" yaml:"tcpHealthCheck"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#timeouts ComputeRegionHealthCheck#timeouts}
	Timeouts *ComputeRegionHealthCheckTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// How long (in seconds) to wait before claiming failure.
	//
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#timeout_sec ComputeRegionHealthCheck#timeout_sec}
	TimeoutSec *float64 `field:"optional" json:"timeoutSec" yaml:"timeoutSec"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_health_check#unhealthy_threshold ComputeRegionHealthCheck#unhealthy_threshold}
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

