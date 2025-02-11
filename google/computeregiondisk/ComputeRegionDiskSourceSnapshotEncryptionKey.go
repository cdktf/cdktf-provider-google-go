// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregiondisk


type ComputeRegionDiskSourceSnapshotEncryptionKey struct {
	// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_region_disk#raw_key ComputeRegionDisk#raw_key}
	RawKey *string `field:"optional" json:"rawKey" yaml:"rawKey"`
}

