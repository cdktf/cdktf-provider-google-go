// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabruntimetemplate


type ColabRuntimeTemplateIdleShutdownConfig struct {
	// The duration after which the runtime is automatically shut down.
	//
	// An input of 0s disables the idle shutdown feature, and a valid range is [10m, 24h].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/colab_runtime_template#idle_timeout ColabRuntimeTemplate#idle_timeout}
	IdleTimeout *string `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

