package computerouternat

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRouterNatConfig struct {
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
	// Name of the NAT service. The name must be 1-63 characters long and comply with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#name ComputeRouterNat#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// How external IPs should be allocated for this NAT.
	//
	// Valid values are
	// 'AUTO_ONLY' for only allowing NAT IPs allocated by Google Cloud
	// Platform, or 'MANUAL_ONLY' for only user-allocated NAT IP addresses. Possible values: ["MANUAL_ONLY", "AUTO_ONLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#nat_ip_allocate_option ComputeRouterNat#nat_ip_allocate_option}
	NatIpAllocateOption *string `field:"required" json:"natIpAllocateOption" yaml:"natIpAllocateOption"`
	// The name of the Cloud Router in which this NAT will be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#router ComputeRouterNat#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// How NAT should be configured per Subnetwork.
	//
	// If 'ALL_SUBNETWORKS_ALL_IP_RANGES', all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If 'ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES', all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// 'LIST_OF_SUBNETWORKS': A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region. Possible values: ["ALL_SUBNETWORKS_ALL_IP_RANGES", "ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES", "LIST_OF_SUBNETWORKS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#source_subnetwork_ip_ranges_to_nat ComputeRouterNat#source_subnetwork_ip_ranges_to_nat}
	SourceSubnetworkIpRangesToNat *string `field:"required" json:"sourceSubnetworkIpRangesToNat" yaml:"sourceSubnetworkIpRangesToNat"`
	// A list of URLs of the IP resources to be drained.
	//
	// These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#drain_nat_ips ComputeRouterNat#drain_nat_ips}
	DrainNatIps *[]*string `field:"optional" json:"drainNatIps" yaml:"drainNatIps"`
	// Enable Dynamic Port Allocation.
	//
	// If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
	// If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
	// If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
	// If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
	//
	// Mutually exclusive with enableEndpointIndependentMapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#enable_dynamic_port_allocation ComputeRouterNat#enable_dynamic_port_allocation}
	EnableDynamicPortAllocation interface{} `field:"optional" json:"enableDynamicPortAllocation" yaml:"enableDynamicPortAllocation"`
	// Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#enable_endpoint_independent_mapping ComputeRouterNat#enable_endpoint_independent_mapping}
	EnableEndpointIndependentMapping interface{} `field:"optional" json:"enableEndpointIndependentMapping" yaml:"enableEndpointIndependentMapping"`
	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#icmp_idle_timeout_sec ComputeRouterNat#icmp_idle_timeout_sec}
	IcmpIdleTimeoutSec *float64 `field:"optional" json:"icmpIdleTimeoutSec" yaml:"icmpIdleTimeoutSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#id ComputeRouterNat#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#log_config ComputeRouterNat#log_config}
	LogConfig *ComputeRouterNatLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// Maximum number of ports allocated to a VM from this NAT.
	//
	// This field can only be set when enableDynamicPortAllocation is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#max_ports_per_vm ComputeRouterNat#max_ports_per_vm}
	MaxPortsPerVm *float64 `field:"optional" json:"maxPortsPerVm" yaml:"maxPortsPerVm"`
	// Minimum number of ports allocated to a VM from this NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#min_ports_per_vm ComputeRouterNat#min_ports_per_vm}
	MinPortsPerVm *float64 `field:"optional" json:"minPortsPerVm" yaml:"minPortsPerVm"`
	// Self-links of NAT IPs. Only valid if natIpAllocateOption is set to MANUAL_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#nat_ips ComputeRouterNat#nat_ips}
	NatIps *[]*string `field:"optional" json:"natIps" yaml:"natIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#project ComputeRouterNat#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the router and NAT reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#region ComputeRouterNat#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#rules ComputeRouterNat#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// subnetwork block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#subnetwork ComputeRouterNat#subnetwork}
	Subnetwork interface{} `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// Timeout (in seconds) for TCP established connections. Defaults to 1200s if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#tcp_established_idle_timeout_sec ComputeRouterNat#tcp_established_idle_timeout_sec}
	TcpEstablishedIdleTimeoutSec *float64 `field:"optional" json:"tcpEstablishedIdleTimeoutSec" yaml:"tcpEstablishedIdleTimeoutSec"`
	// Timeout (in seconds) for TCP connections that are in TIME_WAIT state. Defaults to 120s if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#tcp_time_wait_timeout_sec ComputeRouterNat#tcp_time_wait_timeout_sec}
	TcpTimeWaitTimeoutSec *float64 `field:"optional" json:"tcpTimeWaitTimeoutSec" yaml:"tcpTimeWaitTimeoutSec"`
	// Timeout (in seconds) for TCP transitory connections. Defaults to 30s if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#tcp_transitory_idle_timeout_sec ComputeRouterNat#tcp_transitory_idle_timeout_sec}
	TcpTransitoryIdleTimeoutSec *float64 `field:"optional" json:"tcpTransitoryIdleTimeoutSec" yaml:"tcpTransitoryIdleTimeoutSec"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#timeouts ComputeRouterNat#timeouts}
	Timeouts *ComputeRouterNatTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_router_nat#udp_idle_timeout_sec ComputeRouterNat#udp_idle_timeout_sec}
	UdpIdleTimeoutSec *float64 `field:"optional" json:"udpIdleTimeoutSec" yaml:"udpIdleTimeoutSec"`
}

