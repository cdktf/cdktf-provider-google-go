// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritymirroringdeploymentgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityMirroringDeploymentGroupConfig struct {
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
	// The cloud location of the deployment group, currently restricted to 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_security_mirroring_deployment_group#location NetworkSecurityMirroringDeploymentGroup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID to use for the new deployment group, which will become the final component of the deployment group's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_security_mirroring_deployment_group#mirroring_deployment_group_id NetworkSecurityMirroringDeploymentGroup#mirroring_deployment_group_id}
	MirroringDeploymentGroupId *string `field:"required" json:"mirroringDeploymentGroupId" yaml:"mirroringDeploymentGroupId"`
	// The network that will be used for all child deployments, for example: 'projects/{project}/global/networks/{network}'. See https://google.aip.dev/124.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_security_mirroring_deployment_group#network NetworkSecurityMirroringDeploymentGroup#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// User-provided description of the deployment group. Used as additional context for the deployment group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_security_mirroring_deployment_group#description NetworkSecurityMirroringDeploymentGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_security_mirroring_deployment_group#id NetworkSecurityMirroringDeploymentGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels are key/value pairs that help to organize and filter resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_security_mirroring_deployment_group#labels NetworkSecurityMirroringDeploymentGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_security_mirroring_deployment_group#project NetworkSecurityMirroringDeploymentGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/network_security_mirroring_deployment_group#timeouts NetworkSecurityMirroringDeploymentGroup#timeouts}
	Timeouts *NetworkSecurityMirroringDeploymentGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

