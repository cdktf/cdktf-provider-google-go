// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datalosspreventiondiscoveryconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionDiscoveryConfigActionsOutputReference interface {
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
	ExportData() DataLossPreventionDiscoveryConfigActionsExportDataOutputReference
	ExportDataInput() *DataLossPreventionDiscoveryConfigActionsExportData
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PubSubNotification() DataLossPreventionDiscoveryConfigActionsPubSubNotificationOutputReference
	PubSubNotificationInput() *DataLossPreventionDiscoveryConfigActionsPubSubNotification
	TagResources() DataLossPreventionDiscoveryConfigActionsTagResourcesOutputReference
	TagResourcesInput() *DataLossPreventionDiscoveryConfigActionsTagResources
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
	PutExportData(value *DataLossPreventionDiscoveryConfigActionsExportData)
	PutPubSubNotification(value *DataLossPreventionDiscoveryConfigActionsPubSubNotification)
	PutTagResources(value *DataLossPreventionDiscoveryConfigActionsTagResources)
	ResetExportData()
	ResetPubSubNotification()
	ResetTagResources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDiscoveryConfigActionsOutputReference
type jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ExportData() DataLossPreventionDiscoveryConfigActionsExportDataOutputReference {
	var returns DataLossPreventionDiscoveryConfigActionsExportDataOutputReference
	_jsii_.Get(
		j,
		"exportData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ExportDataInput() *DataLossPreventionDiscoveryConfigActionsExportData {
	var returns *DataLossPreventionDiscoveryConfigActionsExportData
	_jsii_.Get(
		j,
		"exportDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) PubSubNotification() DataLossPreventionDiscoveryConfigActionsPubSubNotificationOutputReference {
	var returns DataLossPreventionDiscoveryConfigActionsPubSubNotificationOutputReference
	_jsii_.Get(
		j,
		"pubSubNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) PubSubNotificationInput() *DataLossPreventionDiscoveryConfigActionsPubSubNotification {
	var returns *DataLossPreventionDiscoveryConfigActionsPubSubNotification
	_jsii_.Get(
		j,
		"pubSubNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) TagResources() DataLossPreventionDiscoveryConfigActionsTagResourcesOutputReference {
	var returns DataLossPreventionDiscoveryConfigActionsTagResourcesOutputReference
	_jsii_.Get(
		j,
		"tagResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) TagResourcesInput() *DataLossPreventionDiscoveryConfigActionsTagResources {
	var returns *DataLossPreventionDiscoveryConfigActionsTagResources
	_jsii_.Get(
		j,
		"tagResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionDiscoveryConfigActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionDiscoveryConfigActionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDiscoveryConfigActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionDiscoveryConfigActionsOutputReference_Override(d DataLossPreventionDiscoveryConfigActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) PutExportData(value *DataLossPreventionDiscoveryConfigActionsExportData) {
	if err := d.validatePutExportDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExportData",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) PutPubSubNotification(value *DataLossPreventionDiscoveryConfigActionsPubSubNotification) {
	if err := d.validatePutPubSubNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPubSubNotification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) PutTagResources(value *DataLossPreventionDiscoveryConfigActionsTagResources) {
	if err := d.validatePutTagResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTagResources",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ResetExportData() {
	_jsii_.InvokeVoid(
		d,
		"resetExportData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ResetPubSubNotification() {
	_jsii_.InvokeVoid(
		d,
		"resetPubSubNotification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ResetTagResources() {
	_jsii_.InvokeVoid(
		d,
		"resetTagResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

