package dbadapter

import (
	"io"

	dbm "github.com/tendermint/tm-db"

	"github.com/pokt-network/pocket-core/store/cachekv"
	"github.com/pokt-network/pocket-core/store/tracekv"
	"github.com/pokt-network/pocket-core/store/types"
)

// Wrapper type for dbm.Db with implementation of KVStore
type Store struct {
	dbm.DB
}

func (dsa Store) Get(key []byte) []byte {
	panic("implement me")
}

func (dsa Store) Has(key []byte) bool {
	panic("implement me")
}

func (dsa Store) Set(key, value []byte) {
	panic("implement me")
}

func (dsa Store) Delete(key []byte) {
	panic("implement me")
}

func (dsa Store) Iterator(start, end []byte) types.Iterator {
	panic("implement me")
}

func (dsa Store) ReverseIterator(start, end []byte) types.Iterator {
	panic("implement me")
}

// GetStoreType returns the type of the store.
func (Store) GetStoreType() types.StoreType {
	return types.StoreTypeDB
}

// CacheWrap cache wraps the underlying store.
func (dsa Store) CacheWrap() types.CacheWrap {
	return cachekv.NewStore(dsa)
}

// CacheWrapWithTrace implements KVStore.
func (dsa Store) CacheWrapWithTrace(w io.Writer, tc types.TraceContext) types.CacheWrap {
	return cachekv.NewStore(tracekv.NewStore(dsa, w, tc))
}

// dbm.DB implements KVStore so we can CacheKVStore it.
var _ types.KVStore = Store{}
