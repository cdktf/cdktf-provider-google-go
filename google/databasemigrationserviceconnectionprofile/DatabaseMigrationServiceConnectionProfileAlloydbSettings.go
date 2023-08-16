package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileAlloydbSettings struct {
	// initial_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/database_migration_service_connection_profile#initial_user DatabaseMigrationServiceConnectionProfile#initial_user}
	InitialUser *DatabaseMigrationServiceConnectionProfileAlloydbSettingsInitialUser `field:"required" json:"initialUser" yaml:"initialUser"`
	// Required.
	//
	// The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP. The network must belong to the same project as the cluster.
	// It is specified in the form: 'projects/{project_number}/global/networks/{network_id}'. This is required to create a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/database_migration_service_connection_profile#vpc_network DatabaseMigrationServiceConnectionProfile#vpc_network}
	VpcNetwork *string `field:"required" json:"vpcNetwork" yaml:"vpcNetwork"`
	// Labels for the AlloyDB cluster created by DMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/database_migration_service_connection_profile#labels DatabaseMigrationServiceConnectionProfile#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// primary_instance_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/database_migration_service_connection_profile#primary_instance_settings DatabaseMigrationServiceConnectionProfile#primary_instance_settings}
	PrimaryInstanceSettings *DatabaseMigrationServiceConnectionProfileAlloydbSettingsPrimaryInstanceSettings `field:"optional" json:"primaryInstanceSettings" yaml:"primaryInstanceSettings"`
}

