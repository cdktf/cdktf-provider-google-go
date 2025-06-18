// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabruntimetemplate


type ColabRuntimeTemplateSoftwareConfig struct {
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/colab_runtime_template#env ColabRuntimeTemplate#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// post_startup_script_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/colab_runtime_template#post_startup_script_config ColabRuntimeTemplate#post_startup_script_config}
	PostStartupScriptConfig *ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig `field:"optional" json:"postStartupScriptConfig" yaml:"postStartupScriptConfig"`
}

