// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Autopilot job. Find the best-performing model after you run an
// Autopilot job by calling DescribeAutoMLJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJob.html)
// . For information about how to use Autopilot, see Automate Model Development
// with Amazon SageMaker Autopilot (https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html)
// .
func (c *Client) CreateAutoMLJob(ctx context.Context, params *CreateAutoMLJobInput, optFns ...func(*Options)) (*CreateAutoMLJobOutput, error) {
	if params == nil {
		params = &CreateAutoMLJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAutoMLJob", params, optFns, c.addOperationCreateAutoMLJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAutoMLJobInput struct {

	// Identifies an Autopilot job. The name must be unique to your account and is
	// case insensitive.
	//
	// This member is required.
	AutoMLJobName *string

	// An array of channel objects that describes the input data and its location.
	// Each channel is a named input source. Similar to InputDataConfig supported by
	// HyperParameterTrainingJobDefinition (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html)
	// . Format(s) supported: CSV, Parquet. A minimum of 500 rows is required for the
	// training dataset. There is not a minimum number of rows required for the
	// validation dataset.
	//
	// This member is required.
	InputDataConfig []types.AutoMLChannel

	// Provides information about encryption and the Amazon S3 output path needed to
	// store artifacts from an AutoML job. Format(s) supported: CSV.
	//
	// This member is required.
	OutputDataConfig *types.AutoMLOutputDataConfig

	// The ARN of the role that is used to access the data.
	//
	// This member is required.
	RoleArn *string

	// A collection of settings used to configure an AutoML job.
	AutoMLJobConfig *types.AutoMLJobConfig

	// Defines the objective metric used to measure the predictive quality of an
	// AutoML job. You provide an AutoMLJobObjective$MetricName (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AutoMLJobObjective.html)
	// and Autopilot infers whether to minimize or maximize it. For CreateAutoMLJobV2 (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html)
	// , only Accuracy is supported.
	AutoMLJobObjective *types.AutoMLJobObjective

	// Generates possible candidates without training the models. A candidate is a
	// combination of data preprocessors, algorithms, and algorithm parameter settings.
	GenerateCandidateDefinitionsOnly bool

	// Specifies how to generate the endpoint name for an automatic one-click
	// Autopilot model deployment.
	ModelDeployConfig *types.ModelDeployConfig

	// Defines the type of supervised learning problem available for the candidates.
	// For more information, see Amazon SageMaker Autopilot problem types (https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-datasets-problem-types.html#autopilot-problem-types)
	// .
	ProblemType types.ProblemType

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web ServicesResources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// . Tag keys must be unique per resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAutoMLJobOutput struct {

	// The unique ARN assigned to the AutoML job when it is created.
	//
	// This member is required.
	AutoMLJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAutoMLJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAutoMLJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAutoMLJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAutoMLJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAutoMLJob(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateAutoMLJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateAutoMLJob",
	}
}
