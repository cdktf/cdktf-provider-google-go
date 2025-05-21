// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksruntime

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotebooksRuntimeConfig struct {
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
	// A reference to the zone where the machine resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#location NotebooksRuntime#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name specified for the Notebook runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#name NotebooksRuntime#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// access_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#access_config NotebooksRuntime#access_config}
	AccessConfig *NotebooksRuntimeAccessConfig `field:"optional" json:"accessConfig" yaml:"accessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#id NotebooksRuntime#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels to associate with this runtime.
	//
	// Label **keys** must
	// contain 1 to 63 characters, and must conform to [RFC 1035]
	// (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
	// empty, but, if present, must contain 1 to 63 characters, and must
	// conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
	// more than 32 labels can be associated with a cluster.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#labels NotebooksRuntime#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#project NotebooksRuntime#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#software_config NotebooksRuntime#software_config}
	SoftwareConfig *NotebooksRuntimeSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#timeouts NotebooksRuntime#timeouts}
	Timeouts *NotebooksRuntimeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/notebooks_runtime#virtual_machine NotebooksRuntime#virtual_machine}
	VirtualMachine *NotebooksRuntimeVirtualMachine `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
}

