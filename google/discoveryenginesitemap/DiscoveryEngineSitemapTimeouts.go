// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginesitemap


type DiscoveryEngineSitemapTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/discovery_engine_sitemap#create DiscoveryEngineSitemap#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/discovery_engine_sitemap#delete DiscoveryEngineSitemap#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

