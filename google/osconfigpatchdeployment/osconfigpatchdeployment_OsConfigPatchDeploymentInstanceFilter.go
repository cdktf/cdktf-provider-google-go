package osconfigpatchdeployment


type OsConfigPatchDeploymentInstanceFilter struct {
	// Target all VM instances in the project. If true, no other criteria is permitted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#all OsConfigPatchDeployment#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// group_labels block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#group_labels OsConfigPatchDeployment#group_labels}
	GroupLabels interface{} `field:"optional" json:"groupLabels" yaml:"groupLabels"`
	// Targets VMs whose name starts with one of these prefixes.
	//
	// Similar to labels, this is another way to group
	// VMs when targeting configs, for example prefix="prod-".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#instance_name_prefixes OsConfigPatchDeployment#instance_name_prefixes}
	InstanceNamePrefixes *[]*string `field:"optional" json:"instanceNamePrefixes" yaml:"instanceNamePrefixes"`
	// Targets any of the VM instances specified. Instances are specified by their URI in the 'form zones/{{zone}}/instances/{{instance_name}}', 'projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}', or 'https://www.googleapis.com/compute/v1/projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#instances OsConfigPatchDeployment#instances}
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_config_patch_deployment#zones OsConfigPatchDeployment#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

