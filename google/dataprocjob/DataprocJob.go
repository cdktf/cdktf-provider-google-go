// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/dataprocjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/dataproc_job google_dataproc_job}.
type DataprocJob interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	DriverControlsFilesUri() *string
	DriverOutputResourceUri() *string
	EffectiveLabels() cdktf.StringMap
	ForceDelete() interface{}
	SetForceDelete(val interface{})
	ForceDeleteInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HadoopConfig() DataprocJobHadoopConfigOutputReference
	HadoopConfigInput() *DataprocJobHadoopConfig
	HiveConfig() DataprocJobHiveConfigOutputReference
	HiveConfigInput() *DataprocJobHiveConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PigConfig() DataprocJobPigConfigOutputReference
	PigConfigInput() *DataprocJobPigConfig
	Placement() DataprocJobPlacementOutputReference
	PlacementInput() *DataprocJobPlacement
	PrestoConfig() DataprocJobPrestoConfigOutputReference
	PrestoConfigInput() *DataprocJobPrestoConfig
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PysparkConfig() DataprocJobPysparkConfigOutputReference
	PysparkConfigInput() *DataprocJobPysparkConfig
	// Experimental.
	RawOverrides() interface{}
	Reference() DataprocJobReferenceOutputReference
	ReferenceInput() *DataprocJobReference
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Scheduling() DataprocJobSchedulingOutputReference
	SchedulingInput() *DataprocJobScheduling
	SparkConfig() DataprocJobSparkConfigOutputReference
	SparkConfigInput() *DataprocJobSparkConfig
	SparksqlConfig() DataprocJobSparksqlConfigOutputReference
	SparksqlConfigInput() *DataprocJobSparksqlConfig
	Status() DataprocJobStatusList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataprocJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutHadoopConfig(value *DataprocJobHadoopConfig)
	PutHiveConfig(value *DataprocJobHiveConfig)
	PutPigConfig(value *DataprocJobPigConfig)
	PutPlacement(value *DataprocJobPlacement)
	PutPrestoConfig(value *DataprocJobPrestoConfig)
	PutPysparkConfig(value *DataprocJobPysparkConfig)
	PutReference(value *DataprocJobReference)
	PutScheduling(value *DataprocJobScheduling)
	PutSparkConfig(value *DataprocJobSparkConfig)
	PutSparksqlConfig(value *DataprocJobSparksqlConfig)
	PutTimeouts(value *DataprocJobTimeouts)
	ResetForceDelete()
	ResetHadoopConfig()
	ResetHiveConfig()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPigConfig()
	ResetPrestoConfig()
	ResetProject()
	ResetPysparkConfig()
	ResetReference()
	ResetRegion()
	ResetScheduling()
	ResetSparkConfig()
	ResetSparksqlConfig()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for DataprocJob
