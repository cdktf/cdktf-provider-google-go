// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigSecurityConfig struct {
	// identity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dataproc_cluster#identity_config DataprocCluster#identity_config}
	IdentityConfig *DataprocClusterClusterConfigSecurityConfigIdentityConfig `field:"optional" json:"identityConfig" yaml:"identityConfig"`
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dataproc_cluster#kerberos_config DataprocCluster#kerberos_config}
	KerberosConfig *DataprocClusterClusterConfigSecurityConfigKerberosConfig `field:"optional" json:"kerberosConfig" yaml:"kerberosConfig"`
}

