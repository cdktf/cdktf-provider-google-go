// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionSparkMetastoreServiceConfig struct {
	// Resource name of an existing Dataproc Metastore service in the form of projects/[projectId]/locations/[region]/services/[serviceId].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/bigquery_connection#metastore_service BigqueryConnection#metastore_service}
	MetastoreService *string `field:"optional" json:"metastoreService" yaml:"metastoreService"`
}

