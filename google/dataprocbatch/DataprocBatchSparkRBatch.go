// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch


type DataprocBatchSparkRBatch struct {
	// HCFS URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/dataproc_batch#archive_uris DataprocBatch#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// The arguments to pass to the driver.
	//
	// Do not include arguments that can be set as batch
	// properties, such as --conf, since a collision can occur that causes an incorrect batch submission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/dataproc_batch#args DataprocBatch#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// HCFS URIs of files to be placed in the working directory of each executor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/dataproc_batch#file_uris DataprocBatch#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// The HCFS URI of the main R file to use as the driver.
	//
	// Must be a .R or .r file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/dataproc_batch#main_r_file_uri DataprocBatch#main_r_file_uri}
	MainRFileUri *string `field:"optional" json:"mainRFileUri" yaml:"mainRFileUri"`
}

