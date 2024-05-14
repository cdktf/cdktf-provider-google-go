// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceNetworkConfig struct {
	// consumers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/dataproc_metastore_service#consumers DataprocMetastoreService#consumers}
	Consumers interface{} `field:"required" json:"consumers" yaml:"consumers"`
}

