// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastorefederation


type DataprocMetastoreFederationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/dataproc_metastore_federation#create DataprocMetastoreFederation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/dataproc_metastore_federation#delete DataprocMetastoreFederation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/dataproc_metastore_federation#update DataprocMetastoreFederation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

