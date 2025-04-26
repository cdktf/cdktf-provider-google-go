// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryengineschema


type DiscoveryEngineSchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/discovery_engine_schema#create DiscoveryEngineSchema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/discovery_engine_schema#delete DiscoveryEngineSchema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

