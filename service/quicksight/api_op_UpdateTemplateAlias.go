// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the template alias of a template.
func (c *Client) UpdateTemplateAlias(ctx context.Context, params *UpdateTemplateAliasInput, optFns ...func(*Options)) (*UpdateTemplateAliasOutput, error) {
	if params == nil {
		params = &UpdateTemplateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTemplateAlias", params, optFns, c.addOperationUpdateTemplateAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTemplateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTemplateAliasInput struct {

	// The alias of the template that you want to update. If you name a specific
	// alias, you update the version that the alias points to. You can specify the
	// latest version of the template by providing the keyword $LATEST in the AliasName
	// parameter. The keyword $PUBLISHED doesn't apply to templates.
	//
	// This member is required.
	AliasName *string

	// The ID of the Amazon Web Services account that contains the template alias that
	// you're updating.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the template.
	//
	// This member is required.
	TemplateId *string

	// The version number of the template.
	//
	// This member is required.
	TemplateVersionNumber *int64

	noSmithyDocumentSerde
}

type UpdateTemplateAliasOutput struct {

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The template alias.
	TemplateAlias *types.TemplateAlias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTemplateAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTemplateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTemplateAlias{}, middleware.After)
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
	if err = addOpUpdateTemplateAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTemplateAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTemplateAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateTemplateAlias",
	}
}
