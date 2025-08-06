// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable


type BigqueryTableExternalCatalogTableOptionsStorageDescriptor struct {
	// Specifies the fully qualified class name of the InputFormat (e.g. "org.apache.hadoop.hive.ql.io.orc.OrcInputFormat"). The maximum length is 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/bigquery_table#input_format BigqueryTable#input_format}
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// The physical location of the table (e.g. 'gs://spark-dataproc-data/pangea-data/case_sensitive/' or 'gs://spark-dataproc-data/pangea-data/*'). The maximum length is 2056 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/bigquery_table#location_uri BigqueryTable#location_uri}
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// Specifies the fully qualified class name of the OutputFormat (e.g. "org.apache.hadoop.hive.ql.io.orc.OrcOutputFormat"). The maximum length is 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/bigquery_table#output_format BigqueryTable#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// serde_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/bigquery_table#serde_info BigqueryTable#serde_info}
	SerdeInfo *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo `field:"optional" json:"serdeInfo" yaml:"serdeInfo"`
}

