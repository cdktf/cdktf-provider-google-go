// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob


type StorageBatchOperationsJobBucketListBucketsManifest struct {
	// Specifies objects in a manifest file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/storage_batch_operations_job#manifest_location StorageBatchOperationsJob#manifest_location}
	ManifestLocation *string `field:"optional" json:"manifestLocation" yaml:"manifestLocation"`
}

