// Store AWS SSM Parameter Store Parameters into another AWS Region with AWS CDK
package almacdkcrossregionparameter

import (
	_init_ "github.com/alma-cdk/cross-region-parameter-go/almacdkcrossregionparameter/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/alma-cdk/cross-region-parameter-go/almacdkcrossregionparameter/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Cross-Region SSM Parameter.
// Experimental.
type CrossRegionParameter interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CrossRegionParameter
type jsiiProxy_CrossRegionParameter struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CrossRegionParameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Define a new Cross-Region SSM Parameter.
//
// Example:
//   new CrossRegionParameter(this, 'SayHiToSweden', {
//     region: 'eu-north-1',
//     name: '/parameter/path/message',
//     description: 'Some message for the Swedes',
//     value: 'Hej då!',
//   });
//
// Experimental.
func NewCrossRegionParameter(scope constructs.Construct, name *string, props *CrossRegionParameterProps) CrossRegionParameter {
	_init_.Initialize()

	if err := validateNewCrossRegionParameterParameters(scope, name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CrossRegionParameter{}

	_jsii_.Create(
		"@alma-cdk/cross-region-parameter.CrossRegionParameter",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

// Define a new Cross-Region SSM Parameter.
//
// Example:
//   new CrossRegionParameter(this, 'SayHiToSweden', {
//     region: 'eu-north-1',
//     name: '/parameter/path/message',
//     description: 'Some message for the Swedes',
//     value: 'Hej då!',
//   });
//
// Experimental.
func NewCrossRegionParameter_Override(c CrossRegionParameter, scope constructs.Construct, name *string, props *CrossRegionParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@alma-cdk/cross-region-parameter.CrossRegionParameter",
		[]interface{}{scope, name, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CrossRegionParameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCrossRegionParameter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@alma-cdk/cross-region-parameter.CrossRegionParameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossRegionParameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

