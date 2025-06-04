// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datalosspreventionjobtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference interface {
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
	DeidentifyTemplate() *string
	SetDeidentifyTemplate(val *string)
	DeidentifyTemplateInput() *string
	// Experimental.
	Fqn() *string
	ImageRedactTemplate() *string
	SetImageRedactTemplate(val *string)
	ImageRedactTemplateInput() *string
	InternalValue() *DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig
	SetInternalValue(val *DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig)
	StructuredDeidentifyTemplate() *string
	SetStructuredDeidentifyTemplate(val *string)
	StructuredDeidentifyTemplateInput() *string
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
	ResetDeidentifyTemplate()
	ResetImageRedactTemplate()
	ResetStructuredDeidentifyTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference
type jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) DeidentifyTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) DeidentifyTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ImageRedactTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRedactTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ImageRedactTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRedactTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) InternalValue() *DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig {
	var returns *DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) StructuredDeidentifyTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"structuredDeidentifyTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) StructuredDeidentifyTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"structuredDeidentifyTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference_Override(d DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference)SetDeidentifyTemplate(val *string) {
	if err := j.validateSetDeidentifyTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deidentifyTemplate",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference)SetImageRedactTemplate(val *string) {
	if err := j.validateSetImageRedactTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageRedactTemplate",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference)SetInternalValue(val *DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference)SetStructuredDeidentifyTemplate(val *string) {
	if err := j.validateSetStructuredDeidentifyTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"structuredDeidentifyTemplate",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ResetDeidentifyTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetDeidentifyTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ResetImageRedactTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetImageRedactTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ResetStructuredDeidentifyTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetStructuredDeidentifyTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

