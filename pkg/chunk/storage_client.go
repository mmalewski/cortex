package chunk

import "context"

// StorageClient is a client for the persistent storage for Cortex. (e.g. DynamoDB + S3).
type StorageClient interface {
	// For the write path.
	NewWriteBatch() WriteBatch
	BatchWrite(context.Context, WriteBatch) error

	// For the read path.
	QueryPages(ctx context.Context, query IndexQuery, callback func(result ReadBatch) (shouldContinue bool)) error

	// For fixups
	QueryRows(ctx context.Context, tableName, prefix string, callback func(result ReadBatch) (shouldContinue bool)) error
	NewDeleteBatch() WriteBatch

	// For storing and retrieving chunks.
	PutChunks(ctx context.Context, chunks []Chunk) error
	GetChunks(ctx context.Context, chunks []Chunk) ([]Chunk, error)
}

// WriteBatch represents a batch of writes.
type WriteBatch interface {
	Add(tableName, hashValue string, rangeValue []byte, value []byte)
}

// ReadBatch represents the results of a QueryPages.
type ReadBatch interface {
	Len() int
	HashValue(index int) string // NB only implemented for the BigTable client
	RangeValue(index int) []byte
	Value(index int) []byte
}
