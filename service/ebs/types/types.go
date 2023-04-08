// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A block of data in an Amazon Elastic Block Store snapshot.
type Block struct {

	// The block index.
	BlockIndex *int32

	// The block token for the block index.
	BlockToken *string

	noSmithyDocumentSerde
}

// A block of data in an Amazon Elastic Block Store snapshot that is different
// from another snapshot of the same volume/snapshot lineage.
type ChangedBlock struct {

	// The block index.
	BlockIndex *int32

	// The block token for the block index of the FirstSnapshotId specified in the
	// ListChangedBlocks operation. This value is absent if the first snapshot does not
	// have the changed block that is on the second snapshot.
	FirstBlockToken *string

	// The block token for the block index of the SecondSnapshotId specified in the
	// ListChangedBlocks operation.
	SecondBlockToken *string

	noSmithyDocumentSerde
}

// Describes a tag.
type Tag struct {

	// The key of the tag.
	Key *string

	// The value of the tag.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
