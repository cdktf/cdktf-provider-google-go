package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/datalosspreventionjobtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference interface {
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
	Dictionary() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
	DictionaryInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	ExcludeByHotword() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference
	ExcludeByHotwordInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword
	ExcludeInfoTypes() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference
	ExcludeInfoTypesInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	// Experimental.
	Fqn() *string
	InternalValue() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	SetInternalValue(val *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule)
	MatchingType() *string
	SetMatchingType(val *string)
	MatchingTypeInput() *string
	Regex() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexOutputReference
	RegexInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
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
	PutDictionary(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary)
	PutExcludeByHotword(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword)
	PutExcludeInfoTypes(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	PutRegex(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex)
	ResetDictionary()
	ResetExcludeByHotword()
	ResetExcludeInfoTypes()
	ResetRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference
type jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) Dictionary() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) DictionaryInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByHotword() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordOutputReference
	_jsii_.Get(
		j,
		"excludeByHotword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeByHotwordInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword
	_jsii_.Get(
		j,
		"excludeByHotwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeInfoTypes() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesOutputReference
	_jsii_.Get(
		j,
		"excludeInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ExcludeInfoTypesInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	_jsii_.Get(
		j,
		"excludeInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) InternalValue() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) MatchingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) MatchingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) Regex() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) RegexInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference_Override(d DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetInternalValue(val *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetMatchingType(val *string) {
	if err := j.validateSetMatchingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingType",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) PutDictionary(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleDictionary) {
	if err := d.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDictionary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeByHotword(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword) {
	if err := d.validatePutExcludeByHotwordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludeByHotword",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) PutExcludeInfoTypes(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) {
	if err := d.validatePutExcludeInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludeInfoTypes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) PutRegex(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleRegex) {
	if err := d.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRegex",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		d,
		"resetDictionary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeByHotword() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeByHotword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetExcludeInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

