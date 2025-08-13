// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch


type DataprocBatchEnvironmentConfigPeripheralsConfig struct {
	// Resource name of an existing Dataproc Metastore service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_batch#metastore_service DataprocBatch#metastore_service}
	MetastoreService *string `field:"optional" json:"metastoreService" yaml:"metastoreService"`
	// spark_history_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_batch#spark_history_server_config DataprocBatch#spark_history_server_config}
	SparkHistoryServerConfig *DataprocBatchEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig `field:"optional" json:"sparkHistoryServerConfig" yaml:"sparkHistoryServerConfig"`
}

