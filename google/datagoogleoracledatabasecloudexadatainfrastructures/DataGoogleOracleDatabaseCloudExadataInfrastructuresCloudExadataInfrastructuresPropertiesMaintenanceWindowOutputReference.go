// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabasecloudexadatainfrastructures

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagoogleoracledatabasecloudexadatainfrastructures/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference interface {
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
	CustomActionTimeoutMins() *float64
	DaysOfWeek() *[]*string
	// Experimental.
	Fqn() *string
	HoursOfDay() *[]*float64
	InternalValue() *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindow
	SetInternalValue(val *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindow)
	IsCustomActionTimeoutEnabled() cdktf.IResolvable
	LeadTimeWeek() *float64
	Months() *[]*string
	PatchingMode() *string
	Preference() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeksOfMonth() *[]*float64
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

// The jsii proxy struct for DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference
type jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) CustomActionTimeoutMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customActionTimeoutMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) HoursOfDay() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) InternalValue() *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindow {
	var returns *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) IsCustomActionTimeoutEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isCustomActionTimeoutEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) LeadTimeWeek() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) Months() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) PatchingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) Preference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) WeeksOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}


func NewDataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructures.DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference_Override(d DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructures.DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference)SetInternalValue(val *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

