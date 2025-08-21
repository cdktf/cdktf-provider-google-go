// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lustreinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LustreInstanceConfig struct {
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
	// The storage capacity of the instance in gibibytes (GiB). Allowed values are from '18000' to '954000', in increments of 9000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#capacity_gib LustreInstance#capacity_gib}
	CapacityGib *string `field:"required" json:"capacityGib" yaml:"capacityGib"`
	// The filesystem name for this instance.
	//
	// This name is used by client-side
	// tools, including when mounting the instance. Must be eight characters or
	// less and can only contain letters and numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#filesystem LustreInstance#filesystem}
	Filesystem *string `field:"required" json:"filesystem" yaml:"filesystem"`
	// The name of the Managed Lustre instance.
	//
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#instance_id LustreInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#location LustreInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The full name of the VPC network to which the instance is connected. Must be in the format 'projects/{project_id}/global/networks/{network_name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#network LustreInstance#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The throughput of the instance in MB/s/TiB. Valid values are 125, 250, 500, 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#per_unit_storage_throughput LustreInstance#per_unit_storage_throughput}
	PerUnitStorageThroughput *string `field:"required" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// A user-readable description of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#description LustreInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether you want to enable support for GKE clients. By default, GKE clients are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#gke_support_enabled LustreInstance#gke_support_enabled}
	GkeSupportEnabled interface{} `field:"optional" json:"gkeSupportEnabled" yaml:"gkeSupportEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#id LustreInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#labels LustreInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#project LustreInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/lustre_instance#timeouts LustreInstance#timeouts}
	Timeouts *LustreInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

