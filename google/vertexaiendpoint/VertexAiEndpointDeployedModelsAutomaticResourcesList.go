package vertexaiendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/vertexaiendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiEndpointDeployedModelsAutomaticResourcesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) VertexAiEndpointDeployedModelsAutomaticResourcesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiEndpointDeployedModelsAutomaticResourcesList
type jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVertexAiEndpointDeployedModelsAutomaticResourcesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VertexAiEndpointDeployedModelsAutomaticResourcesList {
	_init_.Initialize()

	if err := validateNewVertexAiEndpointDeployedModelsAutomaticResourcesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpoint.VertexAiEndpointDeployedModelsAutomaticResourcesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVertexAiEndpointDeployedModelsAutomaticResourcesList_Override(v VertexAiEndpointDeployedModelsAutomaticResourcesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpoint.VertexAiEndpointDeployedModelsAutomaticResourcesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) Get(index *float64) VertexAiEndpointDeployedModelsAutomaticResourcesOutputReference {
	if err := v.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns VertexAiEndpointDeployedModelsAutomaticResourcesOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointDeployedModelsAutomaticResourcesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

