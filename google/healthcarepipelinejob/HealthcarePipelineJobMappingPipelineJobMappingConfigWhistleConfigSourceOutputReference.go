// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/healthcarepipelinejob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference interface {
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
	// Experimental.
	Fqn() *string
	ImportUriPrefix() *string
	SetImportUriPrefix(val *string)
	ImportUriPrefixInput() *string
	InternalValue() *HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSource
	SetInternalValue(val *HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSource)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference
type jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) ImportUriPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUriPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) ImportUriPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUriPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) InternalValue() *HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSource {
	var returns *HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference {
	_init_.Initialize()

	if err := validateNewHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcarePipelineJob.HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference_Override(h HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcarePipelineJob.HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference)SetImportUriPrefix(val *string) {
	if err := j.validateSetImportUriPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importUriPrefix",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference)SetInternalValue(val *HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

