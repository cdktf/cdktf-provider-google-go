// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtableschemabundle


type BigtableSchemaBundleProtoSchema struct {
	// Base64 encoded content of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/bigtable_schema_bundle#proto_descriptors BigtableSchemaBundle#proto_descriptors}
	ProtoDescriptors *string `field:"required" json:"protoDescriptors" yaml:"protoDescriptors"`
}

