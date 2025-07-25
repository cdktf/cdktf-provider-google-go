// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/healthcarepipelinejob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcarePipelineJobMappingPipelineJobOutputReference interface {
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
	FhirStoreDestination() *string
	SetFhirStoreDestination(val *string)
	FhirStoreDestinationInput() *string
	FhirStreamingSource() HealthcarePipelineJobMappingPipelineJobFhirStreamingSourceOutputReference
	FhirStreamingSourceInput() *HealthcarePipelineJobMappingPipelineJobFhirStreamingSource
	// Experimental.
	Fqn() *string
	InternalValue() *HealthcarePipelineJobMappingPipelineJob
	SetInternalValue(val *HealthcarePipelineJobMappingPipelineJob)
	MappingConfig() HealthcarePipelineJobMappingPipelineJobMappingConfigOutputReference
	MappingConfigInput() *HealthcarePipelineJobMappingPipelineJobMappingConfig
	ReconciliationDestination() interface{}
	SetReconciliationDestination(val interface{})
	ReconciliationDestinationInput() interface{}
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
	PutFhirStreamingSource(value *HealthcarePipelineJobMappingPipelineJobFhirStreamingSource)
	PutMappingConfig(value *HealthcarePipelineJobMappingPipelineJobMappingConfig)
	ResetFhirStoreDestination()
	ResetFhirStreamingSource()
	ResetReconciliationDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcarePipelineJobMappingPipelineJobOutputReference
type jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) FhirStoreDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fhirStoreDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) FhirStoreDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fhirStoreDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) FhirStreamingSource() HealthcarePipelineJobMappingPipelineJobFhirStreamingSourceOutputReference {
	var returns HealthcarePipelineJobMappingPipelineJobFhirStreamingSourceOutputReference
	_jsii_.Get(
		j,
		"fhirStreamingSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) FhirStreamingSourceInput() *HealthcarePipelineJobMappingPipelineJobFhirStreamingSource {
	var returns *HealthcarePipelineJobMappingPipelineJobFhirStreamingSource
	_jsii_.Get(
		j,
		"fhirStreamingSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) InternalValue() *HealthcarePipelineJobMappingPipelineJob {
	var returns *HealthcarePipelineJobMappingPipelineJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) MappingConfig() HealthcarePipelineJobMappingPipelineJobMappingConfigOutputReference {
	var returns HealthcarePipelineJobMappingPipelineJobMappingConfigOutputReference
	_jsii_.Get(
		j,
		"mappingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) MappingConfigInput() *HealthcarePipelineJobMappingPipelineJobMappingConfig {
	var returns *HealthcarePipelineJobMappingPipelineJobMappingConfig
	_jsii_.Get(
		j,
		"mappingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ReconciliationDestination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reconciliationDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ReconciliationDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reconciliationDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthcarePipelineJobMappingPipelineJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcarePipelineJobMappingPipelineJobOutputReference {
	_init_.Initialize()

	if err := validateNewHealthcarePipelineJobMappingPipelineJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcarePipelineJob.HealthcarePipelineJobMappingPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcarePipelineJobMappingPipelineJobOutputReference_Override(h HealthcarePipelineJobMappingPipelineJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcarePipelineJob.HealthcarePipelineJobMappingPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference)SetFhirStoreDestination(val *string) {
	if err := j.validateSetFhirStoreDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fhirStoreDestination",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference)SetInternalValue(val *HealthcarePipelineJobMappingPipelineJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference)SetReconciliationDestination(val interface{}) {
	if err := j.validateSetReconciliationDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reconciliationDestination",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) PutFhirStreamingSource(value *HealthcarePipelineJobMappingPipelineJobFhirStreamingSource) {
	if err := h.validatePutFhirStreamingSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putFhirStreamingSource",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) PutMappingConfig(value *HealthcarePipelineJobMappingPipelineJobMappingConfig) {
	if err := h.validatePutMappingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putMappingConfig",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ResetFhirStoreDestination() {
	_jsii_.InvokeVoid(
		h,
		"resetFhirStoreDestination",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ResetFhirStreamingSource() {
	_jsii_.InvokeVoid(
		h,
		"resetFhirStreamingSource",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ResetReconciliationDestination() {
	_jsii_.InvokeVoid(
		h,
		"resetReconciliationDestination",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HealthcarePipelineJobMappingPipelineJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

