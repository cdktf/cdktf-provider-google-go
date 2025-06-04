// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocgdcserviceinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocGdcServiceInstanceConfig struct {
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
	// Location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#location DataprocGdcServiceInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Id of the service instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#service_instance_id DataprocGdcServiceInstance#service_instance_id}
	ServiceInstanceId *string `field:"required" json:"serviceInstanceId" yaml:"serviceInstanceId"`
	// User-provided human-readable name to be used in user interfaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#display_name DataprocGdcServiceInstance#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// gdce_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#gdce_cluster DataprocGdcServiceInstance#gdce_cluster}
	GdceCluster *DataprocGdcServiceInstanceGdceCluster `field:"optional" json:"gdceCluster" yaml:"gdceCluster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#id DataprocGdcServiceInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels to associate with this service instance. Labels may be used for filtering and billing tracking.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#labels DataprocGdcServiceInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#project DataprocGdcServiceInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Requested service account to associate with ServiceInstance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#service_account DataprocGdcServiceInstance#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// spark_service_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#spark_service_instance_config DataprocGdcServiceInstance#spark_service_instance_config}
	SparkServiceInstanceConfig *DataprocGdcServiceInstanceSparkServiceInstanceConfig `field:"optional" json:"sparkServiceInstanceConfig" yaml:"sparkServiceInstanceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_gdc_service_instance#timeouts DataprocGdcServiceInstance#timeouts}
	Timeouts *DataprocGdcServiceInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

