// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/healthcarepipelinejob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcarePipelineJobReconciliationPipelineJobOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *HealthcarePipelineJobReconciliationPipelineJob
	SetInternalValue(val *HealthcarePipelineJobReconciliationPipelineJob)
	MatchingUriPrefix() *string
	SetMatchingUriPrefix(val *string)
	MatchingUriPrefixInput() *string
	MergeConfig() HealthcarePipelineJobReconciliationPipelineJobMergeConfigOutputReference
	MergeConfigInput() *HealthcarePipelineJobReconciliationPipelineJobMergeConfig
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
	PutMergeConfig(value *HealthcarePipelineJobReconciliationPipelineJobMergeConfig)
	ResetFhirStoreDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcarePipelineJobReconciliationPipelineJobOutputReference
type jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) FhirStoreDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fhirStoreDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) FhirStoreDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fhirStoreDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) InternalValue() *HealthcarePipelineJobReconciliationPipelineJob {
	var returns *HealthcarePipelineJobReconciliationPipelineJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) MatchingUriPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingUriPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) MatchingUriPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingUriPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) MergeConfig() HealthcarePipelineJobReconciliationPipelineJobMergeConfigOutputReference {
	var returns HealthcarePipelineJobReconciliationPipelineJobMergeConfigOutputReference
	_jsii_.Get(
		j,
		"mergeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) MergeConfigInput() *HealthcarePipelineJobReconciliationPipelineJobMergeConfig {
	var returns *HealthcarePipelineJobReconciliationPipelineJobMergeConfig
	_jsii_.Get(
		j,
		"mergeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthcarePipelineJobReconciliationPipelineJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcarePipelineJobReconciliationPipelineJobOutputReference {
	_init_.Initialize()

	if err := validateNewHealthcarePipelineJobReconciliationPipelineJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcarePipelineJob.HealthcarePipelineJobReconciliationPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcarePipelineJobReconciliationPipelineJobOutputReference_Override(h HealthcarePipelineJobReconciliationPipelineJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcarePipelineJob.HealthcarePipelineJobReconciliationPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference)SetFhirStoreDestination(val *string) {
	if err := j.validateSetFhirStoreDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fhirStoreDestination",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference)SetInternalValue(val *HealthcarePipelineJobReconciliationPipelineJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference)SetMatchingUriPrefix(val *string) {
	if err := j.validateSetMatchingUriPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingUriPrefix",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) PutMergeConfig(value *HealthcarePipelineJobReconciliationPipelineJobMergeConfig) {
	if err := h.validatePutMergeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putMergeConfig",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) ResetFhirStoreDestination() {
	_jsii_.InvokeVoid(
		h,
		"resetFhirStoreDestination",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HealthcarePipelineJobReconciliationPipelineJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

