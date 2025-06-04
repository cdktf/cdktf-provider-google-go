// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiindexendpointdeployedindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/vertexaiindexendpointdeployedindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference interface {
	cdktf.ComplexObject
	AllowedIssuers() *[]*string
	SetAllowedIssuers(val *[]*string)
	AllowedIssuersInput() *[]*string
	Audiences() *[]*string
	SetAudiences(val *[]*string)
	AudiencesInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProvider
	SetInternalValue(val *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProvider)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAllowedIssuers()
	ResetAudiences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference
type jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) AllowedIssuers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIssuers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) AllowedIssuersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIssuersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) Audiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) AudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) InternalValue() *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProvider {
	var returns *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProvider
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference_Override(v VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiIndexEndpointDeployedIndex.VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference)SetAllowedIssuers(val *[]*string) {
	if err := j.validateSetAllowedIssuersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIssuers",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference)SetAudiences(val *[]*string) {
	if err := j.validateSetAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audiences",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference)SetInternalValue(val *VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProvider) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) ResetAllowedIssuers() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowedIssuers",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) ResetAudiences() {
	_jsii_.InvokeVoid(
		v,
		"resetAudiences",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

