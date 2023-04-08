// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the contents of a Spark notebook.
func (c *Client) UpdateNotebook(ctx context.Context, params *UpdateNotebookInput, optFns ...func(*Options)) (*UpdateNotebookOutput, error) {
	if params == nil {
		params = &UpdateNotebookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNotebook", params, optFns, c.addOperationUpdateNotebookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNotebookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNotebookInput struct {

	// The ID of the notebook to update.
	//
	// This member is required.
	NotebookId *string

	// The updated content for the notebook.
	//
	// This member is required.
	Payload *string

	// The notebook content type. Currently, the only valid type is IPYNB .
	//
	// This member is required.
	Type types.NotebookType

	// A unique case-sensitive string used to ensure the request to create the
	// notebook is idempotent (executes only once). This token is listed as not
	// required because Amazon Web Services SDKs (for example the Amazon Web Services
	// SDK for Java) auto-generate the token for you. If you are not using the Amazon
	// Web Services SDK or the Amazon Web Services CLI, you must provide this token or
	// the action will fail.
	ClientRequestToken *string

	// The active notebook session ID. Required if the notebook has an active session.
	SessionId *string

	noSmithyDocumentSerde
}

type UpdateNotebookOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNotebookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNotebook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNotebook{}, middleware.After)
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
	if err = addOpUpdateNotebookValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotebook(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNotebook(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "UpdateNotebook",
	}
}
