package warppipe

import "strings"

// ChangesetKind is the type for changeset kinds
type ChangesetKind string

// ChangsetKind constants
const (
	ChangesetKindInsert ChangesetKind = "insert"
	ChangesetKindUpdate ChangesetKind = "update"
	ChangesetKindDelete ChangesetKind = "delete"
)

// ParseChangesetKind parses a changeset kind from a string.
func ParseChangesetKind(kind string) ChangesetKind {
	switch strings.ToLower(kind) {
	case "insert":
		return ChangesetKindInsert
	case "update":
		return ChangesetKindUpdate
	case "delete":
		return ChangesetKindDelete
	default:
		// TODO: should this error?
		return ""
	}
}

// Changeset represents a changeset for a record on a Postgres table.
type Changeset struct {
	Kind      ChangesetKind      `json:"kind"`
	Schema    string             `json:"schema"`
	Table     string             `json:"table"`
	NewValues []*ChangesetColumn `json:"new_values"`
	OldValues []*ChangesetColumn `json:"old_values"`
}

// ChangesetColumn represents a type and value for a column in a changset.
type ChangesetColumn struct {
	Column string      `json:"column"`
	Value  interface{} `json:"value"`
	Type   string      `json:"type"`
}