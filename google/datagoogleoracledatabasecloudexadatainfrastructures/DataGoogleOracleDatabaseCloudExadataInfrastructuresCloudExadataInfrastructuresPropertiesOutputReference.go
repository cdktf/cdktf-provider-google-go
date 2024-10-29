// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabasecloudexadatainfrastructures

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/datagoogleoracledatabasecloudexadatainfrastructures/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference interface {
	cdktf.ComplexObject
	ActivatedStorageCount() *float64
	AdditionalStorageCount() *float64
	AvailableStorageSizeGb() *float64
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
	ComputeCount() *float64
	CpuCount() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerContacts() DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesCustomerContactsList
	DataStorageSizeTb() *float64
	DbNodeStorageSizeGb() *float64
	DbServerVersion() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresProperties
	SetInternalValue(val *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresProperties)
	MaintenanceWindow() DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowList
	MaxCpuCount() *float64
	MaxDataStorageTb() *float64
	MaxDbNodeStorageSizeGb() *float64
	MaxMemoryGb() *float64
	MemorySizeGb() *float64
	MonthlyDbServerVersion() *string
	MonthlyStorageServerVersion() *string
	NextMaintenanceRunId() *string
	NextMaintenanceRunTime() *string
	NextSecurityMaintenanceRunTime() *string
	Ocid() *string
	OciUrl() *string
	Shape() *string
	State() *string
	StorageCount() *float64
	StorageServerVersion() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalStorageSizeGb() *float64
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

// The jsii proxy struct for DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference
type jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) ActivatedStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activatedStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) AdditionalStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) AvailableStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) CpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) CustomerContacts() DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesCustomerContactsList {
	var returns DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesCustomerContactsList
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) DbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) DbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) InternalValue() *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresProperties {
	var returns *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) MaintenanceWindow() DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowList {
	var returns DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesMaintenanceWindowList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) MaxCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) MaxDataStorageTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataStorageTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) MaxDbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) MaxMemoryGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) MonthlyDbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyDbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) MonthlyStorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyStorageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) NextMaintenanceRunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceRunId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) NextMaintenanceRunTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceRunTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) NextSecurityMaintenanceRunTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSecurityMaintenanceRunTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) StorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) StorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) TotalStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalStorageSizeGb",
		&returns,
	)
	return returns
}


func NewDataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructures.DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference_Override(d DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudExadataInfrastructures.DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference)SetInternalValue(val *DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudExadataInfrastructuresCloudExadataInfrastructuresPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

