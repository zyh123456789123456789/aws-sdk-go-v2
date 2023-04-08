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

// Updates the data catalog that has the specified name.
func (c *Client) UpdateDataCatalog(ctx context.Context, params *UpdateDataCatalogInput, optFns ...func(*Options)) (*UpdateDataCatalogOutput, error) {
	if params == nil {
		params = &UpdateDataCatalogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataCatalog", params, optFns, c.addOperationUpdateDataCatalogMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataCatalogInput struct {

	// The name of the data catalog to update. The catalog name must be unique for the
	// Amazon Web Services account and can use a maximum of 127 alphanumeric,
	// underscore, at sign, or hyphen characters. The remainder of the length
	// constraint of 256 is reserved for use by Athena.
	//
	// This member is required.
	Name *string

	// Specifies the type of data catalog to update. Specify LAMBDA for a federated
	// catalog, HIVE for an external hive metastore, or GLUE for an Glue Data Catalog.
	//
	// This member is required.
	Type types.DataCatalogType

	// New or modified text that describes the data catalog.
	Description *string

	// Specifies the Lambda function or functions to use for updating the data
	// catalog. This is a mapping whose values depend on the catalog type.
	//   - For the HIVE data catalog type, use the following syntax. The
	//   metadata-function parameter is required. The sdk-version parameter is optional
	//   and defaults to the currently supported version.
	//   metadata-function=lambda_arn, sdk-version=version_number
	//   - For the LAMBDA data catalog type, use one of the following sets of required
	//   parameters, but not both.
	//   - If you have one Lambda function that processes metadata and another for
	//   reading the actual data, use the following syntax. Both parameters are required.
	//   metadata-function=lambda_arn, record-function=lambda_arn
	//   - If you have a composite Lambda function that processes both metadata and
	//   data, use the following syntax to specify your Lambda function.
	//   function=lambda_arn
	Parameters map[string]string

	noSmithyDocumentSerde
}

type UpdateDataCatalogOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataCatalogMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDataCatalog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDataCatalog{}, middleware.After)
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
	if err = addOpUpdateDataCatalogValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataCatalog(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataCatalog(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "UpdateDataCatalog",
	}
}
