// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionInstanceTemplateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#disk ComputeRegionInstanceTemplate#disk}
	Disk interface{} `field:"required" json:"disk" yaml:"disk"`
	// The machine type to create.
	//
	// To create a machine with a custom type (such as extended memory), format the value like custom-VCPUS-MEM_IN_MB like custom-6-20480 for 6 vCPU and 20GB of RAM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#machine_type ComputeRegionInstanceTemplate#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// advanced_machine_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#advanced_machine_features ComputeRegionInstanceTemplate#advanced_machine_features}
	AdvancedMachineFeatures *ComputeRegionInstanceTemplateAdvancedMachineFeatures `field:"optional" json:"advancedMachineFeatures" yaml:"advancedMachineFeatures"`
	// Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#can_ip_forward ComputeRegionInstanceTemplate#can_ip_forward}
	CanIpForward interface{} `field:"optional" json:"canIpForward" yaml:"canIpForward"`
	// confidential_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#confidential_instance_config ComputeRegionInstanceTemplate#confidential_instance_config}
	ConfidentialInstanceConfig *ComputeRegionInstanceTemplateConfidentialInstanceConfig `field:"optional" json:"confidentialInstanceConfig" yaml:"confidentialInstanceConfig"`
	// A brief description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#description ComputeRegionInstanceTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// guest_accelerator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#guest_accelerator ComputeRegionInstanceTemplate#guest_accelerator}
	GuestAccelerator interface{} `field:"optional" json:"guestAccelerator" yaml:"guestAccelerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#id ComputeRegionInstanceTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A description of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#instance_description ComputeRegionInstanceTemplate#instance_description}
	InstanceDescription *string `field:"optional" json:"instanceDescription" yaml:"instanceDescription"`
	// A set of key/value label pairs to assign to instances created from this template,.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#labels ComputeRegionInstanceTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Metadata key/value pairs to make available from within instances created from this template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#metadata ComputeRegionInstanceTemplate#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// An alternative to using the startup-script metadata key, mostly to match the compute_instance resource.
	//
	// This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#metadata_startup_script ComputeRegionInstanceTemplate#metadata_startup_script}
	MetadataStartupScript *string `field:"optional" json:"metadataStartupScript" yaml:"metadataStartupScript"`
	// Specifies a minimum CPU platform.
	//
	// Applicable values are the friendly names of CPU platforms, such as Intel Haswell or Intel Skylake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#min_cpu_platform ComputeRegionInstanceTemplate#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// The name of the instance template. If you leave this blank, Terraform will auto-generate a unique name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#name ComputeRegionInstanceTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#name_prefix ComputeRegionInstanceTemplate#name_prefix}
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#network_interface ComputeRegionInstanceTemplate#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// network_performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#network_performance_config ComputeRegionInstanceTemplate#network_performance_config}
	NetworkPerformanceConfig *ComputeRegionInstanceTemplateNetworkPerformanceConfig `field:"optional" json:"networkPerformanceConfig" yaml:"networkPerformanceConfig"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#project ComputeRegionInstanceTemplate#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region in which the instance template is located. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#region ComputeRegionInstanceTemplate#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#reservation_affinity ComputeRegionInstanceTemplate#reservation_affinity}
	ReservationAffinity *ComputeRegionInstanceTemplateReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#resource_manager_tags ComputeRegionInstanceTemplate#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
	// A list of self_links of resource policies to attach to the instance.
	//
	// Currently a max of 1 resource policy is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#resource_policies ComputeRegionInstanceTemplate#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#scheduling ComputeRegionInstanceTemplate#scheduling}
	Scheduling *ComputeRegionInstanceTemplateScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#service_account ComputeRegionInstanceTemplate#service_account}
	ServiceAccount *ComputeRegionInstanceTemplateServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#shielded_instance_config ComputeRegionInstanceTemplate#shielded_instance_config}
	ShieldedInstanceConfig *ComputeRegionInstanceTemplateShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Tags to attach to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#tags ComputeRegionInstanceTemplate#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_region_instance_template#timeouts ComputeRegionInstanceTemplate#timeouts}
	Timeouts *ComputeRegionInstanceTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

