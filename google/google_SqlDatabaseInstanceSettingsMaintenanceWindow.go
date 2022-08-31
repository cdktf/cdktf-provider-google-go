// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type SqlDatabaseInstanceSettingsMaintenanceWindow struct {
	// Day of week (1-7), starting on Monday.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#day SqlDatabaseInstance#day}
	Day *float64 `field:"optional" json:"day" yaml:"day"`
	// Hour of day (0-23), ignored if day not set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#hour SqlDatabaseInstance#hour}
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
	// Receive updates earlier (canary) or later (stable).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#update_track SqlDatabaseInstance#update_track}
	UpdateTrack *string `field:"optional" json:"updateTrack" yaml:"updateTrack"`
}

