// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lookerinstance


type LookerInstancePscConfigServiceAttachments struct {
	// Fully qualified domain name that will be used in the private DNS record created for the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/looker_instance#local_fqdn LookerInstance#local_fqdn}
	LocalFqdn *string `field:"optional" json:"localFqdn" yaml:"localFqdn"`
	// URI of the service attachment to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/looker_instance#target_service_attachment_uri LookerInstance#target_service_attachment_uri}
	TargetServiceAttachmentUri *string `field:"optional" json:"targetServiceAttachmentUri" yaml:"targetServiceAttachmentUri"`
}

