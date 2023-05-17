package dataproccluster


type DataprocClusterClusterConfigSecurityConfig struct {
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#kerberos_config DataprocCluster#kerberos_config}
	KerberosConfig *DataprocClusterClusterConfigSecurityConfigKerberosConfig `field:"required" json:"kerberosConfig" yaml:"kerberosConfig"`
}

