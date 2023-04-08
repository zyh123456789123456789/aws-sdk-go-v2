// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants permissions to the principal to access metadata in the Data Catalog and
// data organized in underlying data storage such as Amazon S3. For information
// about permissions, see Security and Access Control to Metadata and Data (https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html)
// .
func (c *Client) GrantPermissions(ctx context.Context, params *GrantPermissionsInput, optFns ...func(*Options)) (*GrantPermissionsOutput, error) {
	if params == nil {
		params = &GrantPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GrantPermissions", params, optFns, c.addOperationGrantPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GrantPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GrantPermissionsInput struct {

	// The permissions granted to the principal on the resource. Lake Formation
	// defines privileges to grant and revoke access to metadata in the Data Catalog
	// and data organized in underlying data storage such as Amazon S3. Lake Formation
	// requires that each principal be authorized to perform a specific task on Lake
	// Formation resources.
	//
	// This member is required.
	Permissions []types.Permission

	// The principal to be granted the permissions on the resource. Supported
	// principals are IAM users or IAM roles, and they are defined by their principal
	// type and their ARN. Note that if you define a resource with a particular ARN,
	// then later delete, and recreate a resource with that same ARN, the resource
	// maintains the permissions already granted.
	//
	// This member is required.
	Principal *types.DataLakePrincipal

	// The resource to which permissions are to be granted. Resources in Lake
	// Formation are the Data Catalog, databases, and tables.
	//
	// This member is required.
	Resource *types.Resource

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your Lake Formation
	// environment.
	CatalogId *string

	// Indicates a list of the granted permissions that the principal may pass to
	// other users. These permissions may only be a subset of the permissions granted
	// in the Privileges .
	PermissionsWithGrantOption []types.Permission

	noSmithyDocumentSerde
}

type GrantPermissionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGrantPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGrantPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGrantPermissions{}, middleware.After)
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
	if err = addOpGrantPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGrantPermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGrantPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GrantPermissions",
	}
}
