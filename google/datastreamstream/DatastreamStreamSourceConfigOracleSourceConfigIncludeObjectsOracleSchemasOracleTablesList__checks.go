//go:build !no_runtime_type_checking

package datastreamstream

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTables:
		val := val.(*[]*DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTables)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTables:
		val_ := val.([]*DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTables)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTables; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

