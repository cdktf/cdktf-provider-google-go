package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig struct {
	// authorized_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/database_migration_service_connection_profile#authorized_networks DatabaseMigrationServiceConnectionProfile#authorized_networks}
	AuthorizedNetworks interface{} `field:"optional" json:"authorizedNetworks" yaml:"authorizedNetworks"`
	// Whether the instance should be assigned an IPv4 address or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/database_migration_service_connection_profile#enable_ipv4 DatabaseMigrationServiceConnectionProfile#enable_ipv4}
	EnableIpv4 interface{} `field:"optional" json:"enableIpv4" yaml:"enableIpv4"`
	// The resource link for the VPC network from which the Cloud SQL instance is accessible for private IP.
	//
	// For example, projects/myProject/global/networks/default.
	// This setting can be updated, but it cannot be removed after it is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/database_migration_service_connection_profile#private_network DatabaseMigrationServiceConnectionProfile#private_network}
	PrivateNetwork *string `field:"optional" json:"privateNetwork" yaml:"privateNetwork"`
	// Whether SSL connections over IP should be enforced or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/database_migration_service_connection_profile#require_ssl DatabaseMigrationServiceConnectionProfile#require_ssl}
	RequireSsl interface{} `field:"optional" json:"requireSsl" yaml:"requireSsl"`
}

