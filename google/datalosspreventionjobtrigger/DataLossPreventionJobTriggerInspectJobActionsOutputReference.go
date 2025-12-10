// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datalosspreventionjobtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionJobTriggerInspectJobActionsOutputReference interface {
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
	Deidentify() DataLossPreventionJobTriggerInspectJobActionsDeidentifyOutputReference
	DeidentifyInput() *DataLossPreventionJobTriggerInspectJobActionsDeidentify
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JobNotificationEmails() DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmailsOutputReference
	JobNotificationEmailsInput() *DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails
	PublishFindingsToCloudDataCatalog() DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogOutputReference
	PublishFindingsToCloudDataCatalogInput() *DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	PublishSummaryToCscc() DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCsccOutputReference
	PublishSummaryToCsccInput() *DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc
	PublishToStackdriver() DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriverOutputReference
	PublishToStackdriverInput() *DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver
	PubSub() DataLossPreventionJobTriggerInspectJobActionsPubSubOutputReference
	PubSubInput() *DataLossPreventionJobTriggerInspectJobActionsPubSub
	SaveFindings() DataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputReference
	SaveFindingsInput() *DataLossPreventionJobTriggerInspectJobActionsSaveFindings
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDeidentify(value *DataLossPreventionJobTriggerInspectJobActionsDeidentify)
	PutJobNotificationEmails(value *DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails)
	PutPublishFindingsToCloudDataCatalog(value *DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog)
	PutPublishSummaryToCscc(value *DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc)
	PutPublishToStackdriver(value *DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver)
	PutPubSub(value *DataLossPreventionJobTriggerInspectJobActionsPubSub)
	PutSaveFindings(value *DataLossPreventionJobTriggerInspectJobActionsSaveFindings)
	ResetDeidentify()
	ResetJobNotificationEmails()
	ResetPublishFindingsToCloudDataCatalog()
	ResetPublishSummaryToCscc()
	ResetPublishToStackdriver()
	ResetPubSub()
	ResetSaveFindings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobActionsOutputReference
type jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) Deidentify() DataLossPreventionJobTriggerInspectJobActionsDeidentifyOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobActionsDeidentifyOutputReference
	_jsii_.Get(
		j,
		"deidentify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) DeidentifyInput() *DataLossPreventionJobTriggerInspectJobActionsDeidentify {
	var returns *DataLossPreventionJobTriggerInspectJobActionsDeidentify
	_jsii_.Get(
		j,
		"deidentifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) JobNotificationEmails() DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmailsOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmailsOutputReference
	_jsii_.Get(
		j,
		"jobNotificationEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) JobNotificationEmailsInput() *DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails {
	var returns *DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails
	_jsii_.Get(
		j,
		"jobNotificationEmailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishFindingsToCloudDataCatalog() DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalogOutputReference
	_jsii_.Get(
		j,
		"publishFindingsToCloudDataCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishFindingsToCloudDataCatalogInput() *DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog {
	var returns *DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog
	_jsii_.Get(
		j,
		"publishFindingsToCloudDataCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishSummaryToCscc() DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCsccOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCsccOutputReference
	_jsii_.Get(
		j,
		"publishSummaryToCscc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishSummaryToCsccInput() *DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc {
	var returns *DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc
	_jsii_.Get(
		j,
		"publishSummaryToCsccInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishToStackdriver() DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriverOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriverOutputReference
	_jsii_.Get(
		j,
		"publishToStackdriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PublishToStackdriverInput() *DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver {
	var returns *DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver
	_jsii_.Get(
		j,
		"publishToStackdriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PubSub() DataLossPreventionJobTriggerInspectJobActionsPubSubOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobActionsPubSubOutputReference
	_jsii_.Get(
		j,
		"pubSub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PubSubInput() *DataLossPreventionJobTriggerInspectJobActionsPubSub {
	var returns *DataLossPreventionJobTriggerInspectJobActionsPubSub
	_jsii_.Get(
		j,
		"pubSubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) SaveFindings() DataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobActionsSaveFindingsOutputReference
	_jsii_.Get(
		j,
		"saveFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) SaveFindingsInput() *DataLossPreventionJobTriggerInspectJobActionsSaveFindings {
	var returns *DataLossPreventionJobTriggerInspectJobActionsSaveFindings
	_jsii_.Get(
		j,
		"saveFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionJobTriggerInspectJobActionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobActionsOutputReference_Override(d DataLossPreventionJobTriggerInspectJobActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PutDeidentify(value *DataLossPreventionJobTriggerInspectJobActionsDeidentify) {
	if err := d.validatePutDeidentifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeidentify",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PutJobNotificationEmails(value *DataLossPreventionJobTriggerInspectJobActionsJobNotificationEmails) {
	if err := d.validatePutJobNotificationEmailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJobNotificationEmails",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PutPublishFindingsToCloudDataCatalog(value *DataLossPreventionJobTriggerInspectJobActionsPublishFindingsToCloudDataCatalog) {
	if err := d.validatePutPublishFindingsToCloudDataCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPublishFindingsToCloudDataCatalog",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PutPublishSummaryToCscc(value *DataLossPreventionJobTriggerInspectJobActionsPublishSummaryToCscc) {
	if err := d.validatePutPublishSummaryToCsccParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPublishSummaryToCscc",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PutPublishToStackdriver(value *DataLossPreventionJobTriggerInspectJobActionsPublishToStackdriver) {
	if err := d.validatePutPublishToStackdriverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPublishToStackdriver",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PutPubSub(value *DataLossPreventionJobTriggerInspectJobActionsPubSub) {
	if err := d.validatePutPubSubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPubSub",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) PutSaveFindings(value *DataLossPreventionJobTriggerInspectJobActionsSaveFindings) {
	if err := d.validatePutSaveFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSaveFindings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetDeidentify() {
	_jsii_.InvokeVoid(
		d,
		"resetDeidentify",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetJobNotificationEmails() {
	_jsii_.InvokeVoid(
		d,
		"resetJobNotificationEmails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetPublishFindingsToCloudDataCatalog() {
	_jsii_.InvokeVoid(
		d,
		"resetPublishFindingsToCloudDataCatalog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetPublishSummaryToCscc() {
	_jsii_.InvokeVoid(
		d,
		"resetPublishSummaryToCscc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetPublishToStackdriver() {
	_jsii_.InvokeVoid(
		d,
		"resetPublishToStackdriver",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetPubSub() {
	_jsii_.InvokeVoid(
		d,
		"resetPubSub",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ResetSaveFindings() {
	_jsii_.InvokeVoid(
		d,
		"resetSaveFindings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

