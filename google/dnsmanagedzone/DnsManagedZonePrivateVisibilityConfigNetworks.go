package dnsmanagedzone


type DnsManagedZonePrivateVisibilityConfigNetworks struct {
	// The id or fully qualified URL of the VPC network to bind to. This should be formatted like 'projects/{project}/global/networks/{network}' or 'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/dns_managed_zone#network_url DnsManagedZone#network_url}
	NetworkUrl *string `field:"required" json:"networkUrl" yaml:"networkUrl"`
}

