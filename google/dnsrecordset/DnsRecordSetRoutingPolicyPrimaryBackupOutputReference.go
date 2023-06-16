package dnsrecordset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/dnsrecordset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsRecordSetRoutingPolicyPrimaryBackupOutputReference interface {
	cdktf.ComplexObject
	BackupGeo() DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoList
	BackupGeoInput() interface{}
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
	EnableGeoFencingForBackups() interface{}
	SetEnableGeoFencingForBackups(val interface{})
	EnableGeoFencingForBackupsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DnsRecordSetRoutingPolicyPrimaryBackup
	SetInternalValue(val *DnsRecordSetRoutingPolicyPrimaryBackup)
	Primary() DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference
	PrimaryInput() *DnsRecordSetRoutingPolicyPrimaryBackupPrimary
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrickleRatio() *float64
	SetTrickleRatio(val *float64)
	TrickleRatioInput() *float64
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
	PutBackupGeo(value interface{})
	PutPrimary(value *DnsRecordSetRoutingPolicyPrimaryBackupPrimary)
	ResetEnableGeoFencingForBackups()
	ResetTrickleRatio()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DnsRecordSetRoutingPolicyPrimaryBackupOutputReference
type jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) BackupGeo() DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoList {
	var returns DnsRecordSetRoutingPolicyPrimaryBackupBackupGeoList
	_jsii_.Get(
		j,
		"backupGeo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) BackupGeoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupGeoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) EnableGeoFencingForBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGeoFencingForBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) EnableGeoFencingForBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGeoFencingForBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) InternalValue() *DnsRecordSetRoutingPolicyPrimaryBackup {
	var returns *DnsRecordSetRoutingPolicyPrimaryBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) Primary() DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference {
	var returns DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) PrimaryInput() *DnsRecordSetRoutingPolicyPrimaryBackupPrimary {
	var returns *DnsRecordSetRoutingPolicyPrimaryBackupPrimary
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) TrickleRatio() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trickleRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) TrickleRatioInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trickleRatioInput",
		&returns,
	)
	return returns
}


func NewDnsRecordSetRoutingPolicyPrimaryBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DnsRecordSetRoutingPolicyPrimaryBackupOutputReference {
	_init_.Initialize()

	if err := validateNewDnsRecordSetRoutingPolicyPrimaryBackupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dnsRecordSet.DnsRecordSetRoutingPolicyPrimaryBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDnsRecordSetRoutingPolicyPrimaryBackupOutputReference_Override(d DnsRecordSetRoutingPolicyPrimaryBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dnsRecordSet.DnsRecordSetRoutingPolicyPrimaryBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference)SetEnableGeoFencingForBackups(val interface{}) {
	if err := j.validateSetEnableGeoFencingForBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGeoFencingForBackups",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference)SetInternalValue(val *DnsRecordSetRoutingPolicyPrimaryBackup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference)SetTrickleRatio(val *float64) {
	if err := j.validateSetTrickleRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trickleRatio",
		val,
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) PutBackupGeo(value interface{}) {
	if err := d.validatePutBackupGeoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBackupGeo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) PutPrimary(value *DnsRecordSetRoutingPolicyPrimaryBackupPrimary) {
	if err := d.validatePutPrimaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrimary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) ResetEnableGeoFencingForBackups() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableGeoFencingForBackups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) ResetTrickleRatio() {
	_jsii_.InvokeVoid(
		d,
		"resetTrickleRatio",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

