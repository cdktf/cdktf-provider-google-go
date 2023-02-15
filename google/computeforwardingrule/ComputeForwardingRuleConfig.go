package computeforwardingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeForwardingRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Name of the resource;
	//
	// provided by the client when the resource is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#name ComputeForwardingRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// This field is used along with the `backend_service` field for internal load balancing or with the `target` field for internal TargetInstance.
	//
	// If the field is set to `TRUE`, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#allow_global_access ComputeForwardingRule#allow_global_access}
	AllowGlobalAccess interface{} `field:"optional" json:"allowGlobalAccess" yaml:"allowGlobalAccess"`
	// This field is used along with the `backend_service` field for internal load balancing or with the `target` field for internal TargetInstance.
	//
	// This field cannot be used with `port` or `portRange` fields. When the load balancing scheme is `INTERNAL` and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#all_ports ComputeForwardingRule#all_ports}
	AllPorts interface{} `field:"optional" json:"allPorts" yaml:"allPorts"`
	// This field is only used for `INTERNAL` load balancing.
	//
	// For internal load balancing, this field identifies the BackendService resource to receive the matched traffic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#backend_service ComputeForwardingRule#backend_service}
	BackendService *string `field:"optional" json:"backendService" yaml:"backendService"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#description ComputeForwardingRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#id ComputeForwardingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IP address that this forwarding rule serves.
	//
	// When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule. If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address: * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in `https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name` * Partial URL or by name, as in: * `projects/project_id/regions/region/addresses/address-name` * `regions/region/addresses/address-name` * `global/addresses/address-name` * `address-name` The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#ip_address ComputeForwardingRule#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The IP protocol to which this rule applies.
	//
	// For protocol forwarding, valid options are `TCP`, `UDP`, `ESP`, `AH`, `SCTP` or `ICMP`. For Internal TCP/UDP Load Balancing, the load balancing scheme is `INTERNAL`, and one of `TCP` or `UDP` are valid. For Traffic Director, the load balancing scheme is `INTERNAL_SELF_MANAGED`, and only `TCP`is valid. For Internal HTTP(S) Load Balancing, the load balancing scheme is `INTERNAL_MANAGED`, and only `TCP` is valid. For HTTP(S), SSL Proxy, and TCP Proxy Load Balancing, the load balancing scheme is `EXTERNAL` and only `TCP` is valid. For Network TCP/UDP Load Balancing, the load balancing scheme is `EXTERNAL`, and one of `TCP` or `UDP` is valid.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#ip_protocol ComputeForwardingRule#ip_protocol}
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring.
	//
	// To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a `PacketMirroring` rule applies to them. This can only be set to true for load balancers that have their `loadBalancingScheme` set to `INTERNAL`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#is_mirroring_collector ComputeForwardingRule#is_mirroring_collector}
	IsMirroringCollector interface{} `field:"optional" json:"isMirroringCollector" yaml:"isMirroringCollector"`
	// Labels to apply to this rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#labels ComputeForwardingRule#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Specifies the forwarding rule type.
	//
	// `EXTERNAL` is used for:
	//   Classic Cloud VPN gateways
	//   Protocol forwarding to VMs from an external IP address
	//   The following load balancers: HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP
	// `INTERNAL` is used for:
	//   Protocol forwarding to VMs from an internal IP address
	//   Internal TCP/UDP load balancers
	// `INTERNAL_MANAGED` is used for:
	//   Internal HTTP(S) load balancers
	// `INTERNAL_SELF_MANAGED` is used for:
	//   Traffic Director
	// `EXTERNAL_MANAGED` is used for:
	//   Global external HTTP(S) load balancers
	//
	// For more information about forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts). Possible values: INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED, EXTERNAL, EXTERNAL_MANAGED
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#load_balancing_scheme ComputeForwardingRule#load_balancing_scheme}
	LoadBalancingScheme *string `field:"optional" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// This field is not used for external load balancing.
	//
	// For `INTERNAL` and `INTERNAL_SELF_MANAGED` load balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#network ComputeForwardingRule#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// This signifies the networking tier used for configuring this load balancer and can only take the following values: `PREMIUM`, `STANDARD`.
	//
	// For regional ForwardingRule, the valid values are `PREMIUM` and `STANDARD`. For GlobalForwardingRule, the valid value is `PREMIUM`. If this field is not specified, it is assumed to be `PREMIUM`. If `IPAddress` is specified, this value must be equal to the networkTier of the Address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#network_tier ComputeForwardingRule#network_tier}
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// When the load balancing scheme is `EXTERNAL`, `INTERNAL_SELF_MANAGED` and `INTERNAL_MANAGED`, you can specify a `port_range`.
	//
	// Use with a forwarding rule that points to a target proxy or a target pool. Do not use with a forwarding rule that points to a backend service. This field is used along with the `target` field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when `IPProtocol` is `TCP`, `UDP`, or `SCTP`, only packets addressed to ports in the specified range will be forwarded to `target`. Forwarding rules with the same `[IPAddress, IPProtocol]` pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable ports:
	//
	// TargetHttpProxy: 80, 8080
	// TargetHttpsProxy: 443
	// TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	// TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	// TargetVpnGateway: 500, 4500.
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
	// This field is used along with the `backend_service` field for internal load balancing.
	//
	// When the load balancing scheme is `INTERNAL`, a list of ports can be configured, for example, ['80'], ['8000','9000']. Only packets addressed to these ports are forwarded to the backends configured with the forwarding rule. If the forwarding rule's loadBalancingScheme is INTERNAL, you can specify ports in one of the following ways: * A list of up to five ports, which can be non-contiguous * Keyword `ALL`, which causes the forwarding rule to forward traffic on any port of the forwarding rule's protocol. @pattern: d+(?:-d+)? For more information, refer to [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#ports ComputeForwardingRule#ports}
	Ports *[]*string `field:"optional" json:"ports" yaml:"ports"`
	// The project this resource belongs in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#project ComputeForwardingRule#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The location of this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#region ComputeForwardingRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// service_directory_registrations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#service_directory_registrations ComputeForwardingRule#service_directory_registrations}
	ServiceDirectoryRegistrations interface{} `field:"optional" json:"serviceDirectoryRegistrations" yaml:"serviceDirectoryRegistrations"`
	// An optional prefix to the service name for this Forwarding Rule.
	//
	// If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#service_label ComputeForwardingRule#service_label}
	ServiceLabel *string `field:"optional" json:"serviceLabel" yaml:"serviceLabel"`
	// This field is only used for `INTERNAL` load balancing.
	//
	// For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule. If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#subnetwork ComputeForwardingRule#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The URL of the target resource to receive the matched traffic.
	//
	// For regional forwarding rules, this target must live in the same region as the forwarding rule. For global forwarding rules, this target must be a global load balancing resource. The forwarded traffic must be of a type appropriate to the target object. For `INTERNAL_SELF_MANAGED` load balancing, only `targetHttpProxy` is valid, not `targetHttpsProxy`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#target ComputeForwardingRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#timeouts ComputeForwardingRule#timeouts}
	Timeouts *ComputeForwardingRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

