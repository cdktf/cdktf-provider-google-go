// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch


type DataprocBatchSparkSqlBatch struct {
	// HCFS URIs of jar files to be added to the Spark CLASSPATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dataproc_batch#jar_file_uris DataprocBatch#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// The HCFS URI of the script that contains Spark SQL queries to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dataproc_batch#query_file_uri DataprocBatch#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// Mapping of query variable names to values (equivalent to the Spark SQL command: SET name="value";).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dataproc_batch#query_variables DataprocBatch#query_variables}
	QueryVariables *map[string]*string `field:"optional" json:"queryVariables" yaml:"queryVariables"`
}

