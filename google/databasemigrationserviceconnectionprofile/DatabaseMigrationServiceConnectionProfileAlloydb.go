// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileAlloydb struct {
	// Required. The AlloyDB cluster ID that this connection profile is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/database_migration_service_connection_profile#cluster_id DatabaseMigrationServiceConnectionProfile#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/database_migration_service_connection_profile#settings DatabaseMigrationServiceConnectionProfile#settings}
	Settings *DatabaseMigrationServiceConnectionProfileAlloydbSettings `field:"optional" json:"settings" yaml:"settings"`
}

