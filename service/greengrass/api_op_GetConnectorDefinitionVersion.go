// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about a connector definition version, including the
// connectors that the version contains. Connectors are prebuilt modules that
// interact with local infrastructure, device protocols, AWS, and other cloud
// services.
func (c *Client) GetConnectorDefinitionVersion(ctx context.Context, params *GetConnectorDefinitionVersionInput, optFns ...func(*Options)) (*GetConnectorDefinitionVersionOutput, error) {
	if params == nil {
		params = &GetConnectorDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConnectorDefinitionVersion", params, optFns, c.addOperationGetConnectorDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConnectorDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConnectorDefinitionVersionInput struct {

	// The ID of the connector definition.
	//
	// This member is required.
	ConnectorDefinitionId *string

	// The ID of the connector definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListConnectorDefinitionVersions'' requests. If the version is the last one
	// that was associated with a connector definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	//
	// This member is required.
	ConnectorDefinitionVersionId *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetConnectorDefinitionVersionOutput struct {

	// The ARN of the connector definition version.
	Arn *string

	// The time, in milliseconds since the epoch, when the connector definition
	// version was created.
	CreationTimestamp *string

	// Information about the connector definition version.
	Definition *types.ConnectorDefinitionVersion

	// The ID of the connector definition version.
	Id *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// The version of the connector definition version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConnectorDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConnectorDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConnectorDefinitionVersion{}, middleware.After)
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
	if err = addOpGetConnectorDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConnectorDefinitionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConnectorDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetConnectorDefinitionVersion",
	}
}
