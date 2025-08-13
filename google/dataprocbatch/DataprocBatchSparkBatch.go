// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch


type DataprocBatchSparkBatch struct {
	// HCFS URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_batch#archive_uris DataprocBatch#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// The arguments to pass to the driver.
	//
	// Do not include arguments that can be set as batch
	// properties, such as --conf, since a collision can occur that causes an incorrect batch submission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_batch#args DataprocBatch#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// HCFS URIs of files to be placed in the working directory of each executor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_batch#file_uris DataprocBatch#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// HCFS URIs of jar files to add to the classpath of the Spark driver and tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_batch#jar_file_uris DataprocBatch#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// The name of the driver main class.
	//
	// The jar file that contains the class must be in the
	// classpath or specified in jarFileUris.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_batch#main_class DataprocBatch#main_class}
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// The HCFS URI of the jar file that contains the main class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_batch#main_jar_file_uri DataprocBatch#main_jar_file_uri}
	MainJarFileUri *string `field:"optional" json:"mainJarFileUri" yaml:"mainJarFileUri"`
}

