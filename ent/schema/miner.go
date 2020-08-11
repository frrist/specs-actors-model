package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Miner holds the schema definition for the Miner entity.
type Miner struct {
	ent.Schema
}

// Fields of the Miner.
func (Miner) Fields() []ent.Field {
	return []ent.Field{
		field.String("state_root"),
		field.String("miner_id").
			NotEmpty().
			Unique(),
		field.String("owner_addr").
			NotEmpty(),
		field.String("worker_addr").
			NotEmpty(),
		field.String("peer_id"),
		field.String("sector_size").
			NotEmpty(),
	}
}

// Edges of the Miner.
func (Miner) Edges() []ent.Edge {
	return nil
}
