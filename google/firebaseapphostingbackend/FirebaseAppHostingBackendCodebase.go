// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingbackend


type FirebaseAppHostingBackendCodebase struct {
	// The resource name for the Developer Connect ['gitRepositoryLink'](https://cloud.google.com/developer-connect/docs/api/reference/rest/v1/projects.locations.connections.gitRepositoryLinks) connected to this backend, in the format:.
	//
	// projects/{project}/locations/{location}/connections/{connection}/gitRepositoryLinks/{repositoryLink}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firebase_app_hosting_backend#repository FirebaseAppHostingBackend#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// If 'repository' is provided, the directory relative to the root of the repository to use as the root for the deployed web app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firebase_app_hosting_backend#root_directory FirebaseAppHostingBackend#root_directory}
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
}

