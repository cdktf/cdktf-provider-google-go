package notebooksruntime


type NotebooksRuntimeVirtualMachineVirtualMachineConfig struct {
	// data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#data_disk NotebooksRuntime#data_disk}
	DataDisk *NotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk `field:"required" json:"dataDisk" yaml:"dataDisk"`
	// The Compute Engine machine type used for runtimes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#machine_type NotebooksRuntime#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// accelerator_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#accelerator_config NotebooksRuntime#accelerator_config}
	AcceleratorConfig *NotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig `field:"optional" json:"acceleratorConfig" yaml:"acceleratorConfig"`
	// container_images block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#container_images NotebooksRuntime#container_images}
	ContainerImages interface{} `field:"optional" json:"containerImages" yaml:"containerImages"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#encryption_config NotebooksRuntime#encryption_config}
	EncryptionConfig *NotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// If true, runtime will only have internal IP addresses.
	//
	// By default,
	// runtimes are not restricted to internal IP addresses, and will
	// have ephemeral external IP addresses assigned to each vm. This
	// 'internal_ip_only' restriction can only be enabled for subnetwork
	// enabled networks, and all dependencies must be configured to be
	// accessible without external IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#internal_ip_only NotebooksRuntime#internal_ip_only}
	InternalIpOnly interface{} `field:"optional" json:"internalIpOnly" yaml:"internalIpOnly"`
	// The labels to associate with this runtime.
	//
	// Label **keys** must
	// contain 1 to 63 characters, and must conform to [RFC 1035]
	// (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
	// empty, but, if present, must contain 1 to 63 characters, and must
	// conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
	// more than 32 labels can be associated with a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#labels NotebooksRuntime#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The Compute Engine metadata entries to add to virtual machine. (see [Project and instance metadata](https://cloud.google.com /compute/docs/storing-retrieving-metadata#project_and_instance _metadata)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#metadata NotebooksRuntime#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// The Compute Engine network to be used for machine communications.
	//
	// Cannot be specified with subnetwork. If neither 'network' nor
	// 'subnet' is specified, the "default" network of the project is
	// used, if it exists. A full URL or partial URI. Examples:
	// 'https://www.googleapis.com/compute/v1/projects/[project_id]/
	// regions/global/default'
	// 'projects/[project_id]/regions/global/default'
	// Runtimes are managed resources inside Google Infrastructure.
	// Runtimes support the following network configurations:
	// Google Managed Network (Network & subnet are empty)
	// Consumer Project VPC (network & subnet are required). Requires
	// configuring Private Service Access.
	// Shared VPC (network & subnet are required). Requires
	// configuring Private Service Access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#network NotebooksRuntime#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The type of vNIC to be used on this interface.
	//
	// This may be gVNIC
	// or VirtioNet. Possible values: ["UNSPECIFIED_NIC_TYPE", "VIRTIO_NET", "GVNIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#nic_type NotebooksRuntime#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// Reserved IP Range name is used for VPC Peering. The subnetwork allocation will use the range *name* if it's assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#reserved_ip_range NotebooksRuntime#reserved_ip_range}
	ReservedIpRange *string `field:"optional" json:"reservedIpRange" yaml:"reservedIpRange"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#shielded_instance_config NotebooksRuntime#shielded_instance_config}
	ShieldedInstanceConfig *NotebooksRuntimeVirtualMachineVirtualMachineConfigShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// The Compute Engine subnetwork to be used for machine communications.
	//
	// Cannot be specified with network. A full URL or
	// partial URI are valid. Examples:
	// 'https://www.googleapis.com/compute/v1/projects/[project_id]/
	// regions/us-east1/subnetworks/sub0'
	// 'projects/[project_id]/regions/us-east1/subnetworks/sub0'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#subnet NotebooksRuntime#subnet}
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
	// The Compute Engine tags to add to runtime (see [Tagging instances] (https://cloud.google.com/compute/docs/ label-or-tag-resources#tags)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/notebooks_runtime#tags NotebooksRuntime#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

