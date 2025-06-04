// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/vertexaiendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiEndpointPrivateServiceConnectConfigOutputReference interface {
	cdktf.ComplexObject
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
	EnablePrivateServiceConnect() interface{}
	SetEnablePrivateServiceConnect(val interface{})
	EnablePrivateServiceConnectInput() interface{}
	EnableSecurePrivateServiceConnect() interface{}
	SetEnableSecurePrivateServiceConnect(val interface{})
	EnableSecurePrivateServiceConnectInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiEndpointPrivateServiceConnectConfig
	SetInternalValue(val *VertexAiEndpointPrivateServiceConnectConfig)
	ProjectAllowlist() *[]*string
	SetProjectAllowlist(val *[]*string)
	ProjectAllowlistInput() *[]*string
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
	ResetEnableSecurePrivateServiceConnect()
	ResetProjectAllowlist()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiEndpointPrivateServiceConnectConfigOutputReference
type jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) EnablePrivateServiceConnect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateServiceConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) EnablePrivateServiceConnectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateServiceConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) EnableSecurePrivateServiceConnect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSecurePrivateServiceConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) EnableSecurePrivateServiceConnectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSecurePrivateServiceConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) InternalValue() *VertexAiEndpointPrivateServiceConnectConfig {
	var returns *VertexAiEndpointPrivateServiceConnectConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) ProjectAllowlist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) ProjectAllowlistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiEndpointPrivateServiceConnectConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiEndpointPrivateServiceConnectConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiEndpointPrivateServiceConnectConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpoint.VertexAiEndpointPrivateServiceConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiEndpointPrivateServiceConnectConfigOutputReference_Override(v VertexAiEndpointPrivateServiceConnectConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiEndpoint.VertexAiEndpointPrivateServiceConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference)SetEnablePrivateServiceConnect(val interface{}) {
	if err := j.validateSetEnablePrivateServiceConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateServiceConnect",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference)SetEnableSecurePrivateServiceConnect(val interface{}) {
	if err := j.validateSetEnableSecurePrivateServiceConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSecurePrivateServiceConnect",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference)SetInternalValue(val *VertexAiEndpointPrivateServiceConnectConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference)SetProjectAllowlist(val *[]*string) {
	if err := j.validateSetProjectAllowlistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectAllowlist",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) ResetEnableSecurePrivateServiceConnect() {
	_jsii_.InvokeVoid(
		v,
		"resetEnableSecurePrivateServiceConnect",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) ResetProjectAllowlist() {
	_jsii_.InvokeVoid(
		v,
		"resetProjectAllowlist",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiEndpointPrivateServiceConnectConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