type jsiiProxy_DataprocJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataprocJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) DriverControlsFilesUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverControlsFilesUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) DriverOutputResourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverOutputResourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) HadoopConfig() DataprocJobHadoopConfigOutputReference {
	var returns DataprocJobHadoopConfigOutputReference
	_jsii_.Get(
		j,
		"hadoopConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) HadoopConfigInput() *DataprocJobHadoopConfig {
	var returns *DataprocJobHadoopConfig
	_jsii_.Get(
		j,
		"hadoopConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) HiveConfig() DataprocJobHiveConfigOutputReference {
	var returns DataprocJobHiveConfigOutputReference
	_jsii_.Get(
		j,
		"hiveConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) HiveConfigInput() *DataprocJobHiveConfig {
	var returns *DataprocJobHiveConfig
	_jsii_.Get(
		j,
		"hiveConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) PigConfig() DataprocJobPigConfigOutputReference {
	var returns DataprocJobPigConfigOutputReference
	_jsii_.Get(
		j,
		"pigConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) PigConfigInput() *DataprocJobPigConfig {
	var returns *DataprocJobPigConfig
	_jsii_.Get(
		j,
		"pigConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Placement() DataprocJobPlacementOutputReference {
	var returns DataprocJobPlacementOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) PlacementInput() *DataprocJobPlacement {
	var returns *DataprocJobPlacement
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) PrestoConfig() DataprocJobPrestoConfigOutputReference {
	var returns DataprocJobPrestoConfigOutputReference
	_jsii_.Get(
		j,
		"prestoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) PrestoConfigInput() *DataprocJobPrestoConfig {
	var returns *DataprocJobPrestoConfig
	_jsii_.Get(
		j,
		"prestoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) PysparkConfig() DataprocJobPysparkConfigOutputReference {
	var returns DataprocJobPysparkConfigOutputReference
	_jsii_.Get(
		j,
		"pysparkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) PysparkConfigInput() *DataprocJobPysparkConfig {
	var returns *DataprocJobPysparkConfig
	_jsii_.Get(
		j,
		"pysparkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Reference() DataprocJobReferenceOutputReference {
	var returns DataprocJobReferenceOutputReference
	_jsii_.Get(
		j,
		"reference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) ReferenceInput() *DataprocJobReference {
	var returns *DataprocJobReference
	_jsii_.Get(
		j,
		"referenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Scheduling() DataprocJobSchedulingOutputReference {
	var returns DataprocJobSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) SchedulingInput() *DataprocJobScheduling {
	var returns *DataprocJobScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) SparkConfig() DataprocJobSparkConfigOutputReference {
	var returns DataprocJobSparkConfigOutputReference
	_jsii_.Get(
		j,
		"sparkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) SparkConfigInput() *DataprocJobSparkConfig {
	var returns *DataprocJobSparkConfig
	_jsii_.Get(
		j,
		"sparkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) SparksqlConfig() DataprocJobSparksqlConfigOutputReference {
	var returns DataprocJobSparksqlConfigOutputReference
	_jsii_.Get(
		j,
		"sparksqlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) SparksqlConfigInput() *DataprocJobSparksqlConfig {
	var returns *DataprocJobSparksqlConfig
	_jsii_.Get(
		j,
		"sparksqlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Status() DataprocJobStatusList {
	var returns DataprocJobStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) Timeouts() DataprocJobTimeoutsOutputReference {
	var returns DataprocJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/dataproc_job google_dataproc_job} Resource.
func NewDataprocJob(scope constructs.Construct, id *string, config *DataprocJobConfig) DataprocJob {
	_init_.Initialize()

	if err := validateNewDataprocJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocJob{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocJob.DataprocJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/dataproc_job google_dataproc_job} Resource.
func NewDataprocJob_Override(d DataprocJob, scope constructs.Construct, id *string, config *DataprocJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocJob.DataprocJob",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataprocJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataprocJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a DataprocJob resource upon running "cdktf plan <stack-name>".
func DataprocJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataprocJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataprocJob.DataprocJob",
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
func DataprocJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataprocJob.DataprocJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataprocJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataprocJob.DataprocJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataprocJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataprocJob.DataprocJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataprocJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataprocJob.DataprocJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataprocJob) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataprocJob) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataprocJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataprocJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocJob) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataprocJob) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataprocJob) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataprocJob) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataprocJob) PutHadoopConfig(value *DataprocJobHadoopConfig) {
	if err := d.validatePutHadoopConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHadoopConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutHiveConfig(value *DataprocJobHiveConfig) {
	if err := d.validatePutHiveConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHiveConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutPigConfig(value *DataprocJobPigConfig) {
	if err := d.validatePutPigConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPigConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutPlacement(value *DataprocJobPlacement) {
	if err := d.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPlacement",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutPrestoConfig(value *DataprocJobPrestoConfig) {
	if err := d.validatePutPrestoConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrestoConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutPysparkConfig(value *DataprocJobPysparkConfig) {
	if err := d.validatePutPysparkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPysparkConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutReference(value *DataprocJobReference) {
	if err := d.validatePutReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReference",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutScheduling(value *DataprocJobScheduling) {
	if err := d.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScheduling",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutSparkConfig(value *DataprocJobSparkConfig) {
	if err := d.validatePutSparkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutSparksqlConfig(value *DataprocJobSparksqlConfig) {
	if err := d.validatePutSparksqlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparksqlConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) PutTimeouts(value *DataprocJobTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJob) ResetForceDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetHadoopConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetHadoopConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetHiveConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetHiveConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetPigConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPigConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetPrestoConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPrestoConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetPysparkConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPysparkConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetReference() {
	_jsii_.InvokeVoid(
		d,
		"resetReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetScheduling() {
	_jsii_.InvokeVoid(
		d,
		"resetScheduling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetSparkConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetSparksqlConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSparksqlConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

