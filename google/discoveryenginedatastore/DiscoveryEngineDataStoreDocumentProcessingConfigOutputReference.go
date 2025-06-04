// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginedatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/discoveryenginedatastore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference interface {
	cdktf.ComplexObject
	ChunkingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfigOutputReference
	ChunkingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfig
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
	DefaultParsingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference
	DefaultParsingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig
	// Experimental.
	Fqn() *string
	InternalValue() *DiscoveryEngineDataStoreDocumentProcessingConfig
	SetInternalValue(val *DiscoveryEngineDataStoreDocumentProcessingConfig)
	Name() *string
	ParsingConfigOverrides() DiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesList
	ParsingConfigOverridesInput() interface{}
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
	PutChunkingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfig)
	PutDefaultParsingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig)
	PutParsingConfigOverrides(value interface{})
	ResetChunkingConfig()
	ResetDefaultParsingConfig()
	ResetParsingConfigOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference
type jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ChunkingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfigOutputReference {
	var returns DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfigOutputReference
	_jsii_.Get(
		j,
		"chunkingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ChunkingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfig {
	var returns *DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfig
	_jsii_.Get(
		j,
		"chunkingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) DefaultParsingConfig() DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference {
	var returns DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfigOutputReference
	_jsii_.Get(
		j,
		"defaultParsingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) DefaultParsingConfigInput() *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig {
	var returns *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig
	_jsii_.Get(
		j,
		"defaultParsingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) InternalValue() *DiscoveryEngineDataStoreDocumentProcessingConfig {
	var returns *DiscoveryEngineDataStoreDocumentProcessingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ParsingConfigOverrides() DiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesList {
	var returns DiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesList
	_jsii_.Get(
		j,
		"parsingConfigOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ParsingConfigOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parsingConfigOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineDataStoreDocumentProcessingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineDataStoreDocumentProcessingConfigOutputReference_Override(d DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineDataStore.DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetInternalValue(val *DiscoveryEngineDataStoreDocumentProcessingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) PutChunkingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfig) {
	if err := d.validatePutChunkingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChunkingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) PutDefaultParsingConfig(value *DiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig) {
	if err := d.validatePutDefaultParsingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDefaultParsingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) PutParsingConfigOverrides(value interface{}) {
	if err := d.validatePutParsingConfigOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putParsingConfigOverrides",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ResetChunkingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetChunkingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ResetDefaultParsingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultParsingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ResetParsingConfigOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetParsingConfigOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DiscoveryEngineDataStoreDocumentProcessingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

