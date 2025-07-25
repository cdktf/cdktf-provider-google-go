// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/healthcarefhirstore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference interface {
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
	InternalValue() *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig
	SetInternalValue(val *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig)
	LastUpdatedPartitionConfig() HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfigOutputReference
	LastUpdatedPartitionConfigInput() *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfig
	RecursiveStructureDepth() *float64
	SetRecursiveStructureDepth(val *float64)
	RecursiveStructureDepthInput() *float64
	SchemaType() *string
	SetSchemaType(val *string)
	SchemaTypeInput() *string
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
	PutLastUpdatedPartitionConfig(value *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfig)
	ResetLastUpdatedPartitionConfig()
	ResetSchemaType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference
type jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) InternalValue() *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig {
	var returns *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) LastUpdatedPartitionConfig() HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfigOutputReference {
	var returns HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfigOutputReference
	_jsii_.Get(
		j,
		"lastUpdatedPartitionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) LastUpdatedPartitionConfigInput() *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfig {
	var returns *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfig
	_jsii_.Get(
		j,
		"lastUpdatedPartitionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) RecursiveStructureDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recursiveStructureDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) RecursiveStructureDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recursiveStructureDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) SchemaType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) SchemaTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference {
	_init_.Initialize()

	if err := validateNewHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference_Override(h HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference)SetInternalValue(val *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference)SetRecursiveStructureDepth(val *float64) {
	if err := j.validateSetRecursiveStructureDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recursiveStructureDepth",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference)SetSchemaType(val *string) {
	if err := j.validateSetSchemaTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaType",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) PutLastUpdatedPartitionConfig(value *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfig) {
	if err := h.validatePutLastUpdatedPartitionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putLastUpdatedPartitionConfig",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) ResetLastUpdatedPartitionConfig() {
	_jsii_.InvokeVoid(
		h,
		"resetLastUpdatedPartitionConfig",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) ResetSchemaType() {
	_jsii_.InvokeVoid(
		h,
		"resetSchemaType",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

