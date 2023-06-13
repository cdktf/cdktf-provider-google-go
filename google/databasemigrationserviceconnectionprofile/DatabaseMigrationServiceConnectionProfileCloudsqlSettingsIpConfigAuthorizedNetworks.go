package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfigAuthorizedNetworks struct {
	// The allowlisted value for the access control list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/database_migration_service_connection_profile#value DatabaseMigrationServiceConnectionProfile#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The time when this access control entry expires in RFC 3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/database_migration_service_connection_profile#expire_time DatabaseMigrationServiceConnectionProfile#expire_time}
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// A label to identify this entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/database_migration_service_connection_profile#label DatabaseMigrationServiceConnectionProfile#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Input only. The time-to-leave of this access control entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/database_migration_service_connection_profile#ttl DatabaseMigrationServiceConnectionProfile#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

