// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacapool


type PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageUnknownExtendedKeyUsages struct {
	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#object_id_path PrivatecaCaPool#object_id_path}
	ObjectIdPath *[]*float64 `field:"required" json:"objectIdPath" yaml:"objectIdPath"`
}

