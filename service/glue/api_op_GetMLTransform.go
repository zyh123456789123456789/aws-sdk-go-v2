// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets an Glue machine learning transform artifact and all its corresponding
// metadata. Machine learning transforms are a special type of transform that use
// machine learning to learn the details of the transformation to be performed by
// learning from examples provided by humans. These transformations are then saved
// by Glue. You can retrieve their metadata by calling GetMLTransform .
func (c *Client) GetMLTransform(ctx context.Context, params *GetMLTransformInput, optFns ...func(*Options)) (*GetMLTransformOutput, error) {
	if params == nil {
		params = &GetMLTransformInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMLTransform", params, optFns, c.addOperationGetMLTransformMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMLTransformOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMLTransformInput struct {

	// The unique identifier of the transform, generated at the time that the
	// transform was created.
	//
	// This member is required.
	TransformId *string

	noSmithyDocumentSerde
}

type GetMLTransformOutput struct {

	// The date and time when the transform was created.
	CreatedOn *time.Time

	// A description of the transform.
	Description *string

	// The latest evaluation metrics.
	EvaluationMetrics *types.EvaluationMetrics

	// This value determines which version of Glue this machine learning transform is
	// compatible with. Glue 1.0 is recommended for most customers. If the value is not
	// set, the Glue compatibility defaults to Glue 0.9. For more information, see
	// Glue Versions (https://docs.aws.amazon.com/glue/latest/dg/release-notes.html#release-notes-versions)
	// in the developer guide.
	GlueVersion *string

	// A list of Glue table definitions used by the transform.
	InputRecordTables []types.GlueTable

	// The number of labels available for this transform.
	LabelCount int32

	// The date and time when the transform was last modified.
	LastModifiedOn *time.Time

	// The number of Glue data processing units (DPUs) that are allocated to task runs
	// for this transform. You can allocate from 2 to 100 DPUs; the default is 10. A
	// DPU is a relative measure of processing power that consists of 4 vCPUs of
	// compute capacity and 16 GB of memory. For more information, see the Glue
	// pricing page (https://aws.amazon.com/glue/pricing/) . When the WorkerType field
	// is set to a value other than Standard , the MaxCapacity field is set
	// automatically and becomes read-only.
	MaxCapacity *float64

	// The maximum number of times to retry a task for this transform after a task run
	// fails.
	MaxRetries *int32

	// The unique name given to the transform when it was created.
	Name *string

	// The number of workers of a defined workerType that are allocated when this task
	// runs.
	NumberOfWorkers *int32

	// The configuration parameters that are specific to the algorithm used.
	Parameters *types.TransformParameters

	// The name or Amazon Resource Name (ARN) of the IAM role with the required
	// permissions.
	Role *string

	// The Map object that represents the schema that this transform accepts. Has an
	// upper bound of 100 columns.
	Schema []types.SchemaColumn

	// The last known status of the transform (to indicate whether it can be used or
	// not). One of "NOT_READY", "READY", or "DELETING".
	Status types.TransformStatusType

	// The timeout for a task run for this transform in minutes. This is the maximum
	// time that a task run for this transform can consume resources before it is
	// terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	Timeout *int32

	// The encryption-at-rest settings of the transform that apply to accessing user
	// data. Machine learning transforms can access user data encrypted in Amazon S3
	// using KMS.
	TransformEncryption *types.TransformEncryption

	// The unique identifier of the transform, generated at the time that the
	// transform was created.
	TransformId *string

	// The type of predefined worker that is allocated when this task runs. Accepts a
	// value of Standard, G.1X, or G.2X.
	//   - For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory
	//   and a 50GB disk, and 2 executors per worker.
	//   - For the G.1X worker type, each worker provides 4 vCPU, 16 GB of memory and a
	//   64GB disk, and 1 executor per worker.
	//   - For the G.2X worker type, each worker provides 8 vCPU, 32 GB of memory and a
	//   128GB disk, and 1 executor per worker.
	WorkerType types.WorkerType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMLTransformMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMLTransform{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMLTransform{}, middleware.After)
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
	if err = addOpGetMLTransformValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMLTransform(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMLTransform(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetMLTransform",
	}
}
