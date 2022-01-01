package awssolutionsconstructscore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.AddProxyMethodToApiResourceInputParams",
		reflect.TypeOf((*AddProxyMethodToApiResourceInputParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildDeadLetterQueueProps",
		reflect.TypeOf((*BuildDeadLetterQueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildDynamoDBTableProps",
		reflect.TypeOf((*BuildDynamoDBTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildDynamoDBTableWithStreamProps",
		reflect.TypeOf((*BuildDynamoDBTableWithStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildEventBusProps",
		reflect.TypeOf((*BuildEventBusProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildGlueJobProps",
		reflect.TypeOf((*BuildGlueJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildKinesisAnalyticsAppProps",
		reflect.TypeOf((*BuildKinesisAnalyticsAppProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildKinesisStreamProps",
		reflect.TypeOf((*BuildKinesisStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildLambdaFunctionProps",
		reflect.TypeOf((*BuildLambdaFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildQueueProps",
		reflect.TypeOf((*BuildQueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildS3BucketProps",
		reflect.TypeOf((*BuildS3BucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildSagemakerEndpointProps",
		reflect.TypeOf((*BuildSagemakerEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildSagemakerNotebookProps",
		reflect.TypeOf((*BuildSagemakerNotebookProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildTopicProps",
		reflect.TypeOf((*BuildTopicProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildVpcProps",
		reflect.TypeOf((*BuildVpcProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.BuildWebaclProps",
		reflect.TypeOf((*BuildWebaclProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.CfnDomainOptions",
		reflect.TypeOf((*CfnDomainOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.CfnNagSuppressRule",
		reflect.TypeOf((*CfnNagSuppressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.CognitoOptions",
		reflect.TypeOf((*CognitoOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.EventSourceProps",
		reflect.TypeOf((*EventSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.SecurityGroupRuleDefinition",
		reflect.TypeOf((*SecurityGroupRuleDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-solutions-constructs/core.ServiceEndpointTypes",
		reflect.TypeOf((*ServiceEndpointTypes)(nil)).Elem(),
		map[string]interface{}{
			"DYNAMODB": ServiceEndpointTypes_DYNAMODB,
			"SNS": ServiceEndpointTypes_SNS,
			"SQS": ServiceEndpointTypes_SQS,
			"S3": ServiceEndpointTypes_S3,
			"STEP_FUNCTIONS": ServiceEndpointTypes_STEP_FUNCTIONS,
			"SAGEMAKER_RUNTIME": ServiceEndpointTypes_SAGEMAKER_RUNTIME,
			"SECRETS_MANAGER": ServiceEndpointTypes_SECRETS_MANAGER,
			"SSM": ServiceEndpointTypes_SSM,
			"EVENTS": ServiceEndpointTypes_EVENTS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.SinkDataStoreProps",
		reflect.TypeOf((*SinkDataStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-solutions-constructs/core.SinkStoreType",
		reflect.TypeOf((*SinkStoreType)(nil)).Elem(),
		map[string]interface{}{
			"S3": SinkStoreType_S3,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/core.VerifiedProps",
		reflect.TypeOf((*VerifiedProps)(nil)).Elem(),
	)
}
