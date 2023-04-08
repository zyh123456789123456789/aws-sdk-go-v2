// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For a given input image, first detects the largest face in the image, and then
// searches the specified collection for matching faces. The operation compares the
// features of the input face with faces in the specified collection. To search for
// all faces in an input image, you might first call the IndexFaces operation, and
// then use the face IDs returned in subsequent calls to the SearchFaces
// operation. You can also call the DetectFaces operation and use the bounding
// boxes in the response to make face crops, which then you can pass in to the
// SearchFacesByImage operation. You pass the input image either as base64-encoded
// image bytes or as a reference to an image in an Amazon S3 bucket. If you use the
// AWS CLI to call Amazon Rekognition operations, passing image bytes is not
// supported. The image must be either a PNG or JPEG formatted file. The response
// returns an array of faces that match, ordered by similarity score with the
// highest similarity first. More specifically, it is an array of metadata for each
// face match found. Along with the metadata, the response also includes a
// similarity indicating how similar the face is to the input face. In the
// response, the operation also returns the bounding box (and a confidence level
// that the bounding box contains a face) of the face that Amazon Rekognition used
// for the input image. If no faces are detected in the input image,
// SearchFacesByImage returns an InvalidParameterException error. For an example,
// Searching for a Face Using an Image in the Amazon Rekognition Developer Guide.
// The QualityFilter input parameter allows you to filter out detected faces that
// don’t meet a required quality bar. The quality bar is based on a variety of
// common use cases. Use QualityFilter to set the quality bar for filtering by
// specifying LOW , MEDIUM , or HIGH . If you do not want to filter detected faces,
// specify NONE . The default value is NONE . To use quality filtering, you need a
// collection associated with version 3 of the face model or higher. To get the
// version of the face model associated with a collection, call DescribeCollection
// . This operation requires permissions to perform the
// rekognition:SearchFacesByImage action.
func (c *Client) SearchFacesByImage(ctx context.Context, params *SearchFacesByImageInput, optFns ...func(*Options)) (*SearchFacesByImageOutput, error) {
	if params == nil {
		params = &SearchFacesByImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchFacesByImage", params, optFns, c.addOperationSearchFacesByImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchFacesByImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchFacesByImageInput struct {

	// ID of the collection to search.
	//
	// This member is required.
	CollectionId *string

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	Image *types.Image

	// (Optional) Specifies the minimum confidence in the face match to return. For
	// example, don't return any matches where confidence in matches is less than 70%.
	// The default value is 80%.
	FaceMatchThreshold *float32

	// Maximum number of faces to return. The operation returns the maximum number of
	// faces with the highest confidence in the match.
	MaxFaces *int32

	// A filter that specifies a quality bar for how much filtering is done to
	// identify faces. Filtered faces aren't searched for in the collection. If you
	// specify AUTO , Amazon Rekognition chooses the quality bar. If you specify LOW ,
	// MEDIUM , or HIGH , filtering removes all faces that don’t meet the chosen
	// quality bar. The quality bar is based on a variety of common use cases.
	// Low-quality detections can occur for a number of reasons. Some examples are an
	// object that's misidentified as a face, a face that's too blurry, or a face with
	// a pose that's too extreme to use. If you specify NONE , no filtering is
	// performed. The default value is NONE . To use quality filtering, the collection
	// you are using must be associated with version 3 of the face model or higher.
	QualityFilter types.QualityFilter

	noSmithyDocumentSerde
}

type SearchFacesByImageOutput struct {

	// An array of faces that match the input face, along with the confidence in the
	// match.
	FaceMatches []types.FaceMatch

	// Version number of the face detection model associated with the input collection
	// ( CollectionId ).
	FaceModelVersion *string

	// The bounding box around the face in the input image that Amazon Rekognition
	// used for the search.
	SearchedFaceBoundingBox *types.BoundingBox

	// The level of confidence that the searchedFaceBoundingBox , contains a face.
	SearchedFaceConfidence *float32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchFacesByImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchFacesByImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchFacesByImage{}, middleware.After)
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
	if err = addOpSearchFacesByImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchFacesByImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchFacesByImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "SearchFacesByImage",
	}
}
