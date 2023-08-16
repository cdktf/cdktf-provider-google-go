package computeforwardingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeForwardingRuleConfig struct {
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
	// Name of the resource;
	//
	// provided by the client when the resource is created.
	// The name must be 1-63 characters long, and comply with
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
	//
	// Specifically, the name must be 1-63 characters long and match the regular
	// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	// character must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit, except the last character, which
	// cannot be a dash.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, the forwarding rule name must be a 1-20 characters string with
	// lowercase letters and numbers and must start with a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#name ComputeForwardingRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// This field is used along with the 'backend_service' field for internal load balancing or with the 'target' field for internal TargetInstance.
	//
	// If the field is set to 'TRUE', clients can access ILB from all
	// regions.
	//
	// Otherwise only allows access from clients in the same region as the
	// internal load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#allow_global_access ComputeForwardingRule#allow_global_access}
	AllowGlobalAccess interface{} `field:"optional" json:"allowGlobalAccess" yaml:"allowGlobalAccess"`
	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#allow_psc_global_access ComputeForwardingRule#allow_psc_global_access}
	AllowPscGlobalAccess interface{} `field:"optional" json:"allowPscGlobalAccess" yaml:"allowPscGlobalAccess"`
	// This field can only be used: If 'IPProtocol' is one of TCP, UDP, or SCTP.
	//
	// By internal TCP/UDP load balancers, backend service-based network load
	// balancers, and internal and external protocol forwarding.
	//
	// This option should be set to TRUE when the Forwarding Rule
	// IPProtocol is set to L3_DEFAULT.
	//
	// Set this field to true to allow packets addressed to any port or packets
	// lacking destination port information (for example, UDP fragments after the
	// first fragment) to be forwarded to the backends configured with this
	// forwarding rule.
	//
	// The 'ports', 'port_range', and
	// 'allPorts' fields are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#all_ports ComputeForwardingRule#all_ports}
	AllPorts interface{} `field:"optional" json:"allPorts" yaml:"allPorts"`
	// Identifies the backend service to which the forwarding rule sends traffic.
	//
	// Required for Internal TCP/UDP Load Balancing and Network Load Balancing;
	// must be omitted for all other load balancer types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#backend_service ComputeForwardingRule#backend_service}
	BackendService *string `field:"optional" json:"backendService" yaml:"backendService"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#description ComputeForwardingRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#id ComputeForwardingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IP address for which this forwarding rule accepts traffic.
	//
	// When a client
	// sends traffic to this IP address, the forwarding rule directs the traffic
	// to the referenced 'target' or 'backendService'.
	//
	// While creating a forwarding rule, specifying an 'IPAddress' is
	// required under the following circumstances:
	//
	// When the 'target' is set to 'targetGrpcProxy' and
	// 'validateForProxyless' is set to 'true', the
	// 'IPAddress' should be set to '0.0.0.0'.
	// When the 'target' is a Private Service Connect Google APIs
	// bundle, you must specify an 'IPAddress'.
	//
	//
	// Otherwise, you can optionally specify an IP address that references an
	// existing static (reserved) IP address resource. When omitted, Google Cloud
	// assigns an ephemeral IP address.
	//
	// Use one of the following formats to specify an IP address while creating a
	// forwarding rule:
	//
	// IP address number, as in '100.1.2.3'
	// IPv6 address range, as in '2600:1234::/96'
	// Full resource URL, as in
	// 'https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name'
	// Partial URL or by name, as in:
	// 'projects/project_id/regions/region/addresses/address-name'
	// 'regions/region/addresses/address-name'
	// 'global/addresses/address-name'
	// 'address-name'
	//
	//
	// The forwarding rule's 'target' or 'backendService',
	// and in most cases, also the 'loadBalancingScheme', determine the
	// type of IP address that you can use. For detailed information, see
	// [IP address
	// specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	//
	// When reading an 'IPAddress', the API always returns the IP
	// address number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#ip_address ComputeForwardingRule#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The IP protocol to which this rule applies.
	//
	// For protocol forwarding, valid
	// options are 'TCP', 'UDP', 'ESP',
	// 'AH', 'SCTP', 'ICMP' and
	// 'L3_DEFAULT'.
	//
	// The valid IP protocols are different for different load balancing products
	// as described in [Load balancing
	// features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
	//
	// A Forwarding Rule with protocol L3_DEFAULT can attach with target instance or
	// backend service with UNSPECIFIED protocol.
	// A forwarding rule with "L3_DEFAULT" IPProtocal cannot be attached to a backend service with TCP or UDP. Possible values: ["TCP", "UDP", "ESP", "AH", "SCTP", "ICMP", "L3_DEFAULT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#ip_protocol ComputeForwardingRule#ip_protocol}
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// The IP address version that will be used by this forwarding rule. Valid options are IPV4 and IPV6.
	//
	// If not set, the IPv4 address will be used by default. Possible values: ["IPV4", "IPV6"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#ip_version ComputeForwardingRule#ip_version}
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring.
	//
	// To prevent mirroring loops, instances behind this
	// load balancer will not have their traffic mirrored even if a
	// 'PacketMirroring' rule applies to them.
	//
	// This can only be set to true for load balancers that have their
	// 'loadBalancingScheme' set to 'INTERNAL'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#is_mirroring_collector ComputeForwardingRule#is_mirroring_collector}
	IsMirroringCollector interface{} `field:"optional" json:"isMirroringCollector" yaml:"isMirroringCollector"`
	// Labels to apply to this forwarding rule.  A list of key->value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#labels ComputeForwardingRule#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Specifies the forwarding rule type.
	//
	// For more information about forwarding rules, refer to
	// [Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts). Default value: "EXTERNAL" Possible values: ["EXTERNAL", "EXTERNAL_MANAGED", "INTERNAL", "INTERNAL_MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#load_balancing_scheme ComputeForwardingRule#load_balancing_scheme}
	LoadBalancingScheme *string `field:"optional" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// This field is not used for external load balancing.
	//
	// For Internal TCP/UDP Load Balancing, this field identifies the network that
	// the load balanced IP should belong to for this Forwarding Rule.
	// If the subnetwork is specified, the network of the subnetwork will be used.
	// If neither subnetwork nor this field is specified, the default network will
	// be used.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google
	// APIs, a network must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#network ComputeForwardingRule#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// This signifies the networking tier used for configuring this load balancer and can only take the following values: 'PREMIUM', 'STANDARD'.
	//
	// For regional ForwardingRule, the valid values are 'PREMIUM' and
	// 'STANDARD'. For GlobalForwardingRule, the valid value is
	// 'PREMIUM'.
	//
	// If this field is not specified, it is assumed to be 'PREMIUM'.
	// If 'IPAddress' is specified, this value must be equal to the
	// networkTier of the Address. Possible values: ["PREMIUM", "STANDARD"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#network_tier ComputeForwardingRule#network_tier}
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// This is used in PSC consumer ForwardingRule to control whether it should try to auto-generate a DNS zone or not.
	//
	// Non-PSC forwarding rules do not use this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#no_automate_dns_zone ComputeForwardingRule#no_automate_dns_zone}
	NoAutomateDnsZone interface{} `field:"optional" json:"noAutomateDnsZone" yaml:"noAutomateDnsZone"`
	// This field can only be used:.
	//
	// If 'IPProtocol' is one of TCP, UDP, or SCTP.
	// By backend service-based network load balancers, target pool-based
	// network load balancers, internal proxy load balancers, external proxy load
	// balancers, Traffic Director, external protocol forwarding, and Classic VPN.
	// Some products have restrictions on what ports can be used. See
	// [port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)
	// for details.
	//
	//
	// Only packets addressed to ports in the specified range will be forwarded to
	// the backends configured with this forwarding rule.
	//
	// The 'ports' and 'port_range' fields are mutually exclusive.
	//
	// For external forwarding rules, two or more forwarding rules cannot use the
	// same '[IPAddress, IPProtocol]' pair, and cannot have
	// overlapping 'portRange's.
	//
	// For internal forwarding rules within the same VPC network, two or more
	// forwarding rules cannot use the same '[IPAddress, IPProtocol]'
	// pair, and cannot have overlapping 'portRange's.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#port_range ComputeForwardingRule#port_range}
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
	// This field can only be used:.
	//
	// If 'IPProtocol' is one of TCP, UDP, or SCTP.
	// By internal TCP/UDP load balancers, backend service-based network load
	// balancers, internal protocol forwarding and when protocol is not L3_DEFAULT.
	//
	//
	// You can specify a list of up to five ports by number, separated by commas.
	// The ports can be contiguous or discontiguous. Only packets addressed to
	// these ports will be forwarded to the backends configured with this
	// forwarding rule.
	//
	// For external forwarding rules, two or more forwarding rules cannot use the
	// same '[IPAddress, IPProtocol]' pair, and cannot share any values
	// defined in 'ports'.
	//
	// For internal forwarding rules within the same VPC network, two or more
	// forwarding rules cannot use the same '[IPAddress, IPProtocol]'
	// pair, and cannot share any values defined in 'ports'.
	//
	// The 'ports' and 'port_range' fields are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#ports ComputeForwardingRule#ports}
	Ports *[]*string `field:"optional" json:"ports" yaml:"ports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#project ComputeForwardingRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region where the regional forwarding rule resides.
	//
	// This field is not applicable to global forwarding rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#region ComputeForwardingRule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// service_directory_registrations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#service_directory_registrations ComputeForwardingRule#service_directory_registrations}
	ServiceDirectoryRegistrations *ComputeForwardingRuleServiceDirectoryRegistrations `field:"optional" json:"serviceDirectoryRegistrations" yaml:"serviceDirectoryRegistrations"`
	// An optional prefix to the service name for this Forwarding Rule.
	//
	// If specified, will be the first label of the fully qualified service
	// name.
	//
	// The label must be 1-63 characters long, and comply with RFC1035.
	// Specifically, the label must be 1-63 characters long and match the
	// regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	// character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// This field is only used for INTERNAL load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#service_label ComputeForwardingRule#service_label}
	ServiceLabel *string `field:"optional" json:"serviceLabel" yaml:"serviceLabel"`
	// If not empty, this Forwarding Rule will only forward the traffic when the source IP address matches one of the IP addresses or CIDR ranges set here.
	//
	// Note that a Forwarding Rule can only have up to 64 source IP ranges, and this field can only be used with a regional Forwarding Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#source_ip_ranges ComputeForwardingRule#source_ip_ranges}
	SourceIpRanges *[]*string `field:"optional" json:"sourceIpRanges" yaml:"sourceIpRanges"`
	// This field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule, used in internal load balancing and network load balancing with IPv6.
	//
	// If the network specified is in auto subnet mode, this field is optional.
	// However, a subnetwork must be specified if the network is in custom subnet
	// mode or when creating external forwarding rule with IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#subnetwork ComputeForwardingRule#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The URL of the target resource to receive the matched traffic.
	//
	// For
	// regional forwarding rules, this target must be in the same region as the
	// forwarding rule. For global forwarding rules, this target must be a global
	// load balancing resource.
	//
	// The forwarded traffic must be of a type appropriate to the target object.
	// For load balancers, see the "Target" column in [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	// For Private Service Connect forwarding rules that forward traffic to Google APIs, provide the name of a supported Google API bundle:
	//  'vpc-sc' - [ APIs that support VPC Service Controls](https://cloud.google.com/vpc-service-controls/docs/supported-products).
	//  'all-apis' - [All supported Google APIs](https://cloud.google.com/vpc/docs/private-service-connect#supported-apis).
	//
	//
	// For Private Service Connect forwarding rules that forward traffic to managed services, the target must be a service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#target ComputeForwardingRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_forwarding_rule#timeouts ComputeForwardingRule#timeouts}
	Timeouts *ComputeForwardingRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

