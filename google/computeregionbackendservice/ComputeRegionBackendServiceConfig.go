package computeregionbackendservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionBackendServiceConfig struct {
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
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#name ComputeRegionBackendService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE.
	//
	// If set to 0, the cookie is non-persistent and lasts
	// only until the end of the browser session (or equivalent). The
	// maximum allowed value for TTL is one day.
	//
	// When the load balancing scheme is INTERNAL, this field is not used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#affinity_cookie_ttl_sec ComputeRegionBackendService#affinity_cookie_ttl_sec}
	AffinityCookieTtlSec *float64 `field:"optional" json:"affinityCookieTtlSec" yaml:"affinityCookieTtlSec"`
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#backend ComputeRegionBackendService#backend}
	Backend interface{} `field:"optional" json:"backend" yaml:"backend"`
	// cdn_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#cdn_policy ComputeRegionBackendService#cdn_policy}
	CdnPolicy *ComputeRegionBackendServiceCdnPolicy `field:"optional" json:"cdnPolicy" yaml:"cdnPolicy"`
	// circuit_breakers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#circuit_breakers ComputeRegionBackendService#circuit_breakers}
	CircuitBreakers *ComputeRegionBackendServiceCircuitBreakers `field:"optional" json:"circuitBreakers" yaml:"circuitBreakers"`
	// Time for which instance will be drained (not accept new connections, but still work to finish started).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#connection_draining_timeout_sec ComputeRegionBackendService#connection_draining_timeout_sec}
	ConnectionDrainingTimeoutSec *float64 `field:"optional" json:"connectionDrainingTimeoutSec" yaml:"connectionDrainingTimeoutSec"`
	// consistent_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#consistent_hash ComputeRegionBackendService#consistent_hash}
	ConsistentHash *ComputeRegionBackendServiceConsistentHash `field:"optional" json:"consistentHash" yaml:"consistentHash"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#description ComputeRegionBackendService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If true, enable Cloud CDN for this RegionBackendService.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#enable_cdn ComputeRegionBackendService#enable_cdn}
	EnableCdn interface{} `field:"optional" json:"enableCdn" yaml:"enableCdn"`
	// failover_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#failover_policy ComputeRegionBackendService#failover_policy}
	FailoverPolicy *ComputeRegionBackendServiceFailoverPolicy `field:"optional" json:"failoverPolicy" yaml:"failoverPolicy"`
	// The set of URLs to HealthCheck resources for health checking this RegionBackendService. Currently at most one health check can be specified.
	//
	// A health check must be specified unless the backend service uses an internet
	// or serverless NEG as a backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#health_checks ComputeRegionBackendService#health_checks}
	HealthChecks *[]*string `field:"optional" json:"healthChecks" yaml:"healthChecks"`
	// iap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#iap ComputeRegionBackendService#iap}
	Iap *ComputeRegionBackendServiceIap `field:"optional" json:"iap" yaml:"iap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#id ComputeRegionBackendService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates what kind of load balancing this regional backend service will be used for.
	//
	// A backend service created for one type of load
	// balancing cannot be used with the other(s). For more information, refer to
	// [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service). Default value: "INTERNAL" Possible values: ["EXTERNAL", "EXTERNAL_MANAGED", "INTERNAL", "INTERNAL_MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#load_balancing_scheme ComputeRegionBackendService#load_balancing_scheme}
	LoadBalancingScheme *string `field:"optional" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// The load balancing algorithm used within the scope of the locality. The possible values are:.
	//
	// 'ROUND_ROBIN': This is a simple policy in which each healthy backend
	//              is selected in round robin order.
	//
	// 'LEAST_REQUEST': An O(1) algorithm which selects two random healthy
	//                hosts and picks the host which has fewer active requests.
	//
	// 'RING_HASH': The ring/modulo hash load balancer implements consistent
	//            hashing to backends. The algorithm has the property that the
	//            addition/removal of a host from a set of N hosts only affects
	//            1/N of the requests.
	//
	// 'RANDOM': The load balancer selects a random healthy host.
	//
	// 'ORIGINAL_DESTINATION': Backend host is selected based on the client
	//                       connection metadata, i.e., connections are opened
	//                       to the same address as the destination address of
	//                       the incoming connection before the connection
	//                       was redirected to the load balancer.
	//
	// 'MAGLEV': used as a drop in replacement for the ring hash load balancer.
	//         Maglev is not as stable as ring hash but has faster table lookup
	//         build times and host selection times. For more information about
	//         Maglev, refer to https://ai.google/research/pubs/pub44824
	//
	// 'WEIGHTED_MAGLEV': Per-instance weighted Load Balancing via health check
	//                  reported weights. If set, the Backend Service must
	//                  configure a non legacy HTTP-based Health Check, and
	//                  health check replies are expected to contain
	//                  non-standard HTTP response header field
	//                  X-Load-Balancing-Endpoint-Weight to specify the
	//                  per-instance weights. If set, Load Balancing is weight
	//                  based on the per-instance weights reported in the last
	//                  processed health check replies, as long as every
	//                  instance either reported a valid weight or had
	//                  UNAVAILABLE_WEIGHT. Otherwise, Load Balancing remains
	//                  equal-weight.
	//
	//
	// This field is applicable to either:
	//
	// A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2,
	// and loadBalancingScheme set to INTERNAL_MANAGED.
	// A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
	// A regional backend service with loadBalancingScheme set to EXTERNAL (External Network
	// Load Balancing). Only MAGLEV and WEIGHTED_MAGLEV values are possible for External
	// Network Load Balancing. The default is MAGLEV.
	//
	//
	// If session_affinity is not NONE, and this field is not set to MAGLEV, WEIGHTED_MAGLEV,
	// or RING_HASH, session affinity settings will not take effect.
	//
	// Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced
	// by a URL map that is bound to target gRPC proxy that has validate_for_proxyless
	// field set to true. Possible values: ["ROUND_ROBIN", "LEAST_REQUEST", "RING_HASH", "RANDOM", "ORIGINAL_DESTINATION", "MAGLEV", "WEIGHTED_MAGLEV"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#locality_lb_policy ComputeRegionBackendService#locality_lb_policy}
	LocalityLbPolicy *string `field:"optional" json:"localityLbPolicy" yaml:"localityLbPolicy"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#log_config ComputeRegionBackendService#log_config}
	LogConfig *ComputeRegionBackendServiceLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// The URL of the network to which this backend service belongs.
	//
	// This field can only be specified when the load balancing scheme is set to INTERNAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#network ComputeRegionBackendService#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// outlier_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#outlier_detection ComputeRegionBackendService#outlier_detection}
	OutlierDetection *ComputeRegionBackendServiceOutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// A named port on a backend instance group representing the port for communication to the backend VMs in that group.
	//
	// Required when the
	// loadBalancingScheme is EXTERNAL, EXTERNAL_MANAGED, INTERNAL_MANAGED, or INTERNAL_SELF_MANAGED
	// and the backends are instance groups. The named port must be defined on each
	// backend instance group. This parameter has no meaning if the backends are NEGs. API sets a
	// default of "http" if not given.
	// Must be omitted when the loadBalancingScheme is INTERNAL (Internal TCP/UDP Load Balancing).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#port_name ComputeRegionBackendService#port_name}
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#project ComputeRegionBackendService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The protocol this RegionBackendService uses to communicate with backends.
	//
	// The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
	// types and may result in errors if used with the GA API. Possible values: ["HTTP", "HTTPS", "HTTP2", "SSL", "TCP", "UDP", "GRPC", "UNSPECIFIED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#protocol ComputeRegionBackendService#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The Region in which the created backend service should reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#region ComputeRegionBackendService#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Type of session affinity to use.
	//
	// The default is NONE. Session affinity is
	// not applicable if the protocol is UDP. Possible values: ["NONE", "CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE", "CLIENT_IP_NO_DESTINATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#session_affinity ComputeRegionBackendService#session_affinity}
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#timeouts ComputeRegionBackendService#timeouts}
	Timeouts *ComputeRegionBackendServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// How many seconds to wait for the backend before considering it a failed request.
	//
	// Default is 30 seconds. Valid range is [1, 86400].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_backend_service#timeout_sec ComputeRegionBackendService#timeout_sec}
	TimeoutSec *float64 `field:"optional" json:"timeoutSec" yaml:"timeoutSec"`
}

