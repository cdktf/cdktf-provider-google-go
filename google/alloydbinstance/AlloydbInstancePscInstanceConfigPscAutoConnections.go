// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstancePscInstanceConfigPscAutoConnections struct {
	// The consumer network for the PSC service automation, example: "projects/vpc-host-project/global/networks/default".
	//
	// The consumer network might be hosted a different project than the
	// consumer project. The API expects the consumer project specified to be
	// the project ID (and not the project number)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/alloydb_instance#consumer_network AlloydbInstance#consumer_network}
	ConsumerNetwork *string `field:"optional" json:"consumerNetwork" yaml:"consumerNetwork"`
	// The consumer project to which the PSC service automation endpoint will be created.
	//
	// The API expects the consumer project to be the project ID(
	// and not the project number).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/alloydb_instance#consumer_project AlloydbInstance#consumer_project}
	ConsumerProject *string `field:"optional" json:"consumerProject" yaml:"consumerProject"`
}

