// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecomputeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagooglecomputeregionbackendservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference interface {
	cdktf.ComplexObject
	BaseEjectionTime() DataGoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeList
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
	ConsecutiveErrors() *float64
	ConsecutiveGatewayFailure() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnforcingConsecutiveErrors() *float64
	EnforcingConsecutiveGatewayFailure() *float64
	EnforcingSuccessRate() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleComputeRegionBackendServiceOutlierDetection
	SetInternalValue(val *DataGoogleComputeRegionBackendServiceOutlierDetection)
	Interval() DataGoogleComputeRegionBackendServiceOutlierDetectionIntervalList
	MaxEjectionPercent() *float64
	SuccessRateMinimumHosts() *float64
	SuccessRateRequestVolume() *float64
	SuccessRateStdevFactor() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference
type jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) BaseEjectionTime() DataGoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeList {
	var returns DataGoogleComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeList
	_jsii_.Get(
		j,
		"baseEjectionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveErrors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveGatewayFailure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveGatewayFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveErrors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveGatewayFailure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveGatewayFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingSuccessRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingSuccessRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) InternalValue() *DataGoogleComputeRegionBackendServiceOutlierDetection {
	var returns *DataGoogleComputeRegionBackendServiceOutlierDetection
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) Interval() DataGoogleComputeRegionBackendServiceOutlierDetectionIntervalList {
	var returns DataGoogleComputeRegionBackendServiceOutlierDetectionIntervalList
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) MaxEjectionPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateMinimumHosts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateMinimumHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateRequestVolume() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateRequestVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateStdevFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateStdevFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeRegionBackendServiceOutlierDetectionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeRegionBackendService.DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference_Override(d DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeRegionBackendService.DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetInternalValue(val *DataGoogleComputeRegionBackendServiceOutlierDetection) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleComputeRegionBackendServiceOutlierDetectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

