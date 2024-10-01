// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydataset


type BigqueryDatasetAccessDataset struct {
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/bigquery_dataset#dataset BigqueryDataset#dataset}
	Dataset *BigqueryDatasetAccessDatasetDataset `field:"required" json:"dataset" yaml:"dataset"`
	// Which resources in the dataset this entry applies to.
	//
	// Currently, only views are supported,
	// but additional target types may be added in the future. Possible values: VIEWS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/bigquery_dataset#target_types BigqueryDataset#target_types}
	TargetTypes *[]*string `field:"required" json:"targetTypes" yaml:"targetTypes"`
}

