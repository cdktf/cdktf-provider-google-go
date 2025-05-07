// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocBatchConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID to use for the batch, which will become the final component of the batch's resource name.
	//
	// This value must be 4-63 characters. Valid characters are /[a-z][0-9]-/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#batch_id DataprocBatch#batch_id}
	BatchId *string `field:"optional" json:"batchId" yaml:"batchId"`
	// environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#environment_config DataprocBatch#environment_config}
	EnvironmentConfig *DataprocBatchEnvironmentConfig `field:"optional" json:"environmentConfig" yaml:"environmentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#id DataprocBatch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels to associate with this batch.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#labels DataprocBatch#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location in which the batch will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#location DataprocBatch#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#project DataprocBatch#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// pyspark_batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#pyspark_batch DataprocBatch#pyspark_batch}
	PysparkBatch *DataprocBatchPysparkBatch `field:"optional" json:"pysparkBatch" yaml:"pysparkBatch"`
	// runtime_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#runtime_config DataprocBatch#runtime_config}
	RuntimeConfig *DataprocBatchRuntimeConfig `field:"optional" json:"runtimeConfig" yaml:"runtimeConfig"`
	// spark_batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#spark_batch DataprocBatch#spark_batch}
	SparkBatch *DataprocBatchSparkBatch `field:"optional" json:"sparkBatch" yaml:"sparkBatch"`
	// spark_r_batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#spark_r_batch DataprocBatch#spark_r_batch}
	SparkRBatch *DataprocBatchSparkRBatch `field:"optional" json:"sparkRBatch" yaml:"sparkRBatch"`
	// spark_sql_batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#spark_sql_batch DataprocBatch#spark_sql_batch}
	SparkSqlBatch *DataprocBatchSparkSqlBatch `field:"optional" json:"sparkSqlBatch" yaml:"sparkSqlBatch"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/dataproc_batch#timeouts DataprocBatch#timeouts}
	Timeouts *DataprocBatchTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

