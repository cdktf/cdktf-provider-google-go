package osconfigospolicyassignment


type OsConfigOsPolicyAssignmentInstanceFilterInclusionLabels struct {
	// Labels are identified by key/value pairs in this map.
	//
	// A VM should contain all the key/value pairs specified in this map to be selected.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_os_policy_assignment#labels OsConfigOsPolicyAssignment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

