// Generated by: main
// TypeWriter: wrapper
// Directive: +gen on PrivKeyInner

package crypto

import (
	"github.com/Sukhavati-Labs/skt-tendermint/go-wire/data"
)

// Auto-generated adapters for happily unmarshaling interfaces
// Apache License 2.0
// Copyright (c) 2017 Ethan Frey (ethan.frey@tendermint.com)

type PrivKey struct {
	PrivKeyInner "json:\"unwrap\""
}

var PrivKeyMapper = data.NewMapper(PrivKey{})

func (h PrivKey) MarshalJSON() ([]byte, error) {
	return PrivKeyMapper.ToJSON(h.PrivKeyInner)
}

func (h *PrivKey) UnmarshalJSON(data []byte) (err error) {
	parsed, err := PrivKeyMapper.FromJSON(data)
	if err == nil && parsed != nil {
		h.PrivKeyInner = parsed.(PrivKeyInner)
	}
	return err
}

// Unwrap recovers the concrete interface safely (regardless of levels of embeds)
func (h PrivKey) Unwrap() PrivKeyInner {
	hi := h.PrivKeyInner
	for wrap, ok := hi.(PrivKey); ok; wrap, ok = hi.(PrivKey) {
		hi = wrap.PrivKeyInner
	}
	return hi
}

func (h PrivKey) Empty() bool {
	return h.PrivKeyInner == nil
}

/*** below are bindings for each implementation ***/

func init() {
	PrivKeyMapper.RegisterImplementation(PrivKeyEd25519{}, "ed25519", 0x1)
}

func (hi PrivKeyEd25519) Wrap() PrivKey {
	return PrivKey{hi}
}

func init() {
	PrivKeyMapper.RegisterImplementation(PrivKeySecp256k1{}, "secp256k1", 0x2)
}

func (hi PrivKeySecp256k1) Wrap() PrivKey {
	return PrivKey{hi}
}
