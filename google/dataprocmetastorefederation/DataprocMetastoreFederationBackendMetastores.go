// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastorefederation


type DataprocMetastoreFederationBackendMetastores struct {
	// The type of the backend metastore. Possible values: ["METASTORE_TYPE_UNSPECIFIED", "DATAPROC_METASTORE", "BIGQUERY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/dataproc_metastore_federation#metastore_type DataprocMetastoreFederation#metastore_type}
	MetastoreType *string `field:"required" json:"metastoreType" yaml:"metastoreType"`
	// The relative resource name of the metastore that is being federated.
	//
	// The formats of the relative resource names for the currently supported metastores are listed below: Dataplex: projects/{projectId}/locations/{location}/lakes/{lake_id} BigQuery: projects/{projectId} Dataproc Metastore: projects/{projectId}/locations/{location}/services/{serviceId}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/dataproc_metastore_federation#name DataprocMetastoreFederation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/dataproc_metastore_federation#rank DataprocMetastoreFederation#rank}.
	Rank *string `field:"required" json:"rank" yaml:"rank"`
}

