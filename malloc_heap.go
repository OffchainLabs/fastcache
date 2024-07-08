//go:build appengine || windows || wasm
// +build appengine windows wasm

package fastcache

func getChunk() []byte {
	return make([]byte, chunkSize)
}

func putChunk(chunk []byte) {
	// No-op.
}
