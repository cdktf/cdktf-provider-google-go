// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileOracle struct {
	// Required. Database service for the Oracle connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#database_service DatabaseMigrationServiceConnectionProfile#database_service}
	DatabaseService *string `field:"required" json:"databaseService" yaml:"databaseService"`
	// Required. The IP or hostname of the source Oracle database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#host DatabaseMigrationServiceConnectionProfile#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Required.
	//
	// Input only. The password for the user that Database Migration Service will be using to connect to the database.
	// This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#password DatabaseMigrationServiceConnectionProfile#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Required. The network port of the source Oracle database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#port DatabaseMigrationServiceConnectionProfile#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Required.
	//
	// The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#username DatabaseMigrationServiceConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// forward_ssh_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#forward_ssh_connectivity DatabaseMigrationServiceConnectionProfile#forward_ssh_connectivity}
	ForwardSshConnectivity *DatabaseMigrationServiceConnectionProfileOracleForwardSshConnectivity `field:"optional" json:"forwardSshConnectivity" yaml:"forwardSshConnectivity"`
	// private_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#private_connectivity DatabaseMigrationServiceConnectionProfile#private_connectivity}
	PrivateConnectivity *DatabaseMigrationServiceConnectionProfileOraclePrivateConnectivity `field:"optional" json:"privateConnectivity" yaml:"privateConnectivity"`
	// ssl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#ssl DatabaseMigrationServiceConnectionProfile#ssl}
	Ssl *DatabaseMigrationServiceConnectionProfileOracleSsl `field:"optional" json:"ssl" yaml:"ssl"`
	// static_service_ip_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/database_migration_service_connection_profile#static_service_ip_connectivity DatabaseMigrationServiceConnectionProfile#static_service_ip_connectivity}
	StaticServiceIpConnectivity *DatabaseMigrationServiceConnectionProfileOracleStaticServiceIpConnectivity `field:"optional" json:"staticServiceIpConnectivity" yaml:"staticServiceIpConnectivity"`
}

