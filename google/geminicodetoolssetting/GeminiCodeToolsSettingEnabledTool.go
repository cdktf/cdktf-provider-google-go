// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminicodetoolssetting


type GeminiCodeToolsSettingEnabledTool struct {
	// Handle used to invoke the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gemini_code_tools_setting#handle GeminiCodeToolsSetting#handle}
	Handle *string `field:"required" json:"handle" yaml:"handle"`
	// Link to the Tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gemini_code_tools_setting#tool GeminiCodeToolsSetting#tool}
	Tool *string `field:"required" json:"tool" yaml:"tool"`
	// Link to the Dev Connect Account Connector that holds the user credentials. projects/{project}/locations/{location}/accountConnectors/{account_connector_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gemini_code_tools_setting#account_connector GeminiCodeToolsSetting#account_connector}
	AccountConnector *string `field:"optional" json:"accountConnector" yaml:"accountConnector"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gemini_code_tools_setting#config GeminiCodeToolsSetting#config}
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Overridden URI, if allowed by Tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gemini_code_tools_setting#uri_override GeminiCodeToolsSetting#uri_override}
	UriOverride *string `field:"optional" json:"uriOverride" yaml:"uriOverride"`
}

