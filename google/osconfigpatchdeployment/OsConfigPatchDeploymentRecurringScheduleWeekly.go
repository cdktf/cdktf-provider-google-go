package osconfigpatchdeployment


type OsConfigPatchDeploymentRecurringScheduleWeekly struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York". Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#day_of_week OsConfigPatchDeployment#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
}

