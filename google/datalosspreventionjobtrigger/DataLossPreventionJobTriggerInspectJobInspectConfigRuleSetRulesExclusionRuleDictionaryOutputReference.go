// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/datalosspreventionjobtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference interface {
	cdktf.ComplexObject
	CloudStoragePath() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathOutputReference
	CloudStoragePathInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
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
	InternalValue() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	SetInternalValue(val *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WordList() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStructOutputReference
	WordListInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStruct
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
	PutCloudStoragePath(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
	PutWordList(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStruct)
	ResetCloudStoragePath()
	ResetWordList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
type jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) CloudStoragePath() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathOutputReference
	_jsii_.Get(
		j,
		"cloudStoragePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) CloudStoragePathInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	_jsii_.Get(
		j,
		"cloudStoragePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) InternalValue() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) WordList() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStructOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStructOutputReference
	_jsii_.Get(
		j,
		"wordList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) WordListInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStruct {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStruct
	_jsii_.Get(
		j,
		"wordListInput",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference_Override(d DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference)SetInternalValue(val *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) PutCloudStoragePath(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) {
	if err := d.validatePutCloudStoragePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudStoragePath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) PutWordList(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryWordListStruct) {
	if err := d.validatePutWordListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWordList",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) ResetCloudStoragePath() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStoragePath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) ResetWordList() {
	_jsii_.InvokeVoid(
		d,
		"resetWordList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

