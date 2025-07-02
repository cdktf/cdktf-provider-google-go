// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubcuration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApihubCurationConfig struct {
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
	// The ID to use for the curation resource, which will become the final component of the curations's resource name.
	//
	// This field is optional.
	//
	// * If provided, the same will be used. The service will throw an error if
	// the specified ID is already used by another curation resource in the API
	// hub.
	// * If not provided, a system generated ID will be used.
	//
	// This value should be 4-500 characters, and valid characters
	// are /a-z[0-9]-_/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apihub_curation#curation_id ApihubCuration#curation_id}
	CurationId *string `field:"required" json:"curationId" yaml:"curationId"`
	// The display name of the curation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apihub_curation#display_name ApihubCuration#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apihub_curation#endpoint ApihubCuration#endpoint}
	Endpoint *ApihubCurationEndpoint `field:"required" json:"endpoint" yaml:"endpoint"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apihub_curation#location ApihubCuration#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The description of the curation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apihub_curation#description ApihubCuration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apihub_curation#id ApihubCuration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apihub_curation#project ApihubCuration#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/apihub_curation#timeouts ApihubCuration#timeouts}
	Timeouts *ApihubCurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

