package hasher

import (
	"crypto/sha512"
	"fmt"

	"github.com/Shteyd/wallet-app/src/backend/internal/core/contract"
	"github.com/Shteyd/wallet-app/src/backend/internal/lib/byteconv"
)

var _ contract.SecretHasher = (*HashAdapter)(nil)

type HashAdapter struct {
	salt []byte
}

func New(salt string) HashAdapter {
	return HashAdapter{
		salt: byteconv.ToSlice(salt),
	}
}

func (h HashAdapter) Hash(raw string) (string, error) {
	const op = "adapter.hasher.hash"

	engine := sha512.New()

	_, err := engine.Write(byteconv.ToSlice(raw))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	secret := engine.Sum(h.salt)

	return byteconv.ToString(secret), nil
}
