package sqldatabaseinstance


type SqlDatabaseInstanceSettingsIpConfiguration struct {
	// The name of the allocated ip range for the private ip CloudSQL instance.
	//
	// For example: "google-managed-services-default". If set, the instance ip will be created in the allocated range. The range name must comply with RFC 1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sql_database_instance#allocated_ip_range SqlDatabaseInstance#allocated_ip_range}
	AllocatedIpRange *string `field:"optional" json:"allocatedIpRange" yaml:"allocatedIpRange"`
	// authorized_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sql_database_instance#authorized_networks SqlDatabaseInstance#authorized_networks}
	AuthorizedNetworks interface{} `field:"optional" json:"authorizedNetworks" yaml:"authorizedNetworks"`
	// Whether Google Cloud services such as BigQuery are allowed to access data in this Cloud SQL instance over a private IP connection.
	//
	// SQLSERVER database type is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sql_database_instance#enable_private_path_for_google_cloud_services SqlDatabaseInstance#enable_private_path_for_google_cloud_services}
	EnablePrivatePathForGoogleCloudServices interface{} `field:"optional" json:"enablePrivatePathForGoogleCloudServices" yaml:"enablePrivatePathForGoogleCloudServices"`
	// Whether this Cloud SQL instance should be assigned a public IPV4 address.
	//
	// At least ipv4_enabled must be enabled or a private_network must be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sql_database_instance#ipv4_enabled SqlDatabaseInstance#ipv4_enabled}
	Ipv4Enabled interface{} `field:"optional" json:"ipv4Enabled" yaml:"ipv4Enabled"`
	// The VPC network from which the Cloud SQL instance is accessible for private IP.
	//
	// For example, projects/myProject/global/networks/default. Specifying a network enables private IP. At least ipv4_enabled must be enabled or a private_network must be configured. This setting can be updated, but it cannot be removed after it is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sql_database_instance#private_network SqlDatabaseInstance#private_network}
	PrivateNetwork *string `field:"optional" json:"privateNetwork" yaml:"privateNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/sql_database_instance#require_ssl SqlDatabaseInstance#require_ssl}.
	RequireSsl interface{} `field:"optional" json:"requireSsl" yaml:"requireSsl"`
}

