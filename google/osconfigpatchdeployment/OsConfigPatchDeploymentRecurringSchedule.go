package osconfigpatchdeployment


type OsConfigPatchDeploymentRecurringSchedule struct {
	// time_of_day block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#time_of_day OsConfigPatchDeployment#time_of_day}
	TimeOfDay *OsConfigPatchDeploymentRecurringScheduleTimeOfDay `field:"required" json:"timeOfDay" yaml:"timeOfDay"`
	// time_zone block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#time_zone OsConfigPatchDeployment#time_zone}
	TimeZone *OsConfigPatchDeploymentRecurringScheduleTimeZone `field:"required" json:"timeZone" yaml:"timeZone"`
	// The end time at which a recurring patch deployment schedule is no longer active.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#end_time OsConfigPatchDeployment#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// monthly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#monthly OsConfigPatchDeployment#monthly}
	Monthly *OsConfigPatchDeploymentRecurringScheduleMonthly `field:"optional" json:"monthly" yaml:"monthly"`
	// The time that the recurring schedule becomes effective.
	//
	// Defaults to createTime of the patch deployment.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#start_time OsConfigPatchDeployment#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_config_patch_deployment#weekly OsConfigPatchDeployment#weekly}
	Weekly *OsConfigPatchDeploymentRecurringScheduleWeekly `field:"optional" json:"weekly" yaml:"weekly"`
}

