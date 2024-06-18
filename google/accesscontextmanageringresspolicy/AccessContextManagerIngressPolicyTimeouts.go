// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanageringresspolicy


type AccessContextManagerIngressPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/access_context_manager_ingress_policy#create AccessContextManagerIngressPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/access_context_manager_ingress_policy#delete AccessContextManagerIngressPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

