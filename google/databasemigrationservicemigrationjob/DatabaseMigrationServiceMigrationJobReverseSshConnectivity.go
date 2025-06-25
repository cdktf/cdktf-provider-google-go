// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationservicemigrationjob


type DatabaseMigrationServiceMigrationJobReverseSshConnectivity struct {
	// The name of the virtual machine (Compute Engine) used as the bastion server for the SSH tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/database_migration_service_migration_job#vm DatabaseMigrationServiceMigrationJob#vm}
	Vm *string `field:"optional" json:"vm" yaml:"vm"`
	// The IP of the virtual machine (Compute Engine) used as the bastion server for the SSH tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/database_migration_service_migration_job#vm_ip DatabaseMigrationServiceMigrationJob#vm_ip}
	VmIp *string `field:"optional" json:"vmIp" yaml:"vmIp"`
	// The forwarding port of the virtual machine (Compute Engine) used as the bastion server for the SSH tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/database_migration_service_migration_job#vm_port DatabaseMigrationServiceMigrationJob#vm_port}
	VmPort *float64 `field:"optional" json:"vmPort" yaml:"vmPort"`
	// The name of the VPC to peer with the Cloud SQL private network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/database_migration_service_migration_job#vpc DatabaseMigrationServiceMigrationJob#vpc}
	Vpc *string `field:"optional" json:"vpc" yaml:"vpc"`
}

