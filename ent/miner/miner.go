// Code generated by entc, DO NOT EDIT.

package miner

const (
	// Label holds the string label denoting the miner type in the database.
	Label = "miner"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStateRoot holds the string denoting the state_root field in the database.
	FieldStateRoot = "state_root"
	// FieldMinerID holds the string denoting the miner_id field in the database.
	FieldMinerID = "miner_id"
	// FieldOwnerAddr holds the string denoting the owner_addr field in the database.
	FieldOwnerAddr = "owner_addr"
	// FieldWorkerAddr holds the string denoting the worker_addr field in the database.
	FieldWorkerAddr = "worker_addr"
	// FieldPeerID holds the string denoting the peer_id field in the database.
	FieldPeerID = "peer_id"
	// FieldSectorSize holds the string denoting the sector_size field in the database.
	FieldSectorSize = "sector_size"

	// Table holds the table name of the miner in the database.
	Table = "miners"
)

// Columns holds all SQL columns for miner fields.
var Columns = []string{
	FieldID,
	FieldStateRoot,
	FieldMinerID,
	FieldOwnerAddr,
	FieldWorkerAddr,
	FieldPeerID,
	FieldSectorSize,
}

var (
	// MinerIDValidator is a validator for the "miner_id" field. It is called by the builders before save.
	MinerIDValidator func(string) error
	// OwnerAddrValidator is a validator for the "owner_addr" field. It is called by the builders before save.
	OwnerAddrValidator func(string) error
	// WorkerAddrValidator is a validator for the "worker_addr" field. It is called by the builders before save.
	WorkerAddrValidator func(string) error
	// SectorSizeValidator is a validator for the "sector_size" field. It is called by the builders before save.
	SectorSizeValidator func(string) error
)
