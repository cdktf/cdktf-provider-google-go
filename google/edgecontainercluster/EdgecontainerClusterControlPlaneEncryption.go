// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterControlPlaneEncryption struct {
	// The Cloud KMS CryptoKey e.g. projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey} to use for protecting control plane disks. If not specified, a Google-managed key will be used instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgecontainer_cluster#kms_key EdgecontainerCluster#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

