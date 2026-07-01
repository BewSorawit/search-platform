package milvus

import (
	"context"
	"errors"
	"net/http"

	merr "github.com/milvus-io/milvus/pkg/v2/util/merr"
)

var (
	ErrInvalidID              = errors.New("id must be non-negative")
	ErrTextRequired           = errors.New("text required")
	ErrTextTooLong            = errors.New("text too long")
	ErrInvalidVectorDimension = errors.New("invalid vector dimension")
)

func HTTPStatusAndMessage(err error) (int, string) {
	switch {
	case err == nil:
		return http.StatusOK, ""
	case errors.Is(err, ErrInvalidID),
		errors.Is(err, ErrTextRequired),
		errors.Is(err, ErrTextTooLong),
		errors.Is(err, ErrInvalidVectorDimension),
		errors.Is(err, merr.ErrParameterInvalid),
		errors.Is(err, merr.ErrInvalidInsertData):
		return http.StatusBadRequest, err.Error()
	case errors.Is(err, merr.ErrServiceNotReady),
		errors.Is(err, merr.ErrCollectionNotLoaded):
		return http.StatusServiceUnavailable, err.Error()
	case errors.Is(err, merr.ErrCollectionNotFound):
		return http.StatusServiceUnavailable, "collection not available"
	case errors.Is(err, merr.ErrCollectionSchemaMismatch):
		return http.StatusInternalServerError, "collection schema mismatch"
	case errors.Is(err, context.DeadlineExceeded):
		return http.StatusGatewayTimeout, "milvus request timed out"
	case errors.Is(err, context.Canceled):
		return http.StatusRequestTimeout, "request canceled"
	default:
		return http.StatusInternalServerError, "upsert failed"
	}
}
