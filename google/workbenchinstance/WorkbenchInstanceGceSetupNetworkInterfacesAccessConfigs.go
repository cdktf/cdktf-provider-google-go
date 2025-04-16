// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance


type WorkbenchInstanceGceSetupNetworkInterfacesAccessConfigs struct {
	// An external IP address associated with this instance.
	//
	// Specify an unused
	// static external IP address available to the project or leave this field
	// undefined to use an IP from a shared ephemeral IP address pool. If you
	// specify a static external IP address, it must live in the same region as
	// the zone of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/workbench_instance#external_ip WorkbenchInstance#external_ip}
	ExternalIp *string `field:"required" json:"externalIp" yaml:"externalIp"`
}

