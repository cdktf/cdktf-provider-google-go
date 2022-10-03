package composerenvironment


type ComposerEnvironmentConfigMaintenanceWindow struct {
	// Maintenance window end time.
	//
	// It is used only to calculate the duration of the maintenance window. The value for end-time must be in the future, relative to 'start_time'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#end_time ComposerEnvironment#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Maintenance window recurrence.
	//
	// Format is a subset of RFC-5545 (https://tools.ietf.org/html/rfc5545) 'RRULE'. The only allowed values for 'FREQ' field are 'FREQ=DAILY' and 'FREQ=WEEKLY;BYDAY=...'. Example values: 'FREQ=WEEKLY;BYDAY=TU,WE', 'FREQ=DAILY'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#recurrence ComposerEnvironment#recurrence}
	Recurrence *string `field:"required" json:"recurrence" yaml:"recurrence"`
	// Start time of the first recurrence of the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#start_time ComposerEnvironment#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

