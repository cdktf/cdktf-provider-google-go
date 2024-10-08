// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationservicemigrationjob


type DatabaseMigrationServiceMigrationJobPerformanceConfig struct {
	// Initial dump parallelism level. Possible values: ["MIN", "OPTIMAL", "MAX"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/database_migration_service_migration_job#dump_parallel_level DatabaseMigrationServiceMigrationJob#dump_parallel_level}
	DumpParallelLevel *string `field:"optional" json:"dumpParallelLevel" yaml:"dumpParallelLevel"`
}

