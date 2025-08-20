// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesauthzextension


type NetworkServicesAuthzExtensionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_services_authz_extension#create NetworkServicesAuthzExtension#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_services_authz_extension#delete NetworkServicesAuthzExtension#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_services_authz_extension#update NetworkServicesAuthzExtension#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

