// Store AWS SSM Parameter Store Parameters into another AWS Region with AWS CDK
package almacdkcrossregionparameter


// Tag properties.
// Experimental.
type TagProp struct {
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

