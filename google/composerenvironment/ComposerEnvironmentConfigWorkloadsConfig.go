// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment


type ComposerEnvironmentConfigWorkloadsConfig struct {
	// scheduler block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/composer_environment#scheduler ComposerEnvironment#scheduler}
	Scheduler *ComposerEnvironmentConfigWorkloadsConfigScheduler `field:"optional" json:"scheduler" yaml:"scheduler"`
	// triggerer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/composer_environment#triggerer ComposerEnvironment#triggerer}
	Triggerer *ComposerEnvironmentConfigWorkloadsConfigTriggerer `field:"optional" json:"triggerer" yaml:"triggerer"`
	// web_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/composer_environment#web_server ComposerEnvironment#web_server}
	WebServer *ComposerEnvironmentConfigWorkloadsConfigWebServer `field:"optional" json:"webServer" yaml:"webServer"`
	// worker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/composer_environment#worker ComposerEnvironment#worker}
	Worker *ComposerEnvironmentConfigWorkloadsConfigWorker `field:"optional" json:"worker" yaml:"worker"`
}

