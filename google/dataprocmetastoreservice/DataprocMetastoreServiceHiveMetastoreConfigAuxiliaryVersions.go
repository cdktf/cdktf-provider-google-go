// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceHiveMetastoreConfigAuxiliaryVersions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_metastore_service#key DataprocMetastoreService#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Hive metastore version of the auxiliary service. It must be less than the primary Hive metastore service's version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_metastore_service#version DataprocMetastoreService#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// A mapping of Hive metastore configuration key-value pairs to apply to the auxiliary Hive metastore (configured in hive-site.xml) in addition to the primary version's overrides. If keys are present in both the auxiliary version's overrides and the primary version's overrides, the value from the auxiliary version's overrides takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/dataproc_metastore_service#config_overrides DataprocMetastoreService#config_overrides}
	ConfigOverrides *map[string]*string `field:"optional" json:"configOverrides" yaml:"configOverrides"`
}

