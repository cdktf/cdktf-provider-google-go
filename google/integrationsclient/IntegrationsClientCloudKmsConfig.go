// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsclient


type IntegrationsClientCloudKmsConfig struct {
	// A Cloud KMS key is a named object containing one or more key versions, along with metadata for the key.
	//
	// A key exists on exactly one key ring tied to a
	// specific location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/integrations_client#key IntegrationsClient#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Location name of the key ring, e.g. "us-west1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/integrations_client#kms_location IntegrationsClient#kms_location}
	KmsLocation *string `field:"required" json:"kmsLocation" yaml:"kmsLocation"`
	// A key ring organizes keys in a specific Google Cloud location and allows you to manage access control on groups of keys.
	//
	// A key ring's name does not need to be
	// unique across a Google Cloud project, but must be unique within a given location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/integrations_client#kms_ring IntegrationsClient#kms_ring}
	KmsRing *string `field:"required" json:"kmsRing" yaml:"kmsRing"`
	// Each version of a key contains key material used for encryption or signing.
	//
	// A key's version is represented by an integer, starting at 1. To decrypt data
	// or verify a signature, you must use the same key version that was used to
	// encrypt or sign the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/integrations_client#key_version IntegrationsClient#key_version}
	KeyVersion *string `field:"optional" json:"keyVersion" yaml:"keyVersion"`
	// The Google Cloud project id of the project where the kms key stored.
	//
	// If empty,
	// the kms key is stored at the same project as customer's project and ecrypted
	// with CMEK, otherwise, the kms key is stored in the tenant project and
	// encrypted with GMEK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/integrations_client#kms_project_id IntegrationsClient#kms_project_id}
	KmsProjectId *string `field:"optional" json:"kmsProjectId" yaml:"kmsProjectId"`
}

