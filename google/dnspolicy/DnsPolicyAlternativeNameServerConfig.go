package dnspolicy


type DnsPolicyAlternativeNameServerConfig struct {
	// target_name_servers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dns_policy#target_name_servers DnsPolicy#target_name_servers}
	TargetNameServers interface{} `field:"required" json:"targetNameServers" yaml:"targetNameServers"`
}

