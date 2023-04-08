// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests a CloudFront function. To test a function, you provide an event object
// that represents an HTTP request or response that your CloudFront distribution
// could receive in production. CloudFront runs the function, passing it the event
// object that you provided, and returns the function's result (the modified event
// object) in the response. The response also contains function logs and error
// messages, if any exist. For more information about testing functions, see
// Testing functions (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/managing-functions.html#test-function)
// in the Amazon CloudFront Developer Guide. To test a function, you provide the
// function's name and version ( ETag value) along with the event object. To get
// the function's name and version, you can use ListFunctions and DescribeFunction .
func (c *Client) TestFunction(ctx context.Context, params *TestFunctionInput, optFns ...func(*Options)) (*TestFunctionOutput, error) {
	if params == nil {
		params = &TestFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestFunction", params, optFns, c.addOperationTestFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestFunctionInput struct {

	// The event object to test the function with. For more information about the
	// structure of the event object, see Testing functions (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/managing-functions.html#test-function)
	// in the Amazon CloudFront Developer Guide.
	//
	// This member is required.
	EventObject []byte

	// The current version ( ETag value) of the function that you are testing, which
	// you can get using DescribeFunction .
	//
	// This member is required.
	IfMatch *string

	// The name of the function that you are testing.
	//
	// This member is required.
	Name *string

	// The stage of the function that you are testing, either DEVELOPMENT or LIVE .
	Stage types.FunctionStage

	noSmithyDocumentSerde
}

type TestFunctionOutput struct {

	// An object that represents the result of running the function with the provided
	// event object.
	TestResult *types.TestResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpTestFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpTestFunction{}, middleware.After)
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
	if err = addOpTestFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestFunction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "TestFunction",
	}
}
