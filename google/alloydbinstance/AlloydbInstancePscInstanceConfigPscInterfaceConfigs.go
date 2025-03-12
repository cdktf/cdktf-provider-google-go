// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstancePscInstanceConfigPscInterfaceConfigs struct {
	// The network attachment resource created in the consumer project to which the PSC interface will be linked.
	//
	// This is of the format: "projects/${CONSUMER_PROJECT}/regions/${REGION}/networkAttachments/${NETWORK_ATTACHMENT_NAME}".
	// The network attachment must be in the same region as the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/alloydb_instance#network_attachment_resource AlloydbInstance#network_attachment_resource}
	NetworkAttachmentResource *string `field:"optional" json:"networkAttachmentResource" yaml:"networkAttachmentResource"`
}

