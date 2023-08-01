package osconfigpatchdeployment


type OsConfigPatchDeploymentRecurringScheduleTimeOfDay struct {
	// Hours of day in 24 hour format.
	//
	// Should be from 0 to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#hours OsConfigPatchDeployment#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Minutes of hour of day. Must be from 0 to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#minutes OsConfigPatchDeployment#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#nanos OsConfigPatchDeployment#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time.
	//
	// Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/os_config_patch_deployment#seconds OsConfigPatchDeployment#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

