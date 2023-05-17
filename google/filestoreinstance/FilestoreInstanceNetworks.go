package filestoreinstance


type FilestoreInstanceNetworks struct {
	// IP versions for which the instance has IP addresses assigned. Possible values: ["ADDRESS_MODE_UNSPECIFIED", "MODE_IPV4", "MODE_IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/filestore_instance#modes FilestoreInstance#modes}
	Modes *[]*string `field:"required" json:"modes" yaml:"modes"`
	// The name of the GCE VPC network to which the instance is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/filestore_instance#network FilestoreInstance#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The network connect mode of the Filestore instance.
	//
	// If not provided, the connect mode defaults to
	// DIRECT_PEERING. Default value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/filestore_instance#connect_mode FilestoreInstance#connect_mode}
	ConnectMode *string `field:"optional" json:"connectMode" yaml:"connectMode"`
	// A /29 CIDR block that identifies the range of IP addresses reserved for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/filestore_instance#reserved_ip_range FilestoreInstance#reserved_ip_range}
	ReservedIpRange *string `field:"optional" json:"reservedIpRange" yaml:"reservedIpRange"`
}

