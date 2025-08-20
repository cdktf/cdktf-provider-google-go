// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubapihubinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApihubApiHubInstanceConfig struct {
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
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_api_hub_instance#config ApihubApiHubInstance#config}
	Config *ApihubApiHubInstanceConfigA `field:"required" json:"config" yaml:"config"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_api_hub_instance#location ApihubApiHubInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Optional.
	//
	// Identifier to assign to the Api Hub instance. Must be unique within
	// scope of the parent resource. If the field is not provided,
	// system generated id will be used.
	//
	// This value should be 4-40 characters, and valid characters
	// are '/a-z[0-9]-_/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_api_hub_instance#api_hub_instance_id ApihubApiHubInstance#api_hub_instance_id}
	ApiHubInstanceId *string `field:"optional" json:"apiHubInstanceId" yaml:"apiHubInstanceId"`
	// Optional. Description of the ApiHub instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_api_hub_instance#description ApihubApiHubInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_api_hub_instance#id ApihubApiHubInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Instance labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_api_hub_instance#labels ApihubApiHubInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_api_hub_instance#project ApihubApiHubInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apihub_api_hub_instance#timeouts ApihubApiHubInstance#timeouts}
	Timeouts *ApihubApiHubInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

