// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeEnvironmentConfig struct {
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
	// The resource ID of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#name ApigeeEnvironment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Apigee Organization associated with the Apigee environment, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#org_id ApigeeEnvironment#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Optional.
	//
	// API Proxy type supported by the environment. The type can be set when creating
	// the Environment and cannot be changed. Possible values: ["API_PROXY_TYPE_UNSPECIFIED", "PROGRAMMABLE", "CONFIGURABLE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#api_proxy_type ApigeeEnvironment#api_proxy_type}
	ApiProxyType *string `field:"optional" json:"apiProxyType" yaml:"apiProxyType"`
	// Optional.
	//
	// Deployment type supported by the environment. The deployment type can be
	// set when creating the environment and cannot be changed. When you enable archive
	// deployment, you will be prevented from performing a subset of actions within the
	// environment, including:
	// Managing the deployment of API proxy or shared flow revisions;
	// Creating, updating, or deleting resource files;
	// Creating, updating, or deleting target servers. Possible values: ["DEPLOYMENT_TYPE_UNSPECIFIED", "PROXY", "ARCHIVE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#deployment_type ApigeeEnvironment#deployment_type}
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// Description of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#description ApigeeEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#display_name ApigeeEnvironment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Optional.
	//
	// URI of the forward proxy to be applied to the runtime instances in this environment. Must be in the format of {scheme}://{hostname}:{port}. Note that the scheme must be one of "http" or "https", and the port must be supplied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#forward_proxy_uri ApigeeEnvironment#forward_proxy_uri}
	ForwardProxyUri *string `field:"optional" json:"forwardProxyUri" yaml:"forwardProxyUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#id ApigeeEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#node_config ApigeeEnvironment#node_config}
	NodeConfig *ApigeeEnvironmentNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#properties ApigeeEnvironment#properties}
	Properties *ApigeeEnvironmentProperties `field:"optional" json:"properties" yaml:"properties"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#timeouts ApigeeEnvironment#timeouts}
	Timeouts *ApigeeEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Types that can be selected for an Environment.
	//
	// Each of the types are
	// limited by capability and capacity. Refer to Apigee's public documentation
	// to understand about each of these types in details.
	// An Apigee org can support heterogeneous Environments. Possible values: ["ENVIRONMENT_TYPE_UNSPECIFIED", "BASE", "INTERMEDIATE", "COMPREHENSIVE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/apigee_environment#type ApigeeEnvironment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

