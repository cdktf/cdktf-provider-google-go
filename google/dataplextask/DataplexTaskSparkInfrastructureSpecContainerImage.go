package dataplextask


type DataplexTaskSparkInfrastructureSpecContainerImage struct {
	// Container image to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#image DataplexTask#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// A list of Java JARS to add to the classpath.
	//
	// Valid input includes Cloud Storage URIs to Jar binaries. For example, gs://bucket-name/my/path/to/file.jar
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#java_jars DataplexTask#java_jars}
	JavaJars *[]*string `field:"optional" json:"javaJars" yaml:"javaJars"`
	// Override to common configuration of open source components installed on the Dataproc cluster.
	//
	// The properties to set on daemon config files. Property keys are specified in prefix:property format, for example core:hadoop.tmp.dir. For more information, see Cluster properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#properties DataplexTask#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// A list of python packages to be installed.
	//
	// Valid formats include Cloud Storage URI to a PIP installable library. For example, gs://bucket-name/my/path/to/lib.tar.gz
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#python_packages DataplexTask#python_packages}
	PythonPackages *[]*string `field:"optional" json:"pythonPackages" yaml:"pythonPackages"`
}

