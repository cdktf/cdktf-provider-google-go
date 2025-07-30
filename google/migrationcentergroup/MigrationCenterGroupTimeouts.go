// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationcentergroup


type MigrationCenterGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/migration_center_group#create MigrationCenterGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/migration_center_group#delete MigrationCenterGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/migration_center_group#update MigrationCenterGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

