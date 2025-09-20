// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch


type DataprocBatchEnvironmentConfigExecutionConfigAuthenticationConfig struct {
	// Authentication type for the user workload running in containers. Possible values: ["SERVICE_ACCOUNT", "END_USER_CREDENTIALS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dataproc_batch#user_workload_authentication_type DataprocBatch#user_workload_authentication_type}
	UserWorkloadAuthenticationType *string `field:"optional" json:"userWorkloadAuthenticationType" yaml:"userWorkloadAuthenticationType"`
}

