// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/healthcarepipelinejob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference interface {
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
	InternalValue() *HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource
	SetInternalValue(val *HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource)
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

// The jsii proxy struct for HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference
type jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ImportUriPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUriPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ImportUriPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUriPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) InternalValue() *HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource {
	var returns *HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference {
	_init_.Initialize()

	if err := validateNewHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcarePipelineJob.HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference_Override(h HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcarePipelineJob.HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetImportUriPrefix(val *string) {
	if err := j.validateSetImportUriPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importUriPrefix",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetInternalValue(val *HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

