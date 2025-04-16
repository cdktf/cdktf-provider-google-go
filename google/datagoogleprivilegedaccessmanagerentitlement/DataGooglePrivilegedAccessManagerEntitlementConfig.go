// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleprivilegedaccessmanagerentitlement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGooglePrivilegedAccessManagerEntitlementConfig struct {
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
	// The ID to use for this Entitlement.
	//
	// This will become the last part of the resource name.
	// This value should be 4-63 characters, and valid characters are "[a-z]", "[0-9]", and "-". The first character should be from [a-z].
	// This value should be unique among all other Entitlements under the specified 'parent'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/privileged_access_manager_entitlement#entitlement_id DataGooglePrivilegedAccessManagerEntitlement#entitlement_id}
	EntitlementId *string `field:"optional" json:"entitlementId" yaml:"entitlementId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/privileged_access_manager_entitlement#id DataGooglePrivilegedAccessManagerEntitlement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The region of the Entitlement resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/privileged_access_manager_entitlement#location DataGooglePrivilegedAccessManagerEntitlement#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Format: projects/{project-id|project-number} or organizations/{organization-number} or folders/{folder-number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/data-sources/privileged_access_manager_entitlement#parent DataGooglePrivilegedAccessManagerEntitlement#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
}

