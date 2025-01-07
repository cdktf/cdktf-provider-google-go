// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composeruserworkloadsconfigmap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerUserWorkloadsConfigMapConfig struct {
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
	// Environment where the Kubernetes ConfigMap will be stored and used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/composer_user_workloads_config_map#environment ComposerUserWorkloadsConfigMap#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Name of the Kubernetes ConfigMap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/composer_user_workloads_config_map#name ComposerUserWorkloadsConfigMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The "data" field of Kubernetes ConfigMap, organized in key-value pairs. For details see: https://kubernetes.io/docs/concepts/configuration/configmap/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/composer_user_workloads_config_map#data ComposerUserWorkloadsConfigMap#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/composer_user_workloads_config_map#id ComposerUserWorkloadsConfigMap#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/composer_user_workloads_config_map#project ComposerUserWorkloadsConfigMap#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The location or Compute Engine region for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/composer_user_workloads_config_map#region ComposerUserWorkloadsConfigMap#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/composer_user_workloads_config_map#timeouts ComposerUserWorkloadsConfigMap#timeouts}
	Timeouts *ComposerUserWorkloadsConfigMapTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

