package accesscontextmanageraccesslevels

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/accesscontextmanageraccesslevels/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference interface {
	cdktf.ComplexObject
	AllowedDeviceManagementLevels() *[]*string
	SetAllowedDeviceManagementLevels(val *[]*string)
	AllowedDeviceManagementLevelsInput() *[]*string
	AllowedEncryptionStatuses() *[]*string
	SetAllowedEncryptionStatuses(val *[]*string)
	AllowedEncryptionStatusesInput() *[]*string
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
	InternalValue() *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy
	SetInternalValue(val *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy)
	OsConstraints() AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOsConstraintsList
	OsConstraintsInput() interface{}
	RequireAdminApproval() interface{}
	SetRequireAdminApproval(val interface{})
	RequireAdminApprovalInput() interface{}
	RequireCorpOwned() interface{}
	SetRequireCorpOwned(val interface{})
	RequireCorpOwnedInput() interface{}
	RequireScreenLock() interface{}
	SetRequireScreenLock(val interface{})
	RequireScreenLockInput() interface{}
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
	PutOsConstraints(value interface{})
	ResetAllowedDeviceManagementLevels()
	ResetAllowedEncryptionStatuses()
	ResetOsConstraints()
	ResetRequireAdminApproval()
	ResetRequireCorpOwned()
	ResetRequireScreenLock()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference
type jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) AllowedDeviceManagementLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDeviceManagementLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) AllowedDeviceManagementLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDeviceManagementLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) AllowedEncryptionStatuses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEncryptionStatuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) AllowedEncryptionStatusesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEncryptionStatusesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) InternalValue() *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy {
	var returns *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) OsConstraints() AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOsConstraintsList {
	var returns AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOsConstraintsList
	_jsii_.Get(
		j,
		"osConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) OsConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"osConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) RequireAdminApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAdminApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) RequireAdminApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAdminApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) RequireCorpOwned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCorpOwned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) RequireCorpOwnedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCorpOwnedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) RequireScreenLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireScreenLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) RequireScreenLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireScreenLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerAccessLevels.AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference_Override(a AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerAccessLevels.AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetAllowedDeviceManagementLevels(val *[]*string) {
	if err := j.validateSetAllowedDeviceManagementLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDeviceManagementLevels",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetAllowedEncryptionStatuses(val *[]*string) {
	if err := j.validateSetAllowedEncryptionStatusesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedEncryptionStatuses",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetInternalValue(val *AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetRequireAdminApproval(val interface{}) {
	if err := j.validateSetRequireAdminApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAdminApproval",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetRequireCorpOwned(val interface{}) {
	if err := j.validateSetRequireCorpOwnedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCorpOwned",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetRequireScreenLock(val interface{}) {
	if err := j.validateSetRequireScreenLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireScreenLock",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) PutOsConstraints(value interface{}) {
	if err := a.validatePutOsConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOsConstraints",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ResetAllowedDeviceManagementLevels() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedDeviceManagementLevels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ResetAllowedEncryptionStatuses() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedEncryptionStatuses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ResetOsConstraints() {
	_jsii_.InvokeVoid(
		a,
		"resetOsConstraints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ResetRequireAdminApproval() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireAdminApproval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ResetRequireCorpOwned() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireCorpOwned",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ResetRequireScreenLock() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireScreenLock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

