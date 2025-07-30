// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplateEncryptionConfig struct {
	// Optional. The Cloud KMS key name to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dataproc_workflow_template#kms_key DataprocWorkflowTemplate#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

