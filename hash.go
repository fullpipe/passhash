package passhash

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	cost int
}

func NewHash() *Hasher {
	return &Hasher{
		cost: 16,
	}
}

//Cost setup hash cost [4, 31]. Default 16
func (h *Hasher) Cost(cost int) *Hasher {
	h.cost = cost

	return h
}

//HashPassword hashes plain password with bcrypt
func (h *Hasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)

	return string(bytes), err
}

//CheckPasswordHash checks plain password with hashed version
func (h *Hasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
