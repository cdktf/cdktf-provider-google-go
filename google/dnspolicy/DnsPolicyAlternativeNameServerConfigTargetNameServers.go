package dnspolicy


type DnsPolicyAlternativeNameServerConfigTargetNameServers struct {
	// IPv4 address to forward to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#ipv4_address DnsPolicy#ipv4_address}
	Ipv4Address *string `field:"required" json:"ipv4Address" yaml:"ipv4Address"`
	// Forwarding path for this TargetNameServer.
	//
	// If unset or 'default' Cloud DNS will make forwarding
	// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	// to the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: ["default", "private"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#forwarding_path DnsPolicy#forwarding_path}
	ForwardingPath *string `field:"optional" json:"forwardingPath" yaml:"forwardingPath"`
}

