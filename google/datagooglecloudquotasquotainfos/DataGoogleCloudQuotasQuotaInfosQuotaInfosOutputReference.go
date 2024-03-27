// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecloudquotasquotainfos

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/datagooglecloudquotasquotainfos/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference interface {
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
	ContainerType() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimensions() *[]*string
	DimensionsInfos() DataGoogleCloudQuotasQuotaInfosQuotaInfosDimensionsInfosList
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleCloudQuotasQuotaInfosQuotaInfos
	SetInternalValue(val *DataGoogleCloudQuotasQuotaInfosQuotaInfos)
	IsConcurrent() cdktf.IResolvable
	IsFixed() cdktf.IResolvable
	IsPrecise() cdktf.IResolvable
	Metric() *string
	MetricDisplayName() *string
	MetricUnit() *string
	Name() *string
	QuotaDisplayName() *string
	QuotaId() *string
	QuotaIncreaseEligibility() DataGoogleCloudQuotasQuotaInfosQuotaInfosQuotaIncreaseEligibilityList
	RefreshInterval() *string
	Service() *string
	ServiceRequestQuotaUri() *string
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

// The jsii proxy struct for DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference
type jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) ContainerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) Dimensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) DimensionsInfos() DataGoogleCloudQuotasQuotaInfosQuotaInfosDimensionsInfosList {
	var returns DataGoogleCloudQuotasQuotaInfosQuotaInfosDimensionsInfosList
	_jsii_.Get(
		j,
		"dimensionsInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) InternalValue() *DataGoogleCloudQuotasQuotaInfosQuotaInfos {
	var returns *DataGoogleCloudQuotasQuotaInfosQuotaInfos
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) IsConcurrent() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isConcurrent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) IsFixed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isFixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) IsPrecise() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isPrecise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) MetricDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) MetricUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) QuotaDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) QuotaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) QuotaIncreaseEligibility() DataGoogleCloudQuotasQuotaInfosQuotaInfosQuotaIncreaseEligibilityList {
	var returns DataGoogleCloudQuotasQuotaInfosQuotaInfosQuotaIncreaseEligibilityList
	_jsii_.Get(
		j,
		"quotaIncreaseEligibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) RefreshInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) ServiceRequestQuotaUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRequestQuotaUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfos.DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference_Override(d DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleCloudQuotasQuotaInfos.DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference)SetInternalValue(val *DataGoogleCloudQuotasQuotaInfosQuotaInfos) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

