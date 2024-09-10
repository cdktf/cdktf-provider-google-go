// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileOracleForwardSshConnectivity struct {
	// Required. Hostname for the SSH tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/database_migration_service_connection_profile#hostname DatabaseMigrationServiceConnectionProfile#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Port for the SSH tunnel, default value is 22.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/database_migration_service_connection_profile#port DatabaseMigrationServiceConnectionProfile#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Required. Username for the SSH tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/database_migration_service_connection_profile#username DatabaseMigrationServiceConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Input only. SSH password. Only one of 'password' and 'private_key' can be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/database_migration_service_connection_profile#password DatabaseMigrationServiceConnectionProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only. SSH private key. Only one of 'password' and 'private_key' can be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/database_migration_service_connection_profile#private_key DatabaseMigrationServiceConnectionProfile#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
}

