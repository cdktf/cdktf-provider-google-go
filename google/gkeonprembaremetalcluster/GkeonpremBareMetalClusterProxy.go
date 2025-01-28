// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterProxy struct {
	// Specifies the address of your proxy server.
	//
	// For example: http://domain
	// WARNING: Do not provide credentials in the format
	// of http://(username:password@)domain these will be rejected by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gkeonprem_bare_metal_cluster#uri GkeonpremBareMetalCluster#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// A list of IPs, hostnames, and domains that should skip the proxy. For example ["127.0.0.1", "example.com", ".corp", "localhost"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gkeonprem_bare_metal_cluster#no_proxy GkeonpremBareMetalCluster#no_proxy}
	NoProxy *[]*string `field:"optional" json:"noProxy" yaml:"noProxy"`
}

