package dnsmanagedzone


type DnsManagedZoneCloudLoggingConfig struct {
	// If set, enable query logging for this ManagedZone. False by default, making logging opt-in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_managed_zone#enable_logging DnsManagedZone#enable_logging}
	EnableLogging interface{} `field:"required" json:"enableLogging" yaml:"enableLogging"`
}

