// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecloudidentitygrouplookup


type DataGoogleCloudIdentityGroupLookupGroupKey struct {
	// The ID of the entity.
	//
	// For Google-managed entities, the id should be the email address of an existing group or user.
	// For external-identity-mapped entities, the id must be a string conforming to the Identity Source's requirements.
	// Must be unique within a namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/data-sources/cloud_identity_group_lookup#id DataGoogleCloudIdentityGroupLookup#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The namespace in which the entity exists.
	//
	// If not specified, the EntityKey represents a Google-managed entity such as a Google user or a Google Group.
	// If specified, the EntityKey represents an external-identity-mapped group. The namespace must correspond to an identity source created in Admin Console and must be in the form of identitysources/{identity_source}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/data-sources/cloud_identity_group_lookup#namespace DataGoogleCloudIdentityGroupLookup#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

