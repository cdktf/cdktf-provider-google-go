package computeregionbackendservice


type ComputeRegionBackendServiceBackend struct {
	// The fully-qualified URL of an Instance Group or Network Endpoint Group resource.
	//
	// In case of instance group this defines the list
	// of instances that serve traffic. Member virtual machine
	// instances from each instance group must live in the same zone as
	// the instance group itself. No two backends in a backend service
	// are allowed to use same Instance Group resource.
	//
	// For Network Endpoint Groups this defines list of endpoints. All
	// endpoints of Network Endpoint Group must be hosted on instances
	// located in the same zone as the Network Endpoint Group.
	//
	// Backend services cannot mix Instance Group and
	// Network Endpoint Group backends.
	//
	// When the 'load_balancing_scheme' is INTERNAL, only instance groups
	// are supported.
	//
	// Note that you must specify an Instance Group or Network Endpoint
	// Group resource using the fully-qualified URL, rather than a
	// partial URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#group ComputeRegionBackendService#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Specifies the balancing mode for this backend.
	//
	// See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
	// for an explanation of load balancing modes. Default value: "CONNECTION" Possible values: ["UTILIZATION", "RATE", "CONNECTION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#balancing_mode ComputeRegionBackendService#balancing_mode}
	BalancingMode *string `field:"optional" json:"balancingMode" yaml:"balancingMode"`
	// A multiplier applied to the group's maximum servicing capacity (based on UTILIZATION, RATE or CONNECTION).
	//
	// ~>**NOTE**: This field cannot be set for
	// INTERNAL region backend services (default loadBalancingScheme),
	// but is required for non-INTERNAL backend service. The total
	// capacity_scaler for all backends must be non-zero.
	//
	// A setting of 0 means the group is completely drained, offering
	// 0% of its available Capacity. Valid range is [0.0,1.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#capacity_scaler ComputeRegionBackendService#capacity_scaler}
	CapacityScaler *float64 `field:"optional" json:"capacityScaler" yaml:"capacityScaler"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#description ComputeRegionBackendService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// This field designates whether this is a failover backend.
	//
	// More
	// than one failover backend can be configured for a given RegionBackendService.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#failover ComputeRegionBackendService#failover}
	Failover interface{} `field:"optional" json:"failover" yaml:"failover"`
	// The max number of simultaneous connections for the group.
	//
	// Can
	// be used with either CONNECTION or UTILIZATION balancing modes.
	// Cannot be set for INTERNAL backend services.
	//
	// For CONNECTION mode, either maxConnections or one
	// of maxConnectionsPerInstance or maxConnectionsPerEndpoint,
	// as appropriate for group type, must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#max_connections ComputeRegionBackendService#max_connections}
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// The max number of simultaneous connections that a single backend network endpoint can handle. Cannot be set for INTERNAL backend services.
	//
	// This is used to calculate the capacity of the group. Can be
	// used in either CONNECTION or UTILIZATION balancing modes. For
	// CONNECTION mode, either maxConnections or
	// maxConnectionsPerEndpoint must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#max_connections_per_endpoint ComputeRegionBackendService#max_connections_per_endpoint}
	MaxConnectionsPerEndpoint *float64 `field:"optional" json:"maxConnectionsPerEndpoint" yaml:"maxConnectionsPerEndpoint"`
	// The max number of simultaneous connections that a single backend instance can handle. Cannot be set for INTERNAL backend services.
	//
	// This is used to calculate the capacity of the group.
	// Can be used in either CONNECTION or UTILIZATION balancing modes.
	// For CONNECTION mode, either maxConnections or
	// maxConnectionsPerInstance must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#max_connections_per_instance ComputeRegionBackendService#max_connections_per_instance}
	MaxConnectionsPerInstance *float64 `field:"optional" json:"maxConnectionsPerInstance" yaml:"maxConnectionsPerInstance"`
	// The max requests per second (RPS) of the group. Cannot be set for INTERNAL backend services.
	//
	// Can be used with either RATE or UTILIZATION balancing modes,
	// but required if RATE mode. Either maxRate or one
	// of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
	// group type, must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#max_rate ComputeRegionBackendService#max_rate}
	MaxRate *float64 `field:"optional" json:"maxRate" yaml:"maxRate"`
	// The max requests per second (RPS) that a single backend network endpoint can handle.
	//
	// This is used to calculate the capacity of
	// the group. Can be used in either balancing mode. For RATE mode,
	// either maxRate or maxRatePerEndpoint must be set. Cannot be set
	// for INTERNAL backend services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#max_rate_per_endpoint ComputeRegionBackendService#max_rate_per_endpoint}
	MaxRatePerEndpoint *float64 `field:"optional" json:"maxRatePerEndpoint" yaml:"maxRatePerEndpoint"`
	// The max requests per second (RPS) that a single backend instance can handle.
	//
	// This is used to calculate the capacity of
	// the group. Can be used in either balancing mode. For RATE mode,
	// either maxRate or maxRatePerInstance must be set. Cannot be set
	// for INTERNAL backend services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#max_rate_per_instance ComputeRegionBackendService#max_rate_per_instance}
	MaxRatePerInstance *float64 `field:"optional" json:"maxRatePerInstance" yaml:"maxRatePerInstance"`
	// Used when balancingMode is UTILIZATION.
	//
	// This ratio defines the
	// CPU utilization target for the group. Valid range is [0.0, 1.0].
	// Cannot be set for INTERNAL backend services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_backend_service#max_utilization ComputeRegionBackendService#max_utilization}
	MaxUtilization *float64 `field:"optional" json:"maxUtilization" yaml:"maxUtilization"`
}

