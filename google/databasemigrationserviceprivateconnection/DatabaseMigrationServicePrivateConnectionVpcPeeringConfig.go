// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceprivateconnection


type DatabaseMigrationServicePrivateConnectionVpcPeeringConfig struct {
	// A free subnet for peering. (CIDR of /29).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/database_migration_service_private_connection#subnet DatabaseMigrationServicePrivateConnection#subnet}
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
	// Fully qualified name of the VPC that Database Migration Service will peer to. Format: projects/{project}/global/{networks}/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/database_migration_service_private_connection#vpc_name DatabaseMigrationServicePrivateConnection#vpc_name}
	VpcName *string `field:"required" json:"vpcName" yaml:"vpcName"`
}

