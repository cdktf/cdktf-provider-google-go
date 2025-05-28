// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubnamespace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeHubNamespaceConfig struct {
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
	// The name of the Scope instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gke_hub_namespace#scope GkeHubNamespace#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Id of the scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gke_hub_namespace#scope_id GkeHubNamespace#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The client-provided identifier of the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gke_hub_namespace#scope_namespace_id GkeHubNamespace#scope_namespace_id}
	ScopeNamespaceId *string `field:"required" json:"scopeNamespaceId" yaml:"scopeNamespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gke_hub_namespace#id GkeHubNamespace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels for this Namespace.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gke_hub_namespace#labels GkeHubNamespace#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespace-level cluster namespace labels.
	//
	// These labels are applied
	// to the related namespace of the member clusters bound to the parent
	// Scope. Scope-level labels ('namespace_labels' in the Fleet Scope
	// resource) take precedence over Namespace-level labels if they share
	// a key. Keys and values must be Kubernetes-conformant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gke_hub_namespace#namespace_labels GkeHubNamespace#namespace_labels}
	NamespaceLabels *map[string]*string `field:"optional" json:"namespaceLabels" yaml:"namespaceLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gke_hub_namespace#project GkeHubNamespace#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gke_hub_namespace#timeouts GkeHubNamespace#timeouts}
	Timeouts *GkeHubNamespaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

