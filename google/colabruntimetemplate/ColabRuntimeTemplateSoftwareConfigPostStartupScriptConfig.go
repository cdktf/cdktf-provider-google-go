// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabruntimetemplate


type ColabRuntimeTemplateSoftwareConfigPostStartupScriptConfig struct {
	// Post startup script to run after runtime is started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_runtime_template#post_startup_script ColabRuntimeTemplate#post_startup_script}
	PostStartupScript *string `field:"optional" json:"postStartupScript" yaml:"postStartupScript"`
	// Post startup script behavior that defines download and execution behavior. Possible values: ["RUN_ONCE", "RUN_EVERY_START", "DOWNLOAD_AND_RUN_EVERY_START"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_runtime_template#post_startup_script_behavior ColabRuntimeTemplate#post_startup_script_behavior}
	PostStartupScriptBehavior *string `field:"optional" json:"postStartupScriptBehavior" yaml:"postStartupScriptBehavior"`
	// Post startup script url to download. Example: https://bucket/script.sh.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_runtime_template#post_startup_script_url ColabRuntimeTemplate#post_startup_script_url}
	PostStartupScriptUrl *string `field:"optional" json:"postStartupScriptUrl" yaml:"postStartupScriptUrl"`
}

