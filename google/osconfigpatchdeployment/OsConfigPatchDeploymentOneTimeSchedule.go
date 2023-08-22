package osconfigpatchdeployment


type OsConfigPatchDeploymentOneTimeSchedule struct {
	// The desired patch job execution time. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/os_config_patch_deployment#execute_time OsConfigPatchDeployment#execute_time}
	ExecuteTime *string `field:"required" json:"executeTime" yaml:"executeTime"`
}

