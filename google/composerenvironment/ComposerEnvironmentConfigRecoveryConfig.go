package composerenvironment


type ComposerEnvironmentConfigRecoveryConfig struct {
	// scheduled_snapshots_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#scheduled_snapshots_config ComposerEnvironment#scheduled_snapshots_config}
	ScheduledSnapshotsConfig *ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig `field:"optional" json:"scheduledSnapshotsConfig" yaml:"scheduledSnapshotsConfig"`
}

