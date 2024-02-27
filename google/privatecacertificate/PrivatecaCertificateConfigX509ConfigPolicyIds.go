// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificate


type PrivatecaCertificateConfigX509ConfigPolicyIds struct {
	// An ObjectId specifies an object identifier (OID). These provide context and describe types in ASN.1 messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/privateca_certificate#object_id_path PrivatecaCertificate#object_id_path}
	ObjectIdPath *[]*float64 `field:"required" json:"objectIdPath" yaml:"objectIdPath"`
}

