package hash

import (
	"context"

	"github.com/alexedwards/argon2id"
)

type argon2idService struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func NewArgon2idService(parallelism uint8) HashService {
	return &argon2idService{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: parallelism,
		SaltLength:  16,
		KeyLength:   32,
	}
}

func (s *argon2idService) Hash(ctx context.Context, password string) (string, error) {
	hash, err := argon2id.CreateHash(password, &argon2id.Params{
		Memory:      s.Memory,
		Iterations:  s.Iterations,
		Parallelism: s.Parallelism,
		SaltLength:  s.SaltLength,
		KeyLength:   s.KeyLength,
	})
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (s *argon2idService) Compare(ctx context.Context, password string, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}
