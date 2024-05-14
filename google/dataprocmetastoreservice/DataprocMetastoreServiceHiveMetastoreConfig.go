// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceHiveMetastoreConfig struct {
	// The Hive metastore schema version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/dataproc_metastore_service#version DataprocMetastoreService#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// auxiliary_versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/dataproc_metastore_service#auxiliary_versions DataprocMetastoreService#auxiliary_versions}
	AuxiliaryVersions interface{} `field:"optional" json:"auxiliaryVersions" yaml:"auxiliaryVersions"`
	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml). The mappings override system defaults (some keys cannot be overridden).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/dataproc_metastore_service#config_overrides DataprocMetastoreService#config_overrides}
	ConfigOverrides *map[string]*string `field:"optional" json:"configOverrides" yaml:"configOverrides"`
	// The protocol to use for the metastore service endpoint.
	//
	// If unspecified, defaults to 'THRIFT'. Default value: "THRIFT" Possible values: ["THRIFT", "GRPC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/dataproc_metastore_service#endpoint_protocol DataprocMetastoreService#endpoint_protocol}
	EndpointProtocol *string `field:"optional" json:"endpointProtocol" yaml:"endpointProtocol"`
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/dataproc_metastore_service#kerberos_config DataprocMetastoreService#kerberos_config}
	KerberosConfig *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig `field:"optional" json:"kerberosConfig" yaml:"kerberosConfig"`
}

