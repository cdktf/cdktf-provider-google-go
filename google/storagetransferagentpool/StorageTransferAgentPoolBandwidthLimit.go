// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferagentpool


type StorageTransferAgentPoolBandwidthLimit struct {
	// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/storage_transfer_agent_pool#limit_mbps StorageTransferAgentPool#limit_mbps}
	LimitMbps *string `field:"required" json:"limitMbps" yaml:"limitMbps"`
}

