// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectinsightsconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeveloperConnectInsightsConfigConfig struct {
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
	// The name of the App Hub Application. Format: projects/{project}/locations/{location}/applications/{application}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#app_hub_application DeveloperConnectInsightsConfig#app_hub_application}
	AppHubApplication *string `field:"required" json:"appHubApplication" yaml:"appHubApplication"`
	// ID of the requesting InsightsConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#insights_config_id DeveloperConnectInsightsConfig#insights_config_id}
	InsightsConfigId *string `field:"required" json:"insightsConfigId" yaml:"insightsConfigId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#location DeveloperConnectInsightsConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// User specified annotations. See https://google.aip.dev/148#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#annotations DeveloperConnectInsightsConfig#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// artifact_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#artifact_configs DeveloperConnectInsightsConfig#artifact_configs}
	ArtifactConfigs interface{} `field:"optional" json:"artifactConfigs" yaml:"artifactConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#id DeveloperConnectInsightsConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of labels associated with an InsightsConfig.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#labels DeveloperConnectInsightsConfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#project DeveloperConnectInsightsConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/developer_connect_insights_config#timeouts DeveloperConnectInsightsConfig#timeouts}
	Timeouts *DeveloperConnectInsightsConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

