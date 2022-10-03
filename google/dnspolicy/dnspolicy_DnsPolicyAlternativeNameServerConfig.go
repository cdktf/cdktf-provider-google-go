package dnspolicy


type DnsPolicyAlternativeNameServerConfig struct {
	// target_name_servers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_policy#target_name_servers DnsPolicy#target_name_servers}
	TargetNameServers interface{} `field:"required" json:"targetNameServers" yaml:"targetNameServers"`
}

