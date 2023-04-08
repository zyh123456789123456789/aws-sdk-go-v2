// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes an existing configuration set. A configuration set is a set of rules
// that you apply to voice and SMS messages that you send. In a configuration set,
// you can specify a destination for specific types of events related to voice and
// SMS messages.
func (c *Client) DeleteConfigurationSet(ctx context.Context, params *DeleteConfigurationSetInput, optFns ...func(*Options)) (*DeleteConfigurationSetOutput, error) {
	if params == nil {
		params = &DeleteConfigurationSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfigurationSet", params, optFns, c.addOperationDeleteConfigurationSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigurationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConfigurationSetInput struct {

	// The name of the configuration set or the configuration set ARN that you want to
	// delete. The ConfigurationSetName and ConfigurationSetArn can be found using the
	// DescribeConfigurationSets action.
	//
	// This member is required.
	ConfigurationSetName *string

	noSmithyDocumentSerde
}

type DeleteConfigurationSetOutput struct {

	// The Amazon Resource Name (ARN) of the deleted configuration set.
	ConfigurationSetArn *string

	// The name of the deleted configuration set.
	ConfigurationSetName *string

	// The time that the deleted configuration set was created in UNIX epoch time (https://www.epochconverter.com/)
	// format.
	CreatedTimestamp *time.Time

	// The default message type of the configuration set that was deleted.
	DefaultMessageType types.MessageType

	// The default Sender ID of the configuration set that was deleted.
	DefaultSenderId *string

	// An array of any EventDestination objects that were associated with the deleted
	// configuration set.
	EventDestinations []types.EventDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteConfigurationSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteConfigurationSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteConfigurationSet{}, middleware.After)
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
	if err = addOpDeleteConfigurationSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteConfigurationSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DeleteConfigurationSet",
	}
}
