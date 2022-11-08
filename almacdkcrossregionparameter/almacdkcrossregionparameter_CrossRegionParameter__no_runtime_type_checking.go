//go:build no_runtime_type_checking

// Store AWS SSM Parameter Store Parameters into another AWS Region with AWS CDK
package almacdkcrossregionparameter

// Building without runtime type checking enabled, so all the below just return nil

func validateCrossRegionParameter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCrossRegionParameterParameters(scope constructs.Construct, name *string, props *CrossRegionParameterProps) error {
	return nil
}

