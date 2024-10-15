// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerybireservation


type BigqueryBiReservationPreferredTables struct {
	// The ID of the dataset in the above project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/bigquery_bi_reservation#dataset_id BigqueryBiReservation#dataset_id}
	DatasetId *string `field:"optional" json:"datasetId" yaml:"datasetId"`
	// The assigned project ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/bigquery_bi_reservation#project_id BigqueryBiReservation#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The ID of the table in the above dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/bigquery_bi_reservation#table_id BigqueryBiReservation#table_id}
	TableId *string `field:"optional" json:"tableId" yaml:"tableId"`
}

