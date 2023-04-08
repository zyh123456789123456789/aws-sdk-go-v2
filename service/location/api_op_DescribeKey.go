// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the API key resource details. The API keys feature is in preview. We
// may add, change, or remove features before announcing general availability. For
// more information, see Using API keys (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
// .
func (c *Client) DescribeKey(ctx context.Context, params *DescribeKeyInput, optFns ...func(*Options)) (*DescribeKeyOutput, error) {
	if params == nil {
		params = &DescribeKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeKey", params, optFns, c.addOperationDescribeKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeKeyInput struct {

	// The name of the API key resource.
	//
	// This member is required.
	KeyName *string

	noSmithyDocumentSerde
}

type DescribeKeyOutput struct {

	// The timestamp for when the API key resource was created in  ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html)
	// format: YYYY-MM-DDThh:mm:ss.sssZ .
	//
	// This member is required.
	CreateTime *time.Time

	// The timestamp for when the API key resource will expire in  ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html)
	// format: YYYY-MM-DDThh:mm:ss.sssZ .
	//
	// This member is required.
	ExpireTime *time.Time

	// The key value/string of an API key.
	//
	// This member is required.
	Key *string

	// The Amazon Resource Name (ARN) for the API key resource. Used when you need to
	// specify a resource across all Amazon Web Services.
	//   - Format example: arn:aws:geo:region:account-id:key/ExampleKey
	//
	// This member is required.
	KeyArn *string

	// The name of the API key resource.
	//
	// This member is required.
	KeyName *string

	// API Restrictions on the allowed actions, resources, and referers for an API key
	// resource.
	//
	// This member is required.
	Restrictions *types.ApiKeyRestrictions

	// The timestamp for when the API key resource was last updated in  ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html)
	// format: YYYY-MM-DDThh:mm:ss.sssZ .
	//
	// This member is required.
	UpdateTime *time.Time

	// The optional description for the API key resource.
	Description *string

	// Tags associated with the API key resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeKey{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeKeyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeKey(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeKeyMiddleware struct {
}

func (*endpointPrefix_opDescribeKeyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeKeyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "metadata." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeKeyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeKeyMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "DescribeKey",
	}
}
