// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationcenterpreferenceset


type MigrationCenterPreferenceSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/migration_center_preference_set#create MigrationCenterPreferenceSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/migration_center_preference_set#delete MigrationCenterPreferenceSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/migration_center_preference_set#update MigrationCenterPreferenceSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

