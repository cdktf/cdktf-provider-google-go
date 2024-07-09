// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigSecurityConfig struct {
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/dataproc_cluster#kerberos_config DataprocCluster#kerberos_config}
	KerberosConfig *DataprocClusterClusterConfigSecurityConfigKerberosConfig `field:"required" json:"kerberosConfig" yaml:"kerberosConfig"`
}

