// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocsessiontemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocSessionTemplateConfig struct {
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
	// The resource name of the session template in the following format: projects/{project}/locations/{location}/sessionTemplates/{template_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#name DataprocSessionTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#environment_config DataprocSessionTemplate#environment_config}
	EnvironmentConfig *DataprocSessionTemplateEnvironmentConfig `field:"optional" json:"environmentConfig" yaml:"environmentConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#id DataprocSessionTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// jupyter_session block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#jupyter_session DataprocSessionTemplate#jupyter_session}
	JupyterSession *DataprocSessionTemplateJupyterSession `field:"optional" json:"jupyterSession" yaml:"jupyterSession"`
	// The labels to associate with this session template.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#labels DataprocSessionTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location in which the session template will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#location DataprocSessionTemplate#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#project DataprocSessionTemplate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// runtime_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#runtime_config DataprocSessionTemplate#runtime_config}
	RuntimeConfig *DataprocSessionTemplateRuntimeConfig `field:"optional" json:"runtimeConfig" yaml:"runtimeConfig"`
	// spark_connect_session block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#spark_connect_session DataprocSessionTemplate#spark_connect_session}
	SparkConnectSession *DataprocSessionTemplateSparkConnectSession `field:"optional" json:"sparkConnectSession" yaml:"sparkConnectSession"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dataproc_session_template#timeouts DataprocSessionTemplate#timeouts}
	Timeouts *DataprocSessionTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

