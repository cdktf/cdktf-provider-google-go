// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritymirroringendpointgroupassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityMirroringEndpointGroupAssociationConfig struct {
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
	// The cloud location of the association, currently restricted to 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_mirroring_endpoint_group_association#location NetworkSecurityMirroringEndpointGroupAssociation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The endpoint group that this association is connected to, for example: 'projects/123456789/locations/global/mirroringEndpointGroups/my-eg'. See https://google.aip.dev/124.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_mirroring_endpoint_group_association#mirroring_endpoint_group NetworkSecurityMirroringEndpointGroupAssociation#mirroring_endpoint_group}
	MirroringEndpointGroup *string `field:"required" json:"mirroringEndpointGroup" yaml:"mirroringEndpointGroup"`
	// The VPC network that is associated. for example: 'projects/123456789/global/networks/my-network'. See https://google.aip.dev/124.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_mirroring_endpoint_group_association#network NetworkSecurityMirroringEndpointGroupAssociation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_mirroring_endpoint_group_association#id NetworkSecurityMirroringEndpointGroupAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels are key/value pairs that help to organize and filter resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_mirroring_endpoint_group_association#labels NetworkSecurityMirroringEndpointGroupAssociation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The ID to use for the new association, which will become the final component of the endpoint group's resource name.
	//
	// If not provided, the
	// server will generate a unique ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_mirroring_endpoint_group_association#mirroring_endpoint_group_association_id NetworkSecurityMirroringEndpointGroupAssociation#mirroring_endpoint_group_association_id}
	MirroringEndpointGroupAssociationId *string `field:"optional" json:"mirroringEndpointGroupAssociationId" yaml:"mirroringEndpointGroupAssociationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_mirroring_endpoint_group_association#project NetworkSecurityMirroringEndpointGroupAssociation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/network_security_mirroring_endpoint_group_association#timeouts NetworkSecurityMirroringEndpointGroupAssociation#timeouts}
	Timeouts *NetworkSecurityMirroringEndpointGroupAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

