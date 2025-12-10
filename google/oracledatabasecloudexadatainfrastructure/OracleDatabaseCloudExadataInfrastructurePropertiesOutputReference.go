// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabasecloudexadatainfrastructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/oracledatabasecloudexadatainfrastructure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference interface {
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
	SetComputeCount(val *float64)
	ComputeCountInput() *float64
	CpuCount() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerContacts() OracleDatabaseCloudExadataInfrastructurePropertiesCustomerContactsList
	CustomerContactsInput() interface{}
	DataStorageSizeTb() *float64
	DbNodeStorageSizeGb() *float64
	DbServerVersion() *string
	// Experimental.
	Fqn() *string
	InternalValue() *OracleDatabaseCloudExadataInfrastructureProperties
	SetInternalValue(val *OracleDatabaseCloudExadataInfrastructureProperties)
	MaintenanceWindow() OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference
	MaintenanceWindowInput() *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow
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
	SetShape(val *string)
	ShapeInput() *string
	State() *string
	StorageCount() *float64
	SetStorageCount(val *float64)
	StorageCountInput() *float64
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
	SetTotalStorageSizeGb(val *float64)
	TotalStorageSizeGbInput() *float64
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
	PutCustomerContacts(value interface{})
	PutMaintenanceWindow(value *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow)
	ResetComputeCount()
	ResetCustomerContacts()
	ResetMaintenanceWindow()
	ResetStorageCount()
	ResetTotalStorageSizeGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference
type jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ActivatedStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activatedStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) AdditionalStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) AvailableStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) CpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) CustomerContacts() OracleDatabaseCloudExadataInfrastructurePropertiesCustomerContactsList {
	var returns OracleDatabaseCloudExadataInfrastructurePropertiesCustomerContactsList
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) CustomerContactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) DbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) DbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) InternalValue() *OracleDatabaseCloudExadataInfrastructureProperties {
	var returns *OracleDatabaseCloudExadataInfrastructureProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaintenanceWindow() OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference {
	var returns OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaintenanceWindowInput() *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow {
	var returns *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaxCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaxDataStorageTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataStorageTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaxDbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MaxMemoryGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MonthlyDbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyDbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) MonthlyStorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyStorageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) NextMaintenanceRunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceRunId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) NextMaintenanceRunTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceRunTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) NextSecurityMaintenanceRunTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSecurityMaintenanceRunTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) StorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) StorageCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) StorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) TotalStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) TotalStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalStorageSizeGbInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseCloudExadataInfrastructurePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseCloudExadataInfrastructure.OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseCloudExadataInfrastructurePropertiesOutputReference_Override(o OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseCloudExadataInfrastructure.OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetInternalValue(val *OracleDatabaseCloudExadataInfrastructureProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetStorageCount(val *float64) {
	if err := j.validateSetStorageCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference)SetTotalStorageSizeGb(val *float64) {
	if err := j.validateSetTotalStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalStorageSizeGb",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) PutCustomerContacts(value interface{}) {
	if err := o.validatePutCustomerContactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCustomerContacts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) PutMaintenanceWindow(value *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow) {
	if err := o.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetComputeCount() {
	_jsii_.InvokeVoid(
		o,
		"resetComputeCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetCustomerContacts() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomerContacts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		o,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetStorageCount() {
	_jsii_.InvokeVoid(
		o,
		"resetStorageCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ResetTotalStorageSizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetTotalStorageSizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

