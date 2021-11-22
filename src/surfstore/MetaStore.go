package surfstore

import (
	"net/rpc"
)

type MetaStore struct {
	FileMetaMap    map[string]FileMetaData
	BlockStoreRing ConsistentHashRing
}

func (m *MetaStore) GetFileInfoMap(succ *bool, serverFileInfoMap *map[string]FileMetaData) error {
	// copy your code implemented in project 3 here
	panic("todo")
}

func (m *MetaStore) UpdateFile(fileMetaData *FileMetaData, latestVersion *int) (err error) {
	// copy your code implemented in project 3 here
	panic("todo")
}

// Given an input hashlist, returns a mapping of BlockStore addresses to hashlists.
func (m *MetaStore) GetBlockStoreMap(blockHashesIn []string, blockStoreMap *map[string][]string) error {
	// this should be different from your project 3 implementation. Now we have multiple
	// Blockstore servers instead of one Blockstore server in project 3. For each blockHash in
	// blockHashesIn, you want to find the BlockStore server it is in using consistent hash ring.
	panic("todo")
}

// Add the specified BlockStore node to the cluster and migrate the blocks
func (m *MetaStore) AddNode(nodeAddr string, succ *bool) error {
	// compute node index
	panic("todo")

	// find successor node
	panic("todo")

	// call RPC to migrate blocks from successor node to this node
	inst := MigrationInstruction{
		LowerIndex: "todo",
		UpperIndex: "todo",
		DestAddr:   "todo",
	}

	// connect to the server
	conn, e := rpc.DialHTTP("tcp", "todo")
	if e != nil {
		return e
	}

	// perform the call
	e = conn.Call("BlockStore.MigrateBlocks", inst, succ)
	if e != nil {
		conn.Close()
		return e
	}

	// deal with added node in consistent hash ring data structure
	panic("todo")

	// close the connection
	return conn.Close()
}

// Remove the specified BlockStore node from the cluster and migrate the blocks
func (m *MetaStore) RemoveNode(nodeAddr string, succ *bool) error {
	// compute node index
	panic("todo")

	// find successor node
	panic("todo")

	// call RPC to migrate blocks from successor node to this node
	inst := MigrationInstruction{
		LowerIndex: "todo",
		UpperIndex: "todo",
		DestAddr:   "todo",
	}

	// connect to the server
	conn, e := rpc.DialHTTP("tcp", "todo")
	if e != nil {
		return e
	}

	// perform the call
	e = conn.Call("BlockStore.MigrateBlocks", inst, succ)
	if e != nil {
		conn.Close()
		return e
	}

	// deal with removed node
	panic("todo")

	// close the connection
	return conn.Close()
}

var _ MetaStoreInterface = new(MetaStore)

func NewMetaStore(blockStoreRing ConsistentHashRing) MetaStore {
	return MetaStore{
		FileMetaMap:    map[string]FileMetaData{},
		BlockStoreRing: blockStoreRing,
	}
}
