// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment


type ComposerEnvironmentConfigRecoveryConfig struct {
	// scheduled_snapshots_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/composer_environment#scheduled_snapshots_config ComposerEnvironment#scheduled_snapshots_config}
	ScheduledSnapshotsConfig *ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig `field:"optional" json:"scheduledSnapshotsConfig" yaml:"scheduledSnapshotsConfig"`
}

