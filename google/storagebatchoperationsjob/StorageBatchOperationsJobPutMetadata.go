// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob


type StorageBatchOperationsJobPutMetadata struct {
	// Cache-Control directive to specify caching behavior of object data.
	//
	// If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#cache_control StorageBatchOperationsJob#cache_control}
	CacheControl *string `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Content-Disposition of the object data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#content_disposition StorageBatchOperationsJob#content_disposition}
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// Content Encoding of the object data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#content_encoding StorageBatchOperationsJob#content_encoding}
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// Content-Language of the object data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#content_language StorageBatchOperationsJob#content_language}
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// Content-Type of the object data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#content_type StorageBatchOperationsJob#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// User-provided metadata, in key/value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#custom_metadata StorageBatchOperationsJob#custom_metadata}
	CustomMetadata *map[string]*string `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// Updates the objects fixed custom time metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_batch_operations_job#custom_time StorageBatchOperationsJob#custom_time}
	CustomTime *string `field:"optional" json:"customTime" yaml:"customTime"`
}

