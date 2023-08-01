package containeranalysisoccurrence

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAnalysisOccurrenceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// attestation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_analysis_occurrence#attestation ContainerAnalysisOccurrence#attestation}
	Attestation *ContainerAnalysisOccurrenceAttestation `field:"required" json:"attestation" yaml:"attestation"`
	// The analysis note associated with this occurrence, in the form of projects/[PROJECT]/notes/[NOTE_ID].
	//
	// This field can be used as a
	// filter in list requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_analysis_occurrence#note_name ContainerAnalysisOccurrence#note_name}
	NoteName *string `field:"required" json:"noteName" yaml:"noteName"`
	// Required. Immutable. A URI that represents the resource for which the occurrence applies. For example, https://gcr.io/project/image@sha256:123abc for a Docker image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_analysis_occurrence#resource_uri ContainerAnalysisOccurrence#resource_uri}
	ResourceUri *string `field:"required" json:"resourceUri" yaml:"resourceUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_analysis_occurrence#id ContainerAnalysisOccurrence#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_analysis_occurrence#project ContainerAnalysisOccurrence#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A description of actions that can be taken to remedy the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_analysis_occurrence#remediation ContainerAnalysisOccurrence#remediation}
	Remediation *string `field:"optional" json:"remediation" yaml:"remediation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_analysis_occurrence#timeouts ContainerAnalysisOccurrence#timeouts}
	Timeouts *ContainerAnalysisOccurrenceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

