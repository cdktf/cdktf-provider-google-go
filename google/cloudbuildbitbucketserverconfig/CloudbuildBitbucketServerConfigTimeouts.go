// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildbitbucketserverconfig


type CloudbuildBitbucketServerConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/cloudbuild_bitbucket_server_config#create CloudbuildBitbucketServerConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/cloudbuild_bitbucket_server_config#delete CloudbuildBitbucketServerConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/cloudbuild_bitbucket_server_config#update CloudbuildBitbucketServerConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

