// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeappgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeAppGroupConfig struct {
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
	// Name of the AppGroup. Characters you can use in the name are restricted to: A-Z0-9._-$ %.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#name ApigeeAppGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Apigee Organization associated with the Apigee app group, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#org_id ApigeeAppGroup#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#attributes ApigeeAppGroup#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// Channel identifier identifies the owner maintaining this grouping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#channel_id ApigeeAppGroup#channel_id}
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// A reference to the associated storefront/marketplace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#channel_uri ApigeeAppGroup#channel_uri}
	ChannelUri *string `field:"optional" json:"channelUri" yaml:"channelUri"`
	// App group name displayed in the UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#display_name ApigeeAppGroup#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#id ApigeeAppGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Valid values are active or inactive.
	//
	// Note that the status of the AppGroup should be updated via UpdateAppGroupRequest by setting the action as active or inactive. Possible values: ["active", "inactive"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#status ApigeeAppGroup#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_app_group#timeouts ApigeeAppGroup#timeouts}
	Timeouts *ApigeeAppGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

