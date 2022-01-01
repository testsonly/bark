// Core CDK Construct for patterns library
package awssolutionsconstructscore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediastore"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2"
)

type AddProxyMethodToApiResourceInputParams struct {
	ApiGatewayRole awsiam.IRole `json:"apiGatewayRole"`
	ApiMethod *string `json:"apiMethod"`
	ApiResource awsapigateway.IResource `json:"apiResource"`
	RequestTemplate *string `json:"requestTemplate"`
	Service *string `json:"service"`
	Action *string `json:"action"`
	AwsIntegrationProps interface{} `json:"awsIntegrationProps"`
	ContentType *string `json:"contentType"`
	MethodOptions *awsapigateway.MethodOptions `json:"methodOptions"`
	Path *string `json:"path"`
	RequestModel *map[string]awsapigateway.IModel `json:"requestModel"`
	RequestValidator awsapigateway.IRequestValidator `json:"requestValidator"`
}

type BuildDeadLetterQueueProps struct {
	// Optional user provided properties for the dead letter queue.
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	// Whether to deploy a secondary queue to be used as a dead letter queue.
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	// Existing instance of SQS queue object, providing both this and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead letter queue.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
}

type BuildDynamoDBTableProps struct {
	// Optional user provided props to override the default props.
	DynamoTableProps *awsdynamodb.TableProps `json:"dynamoTableProps"`
	// Existing instance of dynamodb table object.
	//
	// Providing both this and `dynamoTableProps` will cause an error.
	ExistingTableObj awsdynamodb.Table `json:"existingTableObj"`
}

type BuildDynamoDBTableWithStreamProps struct {
	// Optional user provided props to override the default props.
	DynamoTableProps *awsdynamodb.TableProps `json:"dynamoTableProps"`
	// Existing instance of dynamodb table object.
	//
	// Providing both this and `dynamoTableProps` will cause an error.
	ExistingTableInterface awsdynamodb.ITable `json:"existingTableInterface"`
}

type BuildEventBusProps struct {
	// Optional user provided props to override the default props for the SNS topic.
	EventBusProps *awsevents.EventBusProps `json:"eventBusProps"`
	// Existing instance of SNS Topic object, providing both this and `topicProps` will cause an error.
	ExistingEventBusInterface awsevents.IEventBus `json:"existingEventBusInterface"`
}

type BuildGlueJobProps struct {
	// AWS Glue database.
	Database awsglue.CfnDatabase `json:"database"`
	// AWS Glue table.
	Table awsglue.CfnTable `json:"table"`
	// Existing instance of the S3 bucket object, if this is set then the script location is ignored.
	ExistingCfnJob awsglue.CfnJob `json:"existingCfnJob"`
	// Glue ETL job properties.
	GlueJobProps interface{} `json:"glueJobProps"`
	// Output storage options.
	OutputDataStore *SinkDataStoreProps `json:"outputDataStore"`
}

type BuildKinesisAnalyticsAppProps struct {
	// A Kinesis Data Firehose for the Kinesis Streams application to connect to.
	KinesisFirehose awskinesisfirehose.CfnDeliveryStream `json:"kinesisFirehose"`
	// Optional user provided props to override the default props for the Kinesis analytics app.
	KinesisAnalyticsProps interface{} `json:"kinesisAnalyticsProps"`
}

type BuildKinesisStreamProps struct {
	// Existing instance of Kinesis Stream, providing both this and `kinesisStreamProps` will cause an error.
	ExistingStreamObj awskinesis.Stream `json:"existingStreamObj"`
	// Optional user provided props to override the default props for the Kinesis stream.
	KinesisStreamProps *awskinesis.StreamProps `json:"kinesisStreamProps"`
}

type BuildLambdaFunctionProps struct {
	// Existing instance of Lambda Function object, Providing both this and lambdaFunctionProps will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// A VPC where the Lambda function will access internal resources.
	Vpc awsec2.IVpc `json:"vpc"`
}

type BuildQueueProps struct {
	// Optional dead letter queue to pass bad requests to after the max receive count is reached.
	DeadLetterQueue *awssqs.DeadLetterQueue `json:"deadLetterQueue"`
	// Use a KMS Key, either managed by this CDK app, or imported.
	//
	// If importing an encryption key, it must be specified in
	// the encryptionKey property for this construct.
	EnableEncryptionWithCustomerManagedKey *bool `json:"enableEncryptionWithCustomerManagedKey"`
	// An optional, imported encryption key to encrypt the SQS queue with.
	EncryptionKey awskms.Key `json:"encryptionKey"`
	// Optional user-provided props to override the default props for the encryption key.
	EncryptionKeyProps *awskms.KeyProps `json:"encryptionKeyProps"`
	// Existing instance of SQS queue object, providing both this and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// Optional user provided props to override the default props for the primary queue.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
}

