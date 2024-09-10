// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglemonitoringistiocanonicalservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/datagooglemonitoringistiocanonicalservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/data-sources/monitoring_istio_canonical_service google_monitoring_istio_canonical_service}.
type DataGoogleMonitoringIstioCanonicalService interface {
	cdktf.TerraformDataSource
	CanonicalService() *string
	SetCanonicalService(val *string)
	CanonicalServiceInput() *string
	CanonicalServiceNamespace() *string
	SetCanonicalServiceNamespace(val *string)
	CanonicalServiceNamespaceInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MeshUid() *string
	SetMeshUid(val *string)
	MeshUidInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ServiceId() *string
	Telemetry() DataGoogleMonitoringIstioCanonicalServiceTelemetryList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserLabels() cdktf.StringMap
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGoogleMonitoringIstioCanonicalService
type jsiiProxy_DataGoogleMonitoringIstioCanonicalService struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) CanonicalService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canonicalService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) CanonicalServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canonicalServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) CanonicalServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canonicalServiceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) CanonicalServiceNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canonicalServiceNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) MeshUid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) MeshUidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) Telemetry() DataGoogleMonitoringIstioCanonicalServiceTelemetryList {
	var returns DataGoogleMonitoringIstioCanonicalServiceTelemetryList
	_jsii_.Get(
		j,
		"telemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) UserLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/data-sources/monitoring_istio_canonical_service google_monitoring_istio_canonical_service} Data Source.
func NewDataGoogleMonitoringIstioCanonicalService(scope constructs.Construct, id *string, config *DataGoogleMonitoringIstioCanonicalServiceConfig) DataGoogleMonitoringIstioCanonicalService {
	_init_.Initialize()

	if err := validateNewDataGoogleMonitoringIstioCanonicalServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleMonitoringIstioCanonicalService{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleMonitoringIstioCanonicalService.DataGoogleMonitoringIstioCanonicalService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/data-sources/monitoring_istio_canonical_service google_monitoring_istio_canonical_service} Data Source.
func NewDataGoogleMonitoringIstioCanonicalService_Override(d DataGoogleMonitoringIstioCanonicalService, scope constructs.Construct, id *string, config *DataGoogleMonitoringIstioCanonicalServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleMonitoringIstioCanonicalService.DataGoogleMonitoringIstioCanonicalService",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetCanonicalService(val *string) {
	if err := j.validateSetCanonicalServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canonicalService",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetCanonicalServiceNamespace(val *string) {
	if err := j.validateSetCanonicalServiceNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canonicalServiceNamespace",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetMeshUid(val *string) {
	if err := j.validateSetMeshUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meshUid",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGoogleMonitoringIstioCanonicalService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataGoogleMonitoringIstioCanonicalService resource upon running "cdktf plan <stack-name>".
func DataGoogleMonitoringIstioCanonicalService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGoogleMonitoringIstioCanonicalService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleMonitoringIstioCanonicalService.DataGoogleMonitoringIstioCanonicalService",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataGoogleMonitoringIstioCanonicalService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleMonitoringIstioCanonicalService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleMonitoringIstioCanonicalService.DataGoogleMonitoringIstioCanonicalService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleMonitoringIstioCanonicalService_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleMonitoringIstioCanonicalService_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleMonitoringIstioCanonicalService.DataGoogleMonitoringIstioCanonicalService",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleMonitoringIstioCanonicalService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleMonitoringIstioCanonicalService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleMonitoringIstioCanonicalService.DataGoogleMonitoringIstioCanonicalService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleMonitoringIstioCanonicalService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataGoogleMonitoringIstioCanonicalService.DataGoogleMonitoringIstioCanonicalService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleMonitoringIstioCanonicalService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

