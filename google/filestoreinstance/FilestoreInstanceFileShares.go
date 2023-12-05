// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestoreinstance


type FilestoreInstanceFileShares struct {
	// File share capacity in GiB.
	//
	// This must be at least 1024 GiB
	// for the standard tier, or 2560 GiB for the premium tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.8.0/docs/resources/filestore_instance#capacity_gb FilestoreInstance#capacity_gb}
	CapacityGb *float64 `field:"required" json:"capacityGb" yaml:"capacityGb"`
	// The name of the fileshare (16 characters or less).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.8.0/docs/resources/filestore_instance#name FilestoreInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// nfs_export_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.8.0/docs/resources/filestore_instance#nfs_export_options FilestoreInstance#nfs_export_options}
	NfsExportOptions interface{} `field:"optional" json:"nfsExportOptions" yaml:"nfsExportOptions"`
}

