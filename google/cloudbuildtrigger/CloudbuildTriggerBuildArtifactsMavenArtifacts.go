// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger


type CloudbuildTriggerBuildArtifactsMavenArtifacts struct {
	// Maven artifactId value used when uploading the artifact to Artifact Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/cloudbuild_trigger#artifact_id CloudbuildTrigger#artifact_id}
	ArtifactId *string `field:"optional" json:"artifactId" yaml:"artifactId"`
	// Maven groupId value used when uploading the artifact to Artifact Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/cloudbuild_trigger#group_id CloudbuildTrigger#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Path to an artifact in the build's workspace to be uploaded to Artifact Registry.
	//
	// This can be either an absolute path, e.g. /workspace/my-app/target/my-app-1.0.SNAPSHOT.jar or a relative path from /workspace, e.g. my-app/target/my-app-1.0.SNAPSHOT.jar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/cloudbuild_trigger#path CloudbuildTrigger#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Artifact Registry repository, in the form "https://$REGION-maven.pkg.dev/$PROJECT/$REPOSITORY".
	//
	// Artifact in the workspace specified by path will be uploaded to Artifact Registry with this location as a prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/cloudbuild_trigger#repository CloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Maven version value used when uploading the artifact to Artifact Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/cloudbuild_trigger#version CloudbuildTrigger#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