type BuildS3BucketProps struct {
	// User provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// User provided props to override the default props for the S3 Logging Bucket.
	LoggingBucketProps *awss3.BucketProps `json:"loggingBucketProps"`
	// Whether to turn on Access Logs for S3.
	//
	// Uses an S3 bucket with associated storage costs.
	// Enabling Access Logging is a best practice.
	LogS3AccessLogs *bool `json:"logS3AccessLogs"`
}

type BuildSagemakerEndpointProps struct {
	// User provided props to create Sagemaker Endpoint Configuration.
	EndpointConfigProps *awssagemaker.CfnEndpointConfigProps `json:"endpointConfigProps"`
	// User provided props to create Sagemaker Endpoint.
	EndpointProps *awssagemaker.CfnEndpointProps `json:"endpointProps"`
	// Existing Sagemaker Enpoint object, if this is set then the modelProps, endpointConfigProps, and endpointProps are ignored.
	ExistingSagemakerEndpointObj awssagemaker.CfnEndpoint `json:"existingSagemakerEndpointObj"`
	// User provided props to create Sagemaker Model.
	ModelProps interface{} `json:"modelProps"`
	// A VPC where the Sagemaker Endpoint will be placed.
	Vpc awsec2.IVpc `json:"vpc"`
}

type BuildSagemakerNotebookProps struct {
	// IAM Role Arn for Sagemaker NoteBookInstance.
	Role awsiam.Role `json:"role"`
	// Optional user provided props to deploy inside vpc.
	DeployInsideVpc *bool `json:"deployInsideVpc"`
	// An optional, Existing instance of notebook object.
	//
	// If this is set then the sagemakerNotebookProps is ignored
	ExistingNotebookObj awssagemaker.CfnNotebookInstance `json:"existingNotebookObj"`
	// Optional user provided props for CfnNotebookInstanceProps.
	SagemakerNotebookProps interface{} `json:"sagemakerNotebookProps"`
}

type BuildTopicProps struct {
	// Use a Customer Managed KMS Key, either managed by this CDK app, or imported.
	//
	// If importing an encryption key, it must be specified in
	// the encryptionKey property for this construct.
	EnableEncryptionWithCustomerManagedKey *bool `json:"enableEncryptionWithCustomerManagedKey"`
	// An optional, imported encryption key to encrypt the SNS topic with.
	EncryptionKey awskms.Key `json:"encryptionKey"`
	// Optional user-provided props to override the default props for the encryption key.
	EncryptionKeyProps *awskms.KeyProps `json:"encryptionKeyProps"`
	// Existing instance of SNS Topic object, providing both this and `topicProps` will cause an error.
	ExistingTopicObj awssns.Topic `json:"existingTopicObj"`
	// Optional user provided props to override the default props for the SNS topic.
	TopicProps *awssns.TopicProps `json:"topicProps"`
}

type BuildVpcProps struct {
	// One of the default VPC configurations available in vpc-defaults.
	DefaultVpcProps *awsec2.VpcProps `json:"defaultVpcProps"`
	// Construct specified props that override both the default props and user props for the VPC.
	ConstructVpcProps *awsec2.VpcProps `json:"constructVpcProps"`
	// Existing instance of a VPC, if this is set then the all Props are ignored.
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the VPC.
	UserVpcProps *awsec2.VpcProps `json:"userVpcProps"`
}

type BuildWebaclProps struct {
	// Existing instance of a WAF web ACL, if this is set then the all props are ignored.
	ExistingWebaclObj awswafv2.CfnWebACL `json:"existingWebaclObj"`
	// User provided props to override the default ACL props for WAF web ACL.
	WebaclProps *awswafv2.CfnWebACLProps `json:"webaclProps"`
}

type CfnDomainOptions struct {
	CognitoAuthorizedRoleARN *string `json:"cognitoAuthorizedRoleARN"`
	Identitypool awscognito.CfnIdentityPool `json:"identitypool"`
	Userpool awscognito.UserPool `json:"userpool"`
	ServiceRoleARN *string `json:"serviceRoleARN"`
}

// The CFN NAG suppress rule interface.
type CfnNagSuppressRule struct {
	Id *string `json:"id"`
	Reason *string `json:"reason"`
}

type CognitoOptions struct {
	Identitypool awscognito.CfnIdentityPool `json:"identitypool"`
	Userpool awscognito.UserPool `json:"userpool"`
	Userpoolclient awscognito.UserPoolClient `json:"userpoolclient"`
}

type EventSourceProps struct {
	DeploySqsDlqQueue *bool `json:"deploySqsDlqQueue"`
	EventSourceProps *awslambdaeventsources.StreamEventSourceProps `json:"eventSourceProps"`
	SqsDlqQueueProps *awssqs.QueueProps `json:"sqsDlqQueueProps"`
}

type SecurityGroupRuleDefinition struct {
	Connection awsec2.Port `json:"connection"`
	Peer awsec2.IPeer `json:"peer"`
	Description *string `json:"description"`
	RemoteRule *bool `json:"remoteRule"`
}

type ServiceEndpointTypes string

