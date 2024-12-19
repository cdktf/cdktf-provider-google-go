// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocgdcserviceinstance


type DataprocGdcServiceInstanceGdceCluster struct {
	// Gdce cluster resource id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/dataproc_gdc_service_instance#gdce_cluster DataprocGdcServiceInstance#gdce_cluster}
	GdceCluster *string `field:"required" json:"gdceCluster" yaml:"gdceCluster"`
}

