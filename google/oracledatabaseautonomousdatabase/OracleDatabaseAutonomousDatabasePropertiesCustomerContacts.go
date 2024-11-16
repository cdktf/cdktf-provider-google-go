// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabaseautonomousdatabase


type OracleDatabaseAutonomousDatabasePropertiesCustomerContacts struct {
	// The email address used by Oracle to send notifications regarding databases and infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/oracle_database_autonomous_database#email OracleDatabaseAutonomousDatabase#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

