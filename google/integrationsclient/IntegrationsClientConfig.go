// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsclient

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationsClientConfig struct {
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
	// Location in which client needs to be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/integrations_client#location IntegrationsClient#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// cloud_kms_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/integrations_client#cloud_kms_config IntegrationsClient#cloud_kms_config}
	CloudKmsConfig *IntegrationsClientCloudKmsConfig `field:"optional" json:"cloudKmsConfig" yaml:"cloudKmsConfig"`
	// Indicates if sample integrations should be created along with provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/integrations_client#create_sample_integrations IntegrationsClient#create_sample_integrations}
	CreateSampleIntegrations interface{} `field:"optional" json:"createSampleIntegrations" yaml:"createSampleIntegrations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/integrations_client#id IntegrationsClient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/integrations_client#project IntegrationsClient#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// User input run-as service account, if empty, will bring up a new default service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/integrations_client#run_as_service_account IntegrationsClient#run_as_service_account}
	RunAsServiceAccount *string `field:"optional" json:"runAsServiceAccount" yaml:"runAsServiceAccount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/integrations_client#timeouts IntegrationsClient#timeouts}
	Timeouts *IntegrationsClientTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

