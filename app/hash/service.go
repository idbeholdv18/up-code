package hash

import "context"

type HashService interface {
	Hash(ctx context.Context, pass string) (string, error)
	Compare(ctx context.Context, passed string, with string) (bool, error)
}
