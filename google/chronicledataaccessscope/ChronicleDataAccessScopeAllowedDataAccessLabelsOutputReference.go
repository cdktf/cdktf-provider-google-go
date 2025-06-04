// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicledataaccessscope

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/chronicledataaccessscope/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference interface {
	cdktf.ComplexObject
	AssetNamespace() *string
	SetAssetNamespace(val *string)
	AssetNamespaceInput() *string
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
	DataAccessLabel() *string
	SetDataAccessLabel(val *string)
	DataAccessLabelInput() *string
	DisplayName() *string
	// Experimental.
	Fqn() *string
	IngestionLabel() ChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelOutputReference
	IngestionLabelInput() *ChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogType() *string
	SetLogType(val *string)
	LogTypeInput() *string
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
	PutIngestionLabel(value *ChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel)
	ResetAssetNamespace()
	ResetDataAccessLabel()
	ResetIngestionLabel()
	ResetLogType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference
type jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) AssetNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) AssetNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) DataAccessLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) DataAccessLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) IngestionLabel() ChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelOutputReference {
	var returns ChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabelOutputReference
	_jsii_.Get(
		j,
		"ingestionLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) IngestionLabelInput() *ChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel {
	var returns *ChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel
	_jsii_.Get(
		j,
		"ingestionLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) LogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDataAccessScopeAllowedDataAccessLabelsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDataAccessScope.ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference_Override(c ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDataAccessScope.ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference)SetAssetNamespace(val *string) {
	if err := j.validateSetAssetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetNamespace",
		val,
	)
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference)SetDataAccessLabel(val *string) {
	if err := j.validateSetDataAccessLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessLabel",
		val,
	)
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference)SetLogType(val *string) {
	if err := j.validateSetLogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logType",
		val,
	)
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) PutIngestionLabel(value *ChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel) {
	if err := c.validatePutIngestionLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIngestionLabel",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) ResetAssetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetAssetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) ResetDataAccessLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetDataAccessLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) ResetIngestionLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetIngestionLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) ResetLogType() {
	_jsii_.InvokeVoid(
		c,
		"resetLogType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDataAccessScopeAllowedDataAccessLabelsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

