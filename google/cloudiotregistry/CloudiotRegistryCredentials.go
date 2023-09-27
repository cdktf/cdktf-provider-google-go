// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudiotregistry


type CloudiotRegistryCredentials struct {
	// A public key certificate format and data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/cloudiot_registry#public_key_certificate CloudiotRegistry#public_key_certificate}
	PublicKeyCertificate *map[string]*string `field:"required" json:"publicKeyCertificate" yaml:"publicKeyCertificate"`
}

