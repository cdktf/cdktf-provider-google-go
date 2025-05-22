// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstancePscInstanceConfig struct {
	// List of consumer projects that are allowed to create PSC endpoints to service-attachments to this instance.
	//
	// These should be specified as project numbers only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/alloydb_instance#allowed_consumer_projects AlloydbInstance#allowed_consumer_projects}
	AllowedConsumerProjects *[]*string `field:"optional" json:"allowedConsumerProjects" yaml:"allowedConsumerProjects"`
	// psc_auto_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/alloydb_instance#psc_auto_connections AlloydbInstance#psc_auto_connections}
	PscAutoConnections interface{} `field:"optional" json:"pscAutoConnections" yaml:"pscAutoConnections"`
	// psc_interface_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/alloydb_instance#psc_interface_configs AlloydbInstance#psc_interface_configs}
	PscInterfaceConfigs interface{} `field:"optional" json:"pscInterfaceConfigs" yaml:"pscInterfaceConfigs"`
}

