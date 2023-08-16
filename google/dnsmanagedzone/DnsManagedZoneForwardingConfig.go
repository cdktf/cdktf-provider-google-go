package dnsmanagedzone


type DnsManagedZoneForwardingConfig struct {
	// target_name_servers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_managed_zone#target_name_servers DnsManagedZone#target_name_servers}
	TargetNameServers interface{} `field:"required" json:"targetNameServers" yaml:"targetNameServers"`
}

