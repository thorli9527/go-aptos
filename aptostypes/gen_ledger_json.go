// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package aptostypes

import (
	"encoding/json"
	"errors"
)

var _ = (*ledgerInfoMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (l LedgerInfo) MarshalJSON() ([]byte, error) {
	type LedgerInfo struct {
		ChainId         int        `json:"chain_id"`
		LedgerVersion   jsonUint64 `json:"ledger_version" gencodec:"required"`
		LedgerTimestamp jsonUint64 `json:"ledger_timestamp" gencodec:"required"`
		BlockHeight 	jsonUint64 `json:"block_height" gencodec:"required"`
		Epoch           jsonUint64 `json:"epoch"`
		NodeRole        string     `json:"node_role"`
	}
	var enc LedgerInfo
	enc.ChainId = l.ChainId
	enc.LedgerVersion = jsonUint64(l.LedgerVersion)
	enc.LedgerTimestamp = jsonUint64(l.LedgerTimestamp)
	enc.BlockHeight = jsonUint64(l.BlockHeight)
	enc.Epoch = jsonUint64(l.Epoch)
	enc.NodeRole = l.NodeRole
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (l *LedgerInfo) UnmarshalJSON(input []byte) error {
	type LedgerInfo struct {
		ChainId         *int        `json:"chain_id"`
		LedgerVersion   *jsonUint64 `json:"ledger_version" gencodec:"required"`
		LedgerTimestamp *jsonUint64 `json:"ledger_timestamp" gencodec:"required"`
		BlockHeight 	*jsonUint64 `json:"block_height" gencodec:"required"`
		Epoch           *jsonUint64 `json:"epoch"`
		NodeRole        *string     `json:"node_role"`
	}
	var dec LedgerInfo
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ChainId != nil {
		l.ChainId = *dec.ChainId
	}
	if dec.LedgerVersion == nil {
		return errors.New("missing required field 'ledger_version' for LedgerInfo")
	}
	l.LedgerVersion = uint64(*dec.LedgerVersion)
	if dec.LedgerTimestamp == nil {
		return errors.New("missing required field 'ledger_timestamp' for LedgerInfo")
	}
	l.LedgerTimestamp = uint64(*dec.LedgerTimestamp)
	if dec.BlockHeight == nil {
		return errors.New("missing required field 'block_height' for LedgerInfo")
	}
	l.BlockHeight = uint64(*dec.BlockHeight)
	if dec.Epoch != nil {
		l.Epoch = uint64(*dec.Epoch)
	}
	if dec.NodeRole != nil {
		l.NodeRole = *dec.NodeRole
	}
	return nil
}
