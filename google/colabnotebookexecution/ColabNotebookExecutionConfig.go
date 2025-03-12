// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabnotebookexecution

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ColabNotebookExecutionConfig struct {
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
	// Required. The display name of the Notebook Execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#display_name ColabNotebookExecution#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Cloud Storage location to upload the result to. Format:'gs://bucket-name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#gcs_output_uri ColabNotebookExecution#gcs_output_uri}
	GcsOutputUri *string `field:"required" json:"gcsOutputUri" yaml:"gcsOutputUri"`
	// The location for the resource: https://cloud.google.com/colab/docs/locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#location ColabNotebookExecution#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// dataform_repository_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#dataform_repository_source ColabNotebookExecution#dataform_repository_source}
	DataformRepositorySource *ColabNotebookExecutionDataformRepositorySource `field:"optional" json:"dataformRepositorySource" yaml:"dataformRepositorySource"`
	// direct_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#direct_notebook_source ColabNotebookExecution#direct_notebook_source}
	DirectNotebookSource *ColabNotebookExecutionDirectNotebookSource `field:"optional" json:"directNotebookSource" yaml:"directNotebookSource"`
	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#execution_timeout ColabNotebookExecution#execution_timeout}
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// The user email to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#execution_user ColabNotebookExecution#execution_user}
	ExecutionUser *string `field:"optional" json:"executionUser" yaml:"executionUser"`
	// gcs_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#gcs_notebook_source ColabNotebookExecution#gcs_notebook_source}
	GcsNotebookSource *ColabNotebookExecutionGcsNotebookSource `field:"optional" json:"gcsNotebookSource" yaml:"gcsNotebookSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#id ColabNotebookExecution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User specified ID for the Notebook Execution Job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#notebook_execution_job_id ColabNotebookExecution#notebook_execution_job_id}
	NotebookExecutionJobId *string `field:"optional" json:"notebookExecutionJobId" yaml:"notebookExecutionJobId"`
	// The NotebookRuntimeTemplate to source compute configuration from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#notebook_runtime_template_resource_name ColabNotebookExecution#notebook_runtime_template_resource_name}
	NotebookRuntimeTemplateResourceName *string `field:"optional" json:"notebookRuntimeTemplateResourceName" yaml:"notebookRuntimeTemplateResourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#project ColabNotebookExecution#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The service account to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#service_account ColabNotebookExecution#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/colab_notebook_execution#timeouts ColabNotebookExecution#timeouts}
	Timeouts *ColabNotebookExecutionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

