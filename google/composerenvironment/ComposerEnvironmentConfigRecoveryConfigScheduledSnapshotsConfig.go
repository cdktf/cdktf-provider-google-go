package composerenvironment


type ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig struct {
	// When enabled, Cloud Composer periodically saves snapshots of your environment to a Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#enabled ComposerEnvironment#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Snapshot schedule, in the unix-cron format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#snapshot_creation_schedule ComposerEnvironment#snapshot_creation_schedule}
	SnapshotCreationSchedule *string `field:"optional" json:"snapshotCreationSchedule" yaml:"snapshotCreationSchedule"`
	// the URI of a bucket folder where to save the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#snapshot_location ComposerEnvironment#snapshot_location}
	SnapshotLocation *string `field:"optional" json:"snapshotLocation" yaml:"snapshotLocation"`
	// A time zone for the schedule.
	//
	// This value is a time offset and does not take into account daylight saving time changes. Valid values are from UTC-12 to UTC+12. Examples: UTC, UTC-01, UTC+03.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#time_zone ComposerEnvironment#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

