// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstancePersistenceConfig struct {
	// aof_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/memorystore_instance#aof_config MemorystoreInstance#aof_config}
	AofConfig *MemorystoreInstancePersistenceConfigAofConfig `field:"optional" json:"aofConfig" yaml:"aofConfig"`
	// Optional. Current persistence mode.   Possible values: DISABLED RDB AOF Possible values: ["DISABLED", "RDB", "AOF"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/memorystore_instance#mode MemorystoreInstance#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// rdb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/memorystore_instance#rdb_config MemorystoreInstance#rdb_config}
	RdbConfig *MemorystoreInstancePersistenceConfigRdbConfig `field:"optional" json:"rdbConfig" yaml:"rdbConfig"`
}

