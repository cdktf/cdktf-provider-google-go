// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceSettingsLocationPreference struct {
	// A Google App Engine application whose zone to remain in. Must be in the same region as this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/sql_database_instance#follow_gae_application SqlDatabaseInstance#follow_gae_application}
	FollowGaeApplication *string `field:"optional" json:"followGaeApplication" yaml:"followGaeApplication"`
	// The preferred Compute Engine zone for the secondary/failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/sql_database_instance#secondary_zone SqlDatabaseInstance#secondary_zone}
	SecondaryZone *string `field:"optional" json:"secondaryZone" yaml:"secondaryZone"`
	// The preferred compute engine zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/sql_database_instance#zone SqlDatabaseInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

