// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerrepository


type SecureSourceManagerRepositoryInitialConfig struct {
	// Default branch name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/secure_source_manager_repository#default_branch SecureSourceManagerRepository#default_branch}
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// List of gitignore template names user can choose from. Valid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/secure_source_manager_repository#gitignores SecureSourceManagerRepository#gitignores}
	Gitignores *[]*string `field:"optional" json:"gitignores" yaml:"gitignores"`
	// License template name user can choose from. Valid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/secure_source_manager_repository#license SecureSourceManagerRepository#license}
	License *string `field:"optional" json:"license" yaml:"license"`
	// README template name. Valid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/secure_source_manager_repository#readme SecureSourceManagerRepository#readme}
	Readme *string `field:"optional" json:"readme" yaml:"readme"`
}

