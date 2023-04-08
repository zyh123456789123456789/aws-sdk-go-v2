// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an entire resource tree. This operation will delete the parent resource
// and its child resources. Child resources are resources that were created from
// another resource. For example, when a forecast is generated from a predictor,
// the forecast is the child resource and the predictor is the parent resource.
// Amazon Forecast resources possess the following parent-child resource
// hierarchies:
//   - Dataset: dataset import jobs
//   - Dataset Group: predictors, predictor backtest export jobs, forecasts,
//     forecast export jobs
//   - Predictor: predictor backtest export jobs, forecasts, forecast export jobs
//   - Forecast: forecast export jobs
//
// DeleteResourceTree will only delete Amazon Forecast resources, and will not
// delete datasets or exported files stored in Amazon S3.
func (c *Client) DeleteResourceTree(ctx context.Context, params *DeleteResourceTreeInput, optFns ...func(*Options)) (*DeleteResourceTreeOutput, error) {
	if params == nil {
		params = &DeleteResourceTreeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResourceTree", params, optFns, c.addOperationDeleteResourceTreeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResourceTreeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResourceTreeInput struct {

	// The Amazon Resource Name (ARN) of the parent resource to delete. All child
	// resources of the parent resource will also be deleted.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type DeleteResourceTreeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteResourceTreeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteResourceTree{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteResourceTree{}, middleware.After)
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
	if err = addOpDeleteResourceTreeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResourceTree(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteResourceTree(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DeleteResourceTree",
	}
}
