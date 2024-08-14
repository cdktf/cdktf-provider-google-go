// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileAlloydbSettingsInitialUser struct {
	// The initial password for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/database_migration_service_connection_profile#password DatabaseMigrationServiceConnectionProfile#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The database username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/database_migration_service_connection_profile#user DatabaseMigrationServiceConnectionProfile#user}
	User *string `field:"required" json:"user" yaml:"user"`
}

