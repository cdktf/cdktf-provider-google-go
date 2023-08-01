package dataprocmetastoreservice


type DataprocMetastoreServiceHiveMetastoreConfig struct {
	// The Hive metastore schema version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_metastore_service#version DataprocMetastoreService#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml). The mappings override system defaults (some keys cannot be overridden).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_metastore_service#config_overrides DataprocMetastoreService#config_overrides}
	ConfigOverrides *map[string]*string `field:"optional" json:"configOverrides" yaml:"configOverrides"`
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_metastore_service#kerberos_config DataprocMetastoreService#kerberos_config}
	KerberosConfig *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig `field:"optional" json:"kerberosConfig" yaml:"kerberosConfig"`
}

