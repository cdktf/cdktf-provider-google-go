// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityinterceptendpointgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityInterceptEndpointGroupConfig struct {
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
	// The deployment group that this endpoint group is connected to, for example: 'projects/123456789/locations/global/interceptDeploymentGroups/my-dg'. See https://google.aip.dev/124.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_intercept_endpoint_group#intercept_deployment_group NetworkSecurityInterceptEndpointGroup#intercept_deployment_group}
	InterceptDeploymentGroup *string `field:"required" json:"interceptDeploymentGroup" yaml:"interceptDeploymentGroup"`
	// The ID to use for the endpoint group, which will become the final component of the endpoint group's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_intercept_endpoint_group#intercept_endpoint_group_id NetworkSecurityInterceptEndpointGroup#intercept_endpoint_group_id}
	InterceptEndpointGroupId *string `field:"required" json:"interceptEndpointGroupId" yaml:"interceptEndpointGroupId"`
	// The cloud location of the endpoint group, currently restricted to 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_intercept_endpoint_group#location NetworkSecurityInterceptEndpointGroup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// User-provided description of the endpoint group. Used as additional context for the endpoint group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_intercept_endpoint_group#description NetworkSecurityInterceptEndpointGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_intercept_endpoint_group#id NetworkSecurityInterceptEndpointGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels are key/value pairs that help to organize and filter resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_intercept_endpoint_group#labels NetworkSecurityInterceptEndpointGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_intercept_endpoint_group#project NetworkSecurityInterceptEndpointGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/network_security_intercept_endpoint_group#timeouts NetworkSecurityInterceptEndpointGroup#timeouts}
	Timeouts *NetworkSecurityInterceptEndpointGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

