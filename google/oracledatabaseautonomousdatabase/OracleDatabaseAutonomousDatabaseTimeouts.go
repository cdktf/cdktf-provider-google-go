// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabaseautonomousdatabase


type OracleDatabaseAutonomousDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/oracle_database_autonomous_database#create OracleDatabaseAutonomousDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/oracle_database_autonomous_database#delete OracleDatabaseAutonomousDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/oracle_database_autonomous_database#update OracleDatabaseAutonomousDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

