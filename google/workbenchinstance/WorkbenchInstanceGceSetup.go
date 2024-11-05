// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance


type WorkbenchInstanceGceSetup struct {
	// accelerator_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#accelerator_configs WorkbenchInstance#accelerator_configs}
	AcceleratorConfigs interface{} `field:"optional" json:"acceleratorConfigs" yaml:"acceleratorConfigs"`
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#boot_disk WorkbenchInstance#boot_disk}
	BootDisk *WorkbenchInstanceGceSetupBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// container_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#container_image WorkbenchInstance#container_image}
	ContainerImage *WorkbenchInstanceGceSetupContainerImage `field:"optional" json:"containerImage" yaml:"containerImage"`
	// data_disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#data_disks WorkbenchInstance#data_disks}
	DataDisks *WorkbenchInstanceGceSetupDataDisks `field:"optional" json:"dataDisks" yaml:"dataDisks"`
	// Optional. If true, no external IP will be assigned to this VM instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#disable_public_ip WorkbenchInstance#disable_public_ip}
	DisablePublicIp interface{} `field:"optional" json:"disablePublicIp" yaml:"disablePublicIp"`
	// Optional. Flag to enable ip forwarding or not, default false/off. https://cloud.google.com/vpc/docs/using-routes#canipforward.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#enable_ip_forwarding WorkbenchInstance#enable_ip_forwarding}
	EnableIpForwarding interface{} `field:"optional" json:"enableIpForwarding" yaml:"enableIpForwarding"`
	// Optional. The machine type of the VM instance. https://cloud.google.com/compute/docs/machine-resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#machine_type WorkbenchInstance#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Optional. Custom metadata to apply to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#metadata WorkbenchInstance#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// network_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#network_interfaces WorkbenchInstance#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// service_accounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#service_accounts WorkbenchInstance#service_accounts}
	ServiceAccounts interface{} `field:"optional" json:"serviceAccounts" yaml:"serviceAccounts"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#shielded_instance_config WorkbenchInstance#shielded_instance_config}
	ShieldedInstanceConfig *WorkbenchInstanceGceSetupShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Optional. The Compute Engine tags to add to instance (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#tags WorkbenchInstance#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// vm_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/workbench_instance#vm_image WorkbenchInstance#vm_image}
	VmImage *WorkbenchInstanceGceSetupVmImage `field:"optional" json:"vmImage" yaml:"vmImage"`
}

