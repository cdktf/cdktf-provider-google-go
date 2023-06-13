package composerenvironment


type ComposerEnvironmentConfigRecoveryConfig struct {
	// scheduled_snapshots_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#scheduled_snapshots_config ComposerEnvironment#scheduled_snapshots_config}
	ScheduledSnapshotsConfig *ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig `field:"optional" json:"scheduledSnapshotsConfig" yaml:"scheduledSnapshotsConfig"`
}

