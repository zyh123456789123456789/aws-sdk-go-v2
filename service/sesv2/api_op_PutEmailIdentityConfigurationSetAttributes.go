// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to associate a configuration set with an email identity.
func (c *Client) PutEmailIdentityConfigurationSetAttributes(ctx context.Context, params *PutEmailIdentityConfigurationSetAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityConfigurationSetAttributesOutput, error) {
	if params == nil {
		params = &PutEmailIdentityConfigurationSetAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEmailIdentityConfigurationSetAttributes", params, optFns, c.addOperationPutEmailIdentityConfigurationSetAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEmailIdentityConfigurationSetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to associate a configuration set with an email identity.
type PutEmailIdentityConfigurationSetAttributesInput struct {

	// The email address or domain to associate with a configuration set.
	//
	// This member is required.
	EmailIdentity *string

	// The configuration set to associate with an email identity.
	ConfigurationSetName *string

	noSmithyDocumentSerde
}

// If the action is successful, the service sends back an HTTP 200 response with
// an empty HTTP body.
type PutEmailIdentityConfigurationSetAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEmailIdentityConfigurationSetAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEmailIdentityConfigurationSetAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEmailIdentityConfigurationSetAttributes{}, middleware.After)
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
	if err = addOpPutEmailIdentityConfigurationSetAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEmailIdentityConfigurationSetAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEmailIdentityConfigurationSetAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutEmailIdentityConfigurationSetAttributes",
	}
}
