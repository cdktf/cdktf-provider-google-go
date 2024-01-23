// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appenginedomainmapping


type AppEngineDomainMappingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/app_engine_domain_mapping#create AppEngineDomainMapping#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/app_engine_domain_mapping#delete AppEngineDomainMapping#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/app_engine_domain_mapping#update AppEngineDomainMapping#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

