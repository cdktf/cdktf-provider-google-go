// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugininstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApihubPluginInstanceConfig struct {
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
	// The display name for this plugin instance. Max length is 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#display_name ApihubPluginInstance#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#location ApihubPluginInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#plugin ApihubPluginInstance#plugin}
	Plugin *string `field:"required" json:"plugin" yaml:"plugin"`
	// The ID to use for the plugin instance, which will become the final component of the plugin instance's resource name.
	//
	// This field is optional.
	//
	// * If provided, the same will be used. The service will throw an error if
	// the specified id is already used by another plugin instance in the plugin
	// resource.
	// * If not provided, a system generated id will be used.
	//
	// This value should be 4-63 characters, and valid characters
	// are /a-z[0-9]-_/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#plugin_instance_id ApihubPluginInstance#plugin_instance_id}
	PluginInstanceId *string `field:"required" json:"pluginInstanceId" yaml:"pluginInstanceId"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#actions ApihubPluginInstance#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#auth_config ApihubPluginInstance#auth_config}
	AuthConfig *ApihubPluginInstanceAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// The display name for this plugin instance. Max length is 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#disable ApihubPluginInstance#disable}
	Disable interface{} `field:"optional" json:"disable" yaml:"disable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#id ApihubPluginInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#project ApihubPluginInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apihub_plugin_instance#timeouts ApihubPluginInstance#timeouts}
	Timeouts *ApihubPluginInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

