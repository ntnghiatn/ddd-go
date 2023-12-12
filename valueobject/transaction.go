package valueobject

import "github.com/google/uuid"

// Tranaction is
type Transaction struct {
	amount int
	from   uuid.UUID
	to     uuid.UUID
}
