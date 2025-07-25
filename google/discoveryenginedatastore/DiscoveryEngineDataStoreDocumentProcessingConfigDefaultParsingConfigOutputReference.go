// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginedatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/discoveryenginedatastore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference interface {
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
	DigitalParsingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfigOutputReference
	DigitalParsingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfig
	// Experimental.
	Fqn() *string
	InternalValue() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig
	SetInternalValue(val *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig)
	LayoutParsingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfigOutputReference
	LayoutParsingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig
	OcrParsingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfigOutputReference
	OcrParsingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfig
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
	PutDigitalParsingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfig)
	PutLayoutParsingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig)
	PutOcrParsingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfig)
	ResetDigitalParsingConfig()
	ResetLayoutParsingConfig()
	ResetOcrParsingConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference
type jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) DigitalParsingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfigOutputReference {
	var returns DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfigOutputReference
	_jsii_.Get(
		j,
		"digitalParsingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) DigitalParsingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfig {
	var returns *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfig
	_jsii_.Get(
		j,
		"digitalParsingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) InternalValue() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig {
	var returns *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) LayoutParsingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfigOutputReference {
	var returns DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfigOutputReference
	_jsii_.Get(
		j,
		"layoutParsingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) LayoutParsingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig {
	var returns *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig
	_jsii_.Get(
		j,
		"layoutParsingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) OcrParsingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfigOutputReference {
	var returns DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfigOutputReference
	_jsii_.Get(
		j,
		"ocrParsingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) OcrParsingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfig {
	var returns *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfig
	_jsii_.Get(
		j,
		"ocrParsingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference_Override(d DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference)SetInternalValue(val *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) PutDigitalParsingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigDigitalParsingConfig) {
	if err := d.validatePutDigitalParsingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDigitalParsingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) PutLayoutParsingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigLayoutParsingConfig) {
	if err := d.validatePutLayoutParsingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLayoutParsingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) PutOcrParsingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOcrParsingConfig) {
	if err := d.validatePutOcrParsingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOcrParsingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) ResetDigitalParsingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetDigitalParsingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) ResetLayoutParsingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLayoutParsingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) ResetOcrParsingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetOcrParsingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

