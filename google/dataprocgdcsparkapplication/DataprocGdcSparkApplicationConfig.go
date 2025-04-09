// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocgdcsparkapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocGdcSparkApplicationConfig struct {
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
	// The location of the spark application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#location DataprocGdcSparkApplication#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The id of the service instance to which this spark application belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#serviceinstance DataprocGdcSparkApplication#serviceinstance}
	Serviceinstance *string `field:"required" json:"serviceinstance" yaml:"serviceinstance"`
	// The id of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#spark_application_id DataprocGdcSparkApplication#spark_application_id}
	SparkApplicationId *string `field:"required" json:"sparkApplicationId" yaml:"sparkApplicationId"`
	// The annotations to associate with this application.
	//
	// Annotations may be used to store client information, but are not used by the server.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#annotations DataprocGdcSparkApplication#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// An ApplicationEnvironment from which to inherit configuration properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#application_environment DataprocGdcSparkApplication#application_environment}
	ApplicationEnvironment *string `field:"optional" json:"applicationEnvironment" yaml:"applicationEnvironment"`
	// List of container image uris for additional file dependencies.
	//
	// Dependent files are sequentially copied from each image. If a file with the same name exists in 2 images then the file from later image is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#dependency_images DataprocGdcSparkApplication#dependency_images}
	DependencyImages *[]*string `field:"optional" json:"dependencyImages" yaml:"dependencyImages"`
	// User-provided human-readable name to be used in user interfaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#display_name DataprocGdcSparkApplication#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#id DataprocGdcSparkApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels to associate with this application. Labels may be used for filtering and billing tracking.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#labels DataprocGdcSparkApplication#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The Kubernetes namespace in which to create the application. This namespace must already exist on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#namespace DataprocGdcSparkApplication#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#project DataprocGdcSparkApplication#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// application-specific properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#properties DataprocGdcSparkApplication#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// pyspark_application_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#pyspark_application_config DataprocGdcSparkApplication#pyspark_application_config}
	PysparkApplicationConfig *DataprocGdcSparkApplicationPysparkApplicationConfig `field:"optional" json:"pysparkApplicationConfig" yaml:"pysparkApplicationConfig"`
	// spark_application_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#spark_application_config DataprocGdcSparkApplication#spark_application_config}
	SparkApplicationConfig *DataprocGdcSparkApplicationSparkApplicationConfig `field:"optional" json:"sparkApplicationConfig" yaml:"sparkApplicationConfig"`
	// spark_r_application_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#spark_r_application_config DataprocGdcSparkApplication#spark_r_application_config}
	SparkRApplicationConfig *DataprocGdcSparkApplicationSparkRApplicationConfig `field:"optional" json:"sparkRApplicationConfig" yaml:"sparkRApplicationConfig"`
	// spark_sql_application_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#spark_sql_application_config DataprocGdcSparkApplication#spark_sql_application_config}
	SparkSqlApplicationConfig *DataprocGdcSparkApplicationSparkSqlApplicationConfig `field:"optional" json:"sparkSqlApplicationConfig" yaml:"sparkSqlApplicationConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#timeouts DataprocGdcSparkApplication#timeouts}
	Timeouts *DataprocGdcSparkApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The Dataproc version of this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataproc_gdc_spark_application#version DataprocGdcSparkApplication#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

