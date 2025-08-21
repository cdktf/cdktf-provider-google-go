// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement


type PrivilegedAccessManagerEntitlementAdditionalNotificationTargets struct {
	// Optional. Additional email addresses to be notified when a principal(requester) is granted access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/privileged_access_manager_entitlement#admin_email_recipients PrivilegedAccessManagerEntitlement#admin_email_recipients}
	AdminEmailRecipients *[]*string `field:"optional" json:"adminEmailRecipients" yaml:"adminEmailRecipients"`
	// Optional. Additional email address to be notified about an eligible entitlement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/privileged_access_manager_entitlement#requester_email_recipients PrivilegedAccessManagerEntitlement#requester_email_recipients}
	RequesterEmailRecipients *[]*string `field:"optional" json:"requesterEmailRecipients" yaml:"requesterEmailRecipients"`
}

