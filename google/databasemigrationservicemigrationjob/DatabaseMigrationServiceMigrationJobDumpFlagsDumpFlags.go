// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationservicemigrationjob


type DatabaseMigrationServiceMigrationJobDumpFlagsDumpFlags struct {
	// The name of the flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/database_migration_service_migration_job#name DatabaseMigrationServiceMigrationJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The vale of the flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/database_migration_service_migration_job#value DatabaseMigrationServiceMigrationJob#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

