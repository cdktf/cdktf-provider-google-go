// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan


type DataplexDatascanDataDiscoverySpecBigqueryPublishingConfig struct {
	// The BigQuery connection used to create BigLake tables. Must be in the form 'projects/{projectId}/locations/{locationId}/connections/{connection_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dataplex_datascan#connection DataplexDatascan#connection}
	Connection *string `field:"optional" json:"connection" yaml:"connection"`
	// The location of the BigQuery dataset to publish BigLake external or non-BigLake external tables to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dataplex_datascan#location DataplexDatascan#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The project of the BigQuery dataset to publish BigLake external or non-BigLake external tables to.
	//
	// If not specified, the project of the Cloud Storage bucket will be used. The format is "projects/{project_id_or_number}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dataplex_datascan#project DataplexDatascan#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Determines whether to publish discovered tables as BigLake external tables or non-BigLake external tables. Possible values: ["TABLE_TYPE_UNSPECIFIED", "EXTERNAL", "BIGLAKE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/dataplex_datascan#table_type DataplexDatascan#table_type}
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
}

