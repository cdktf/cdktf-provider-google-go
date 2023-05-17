package notebooksruntime


type NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams struct {
	// Provide this property when creating the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#description NotebooksRuntime#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the disk name.
	//
	// If not specified, the default is
	// to use the name of the instance. If the disk with the
	// instance name exists already in the given zone/region, a
	// new name will be automatically generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#disk_name NotebooksRuntime#disk_name}
	DiskName *string `field:"optional" json:"diskName" yaml:"diskName"`
	// Specifies the size of the disk in base-2 GB.
	//
	// If not
	// specified, the disk will be the same size as the image
	// (usually 10GB). If specified, the size must be equal to
	// or larger than 10GB. Default 100 GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#disk_size_gb NotebooksRuntime#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The type of the boot disk attached to this runtime, defaults to standard persistent disk. For valid values, see 'https://cloud.google.com/vertex-ai/docs/workbench/ reference/rest/v1/projects.locations.runtimes#disktype'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#disk_type NotebooksRuntime#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Labels to apply to this disk.
	//
	// These can be later modified
	// by the disks.setLabels method. This field is only
	// applicable for persistent disks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#labels NotebooksRuntime#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