const (
	ServiceEndpointTypes_DYNAMODB ServiceEndpointTypes = "DYNAMODB"
	ServiceEndpointTypes_SNS ServiceEndpointTypes = "SNS"
	ServiceEndpointTypes_SQS ServiceEndpointTypes = "SQS"
	ServiceEndpointTypes_S3 ServiceEndpointTypes = "S3"
	ServiceEndpointTypes_STEP_FUNCTIONS ServiceEndpointTypes = "STEP_FUNCTIONS"
	ServiceEndpointTypes_SAGEMAKER_RUNTIME ServiceEndpointTypes = "SAGEMAKER_RUNTIME"
	ServiceEndpointTypes_SECRETS_MANAGER ServiceEndpointTypes = "SECRETS_MANAGER"
	ServiceEndpointTypes_SSM ServiceEndpointTypes = "SSM"
	ServiceEndpointTypes_EVENTS ServiceEndpointTypes = "EVENTS"
)

// Interface to define potential outputs to allow the construct define additional output destinations for ETL transformation.
type SinkDataStoreProps struct {
	// Sink data store type.
	DatastoreType SinkStoreType `json:"datastoreType"`
	// The output S3 location where the data should be written.
	//
	// The provided S3 bucket will be used to pass
	// the output location to the etl script as an argument to the AWS Glue job.
	//
	// If no location is provided, it will check if @outputBucketProps are provided. If not it will create a new
	// bucket if the @datastoreType is S3.
	//
	// The argument key is `output_path`. The value of the argument can be retrieve in the python script
	// as follows:
	//   getResolvedOptions(sys.argv, ["JOB_NAME", "output_path", <other arguments that are passed> ])
	//   output_path = args["output_path"]
	ExistingS3OutputBucket awss3.Bucket `json:"existingS3OutputBucket"`
	// If @existingS3OutputBUcket is provided, this parameter is ignored.
	//
	// If this parameter is not provided,
	// the construct will create a new bucket if the @datastoreType is S3.
	OutputBucketProps *awss3.BucketProps `json:"outputBucketProps"`
}

// Enumeration of data store types that could include S3, DynamoDB, DocumentDB, RDS or Redshift.
//
// Current
// construct implementation only supports S3, but potential to add other output types in the future
type SinkStoreType string

const (
	SinkStoreType_S3 SinkStoreType = "S3"
)

type VerifiedProps struct {
	AlbLoggingBucketProps *awss3.BucketProps `json:"albLoggingBucketProps"`
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	DeployVpc *bool `json:"deployVpc"`
	DynamoTableProps *awsdynamodb.TableProps `json:"dynamoTableProps"`
	EncryptionKey awskms.Key `json:"encryptionKey"`
	EncryptionKeyProps *awskms.KeyProps `json:"encryptionKeyProps"`
	EndpointProps *awssagemaker.CfnEndpointProps `json:"endpointProps"`
	ExistingBucketInterface awss3.IBucket `json:"existingBucketInterface"`
	ExistingBucketObj awss3.Bucket `json:"existingBucketObj"`
	ExistingGlueJob awsglue.CfnJob `json:"existingGlueJob"`
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	ExistingLoadBalancerObj awselasticloadbalancingv2.ApplicationLoadBalancer `json:"existingLoadBalancerObj"`
	ExistingLoggingBucketObj awss3.IBucket `json:"existingLoggingBucketObj"`
	ExistingMediaStoreContainerObj awsmediastore.CfnContainer `json:"existingMediaStoreContainerObj"`
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	ExistingSagemakerEndpointObj awssagemaker.CfnEndpoint `json:"existingSagemakerEndpointObj"`
	ExistingSecretObj awssecretsmanager.Secret `json:"existingSecretObj"`
	ExistingStreamObj awskinesis.Stream `json:"existingStreamObj"`
	ExistingTableInterface awsdynamodb.ITable `json:"existingTableInterface"`
	ExistingTableObj awsdynamodb.Table `json:"existingTableObj"`
	ExistingTopicObj awssns.Topic `json:"existingTopicObj"`
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	GlueJobProps *awsglue.CfnJobProps `json:"glueJobProps"`
	KinesisStreamProps *awskinesis.StreamProps `json:"kinesisStreamProps"`
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	LoadBalancerProps *awselasticloadbalancingv2.ApplicationLoadBalancerProps `json:"loadBalancerProps"`
	LogAlbAccessLogs *bool `json:"logAlbAccessLogs"`
	LoggingBucketProps *awss3.BucketProps `json:"loggingBucketProps"`
	LogS3AccessLogs *bool `json:"logS3AccessLogs"`
	MediaStoreContainerProps *awsmediastore.CfnContainerProps `json:"mediaStoreContainerProps"`
	QueueProps *awssqs.QueueProps `json:"queueProps"`
	SecretProps *awssecretsmanager.SecretProps `json:"secretProps"`
	TopicProps *awssns.TopicProps `json:"topicProps"`
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

