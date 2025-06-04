// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/dataplexdatascan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference interface {
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
	InternalValue() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReport
	SetInternalValue(val *DataplexDatascanDataQualitySpecPostScanActionsNotificationReport)
	JobEndTrigger() DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTriggerOutputReference
	JobEndTriggerInput() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger
	JobFailureTrigger() DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTriggerOutputReference
	JobFailureTriggerInput() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger
	Recipients() DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipientsOutputReference
	RecipientsInput() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients
	ScoreThresholdTrigger() DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTriggerOutputReference
	ScoreThresholdTriggerInput() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger
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
	PutJobEndTrigger(value *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger)
	PutJobFailureTrigger(value *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger)
	PutRecipients(value *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients)
	PutScoreThresholdTrigger(value *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger)
	ResetJobEndTrigger()
	ResetJobFailureTrigger()
	ResetScoreThresholdTrigger()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference
type jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) InternalValue() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReport {
	var returns *DataplexDatascanDataQualitySpecPostScanActionsNotificationReport
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) JobEndTrigger() DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTriggerOutputReference {
	var returns DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTriggerOutputReference
	_jsii_.Get(
		j,
		"jobEndTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) JobEndTriggerInput() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger {
	var returns *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger
	_jsii_.Get(
		j,
		"jobEndTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) JobFailureTrigger() DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTriggerOutputReference {
	var returns DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTriggerOutputReference
	_jsii_.Get(
		j,
		"jobFailureTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) JobFailureTriggerInput() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger {
	var returns *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger
	_jsii_.Get(
		j,
		"jobFailureTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) Recipients() DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipientsOutputReference {
	var returns DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipientsOutputReference
	_jsii_.Get(
		j,
		"recipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) RecipientsInput() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients {
	var returns *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients
	_jsii_.Get(
		j,
		"recipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ScoreThresholdTrigger() DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTriggerOutputReference {
	var returns DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTriggerOutputReference
	_jsii_.Get(
		j,
		"scoreThresholdTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ScoreThresholdTriggerInput() *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger {
	var returns *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger
	_jsii_.Get(
		j,
		"scoreThresholdTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference_Override(d DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetInternalValue(val *DataplexDatascanDataQualitySpecPostScanActionsNotificationReport) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) PutJobEndTrigger(value *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger) {
	if err := d.validatePutJobEndTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJobEndTrigger",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) PutJobFailureTrigger(value *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger) {
	if err := d.validatePutJobFailureTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJobFailureTrigger",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) PutRecipients(value *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients) {
	if err := d.validatePutRecipientsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRecipients",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) PutScoreThresholdTrigger(value *DataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger) {
	if err := d.validatePutScoreThresholdTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScoreThresholdTrigger",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ResetJobEndTrigger() {
	_jsii_.InvokeVoid(
		d,
		"resetJobEndTrigger",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ResetJobFailureTrigger() {
	_jsii_.InvokeVoid(
		d,
		"resetJobFailureTrigger",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ResetScoreThresholdTrigger() {
	_jsii_.InvokeVoid(
		d,
		"resetScoreThresholdTrigger",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

