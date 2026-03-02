package meta

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// NullUUID is a wrapper around uuid.UUID that supports unmarshaling from an empty string as nil.
// It also tracks if the field was present in the JSON.
type NullUUID struct {
	ID    *uuid.UUID
	Valid bool // True if the field was present in the JSON (even if null or empty)
}

func (n *NullUUID) UnmarshalJSON(data []byte) error {
	n.Valid = true

	// Handle explicitly "null"
	if string(data) == "null" {
		n.ID = nil
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		// Try unmarshaling as pointer to uuid directly (for safety/compatibility)
		var u *uuid.UUID
		if err := json.Unmarshal(data, &u); err != nil {
			return err
		}
		n.ID = u
		return nil
	}

	if s == "" {
		n.ID = nil
		return nil
	}

	u, err := uuid.Parse(s)
	if err != nil {
		return fmt.Errorf("invalid UUID length: %d", len(s))
	}

	n.ID = &u
	return nil
}

func (n NullUUID) MarshalJSON() ([]byte, error) {
	if n.ID == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(n.ID)
}

func (n NullUUID) String() string {
	if n.ID == nil {
		return ""
	}
	return n.ID.String()
}
