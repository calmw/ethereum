// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/calmw/ethereum/common"
	"github.com/calmw/ethereum/common/hexutil"
	"github.com/holiman/uint256"
)

var _ = (*authorizationMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (s SetCodeAuthorization) MarshalJSON() ([]byte, error) {
	type SetCodeAuthorization struct {
		ChainID hexutil.U256   `json:"chainId" gencodec:"required"`
		Address common.Address `json:"address" gencodec:"required"`
		Nonce   hexutil.Uint64 `json:"nonce" gencodec:"required"`
		V       hexutil.Uint64 `json:"yParity" gencodec:"required"`
		R       hexutil.U256   `json:"r" gencodec:"required"`
		S       hexutil.U256   `json:"s" gencodec:"required"`
	}
	var enc SetCodeAuthorization
	enc.ChainID = hexutil.U256(s.ChainID)
	enc.Address = s.Address
	enc.Nonce = hexutil.Uint64(s.Nonce)
	enc.V = hexutil.Uint64(s.V)
	enc.R = hexutil.U256(s.R)
	enc.S = hexutil.U256(s.S)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (s *SetCodeAuthorization) UnmarshalJSON(input []byte) error {
	type SetCodeAuthorization struct {
		ChainID *hexutil.U256   `json:"chainId" gencodec:"required"`
		Address *common.Address `json:"address" gencodec:"required"`
		Nonce   *hexutil.Uint64 `json:"nonce" gencodec:"required"`
		V       *hexutil.Uint64 `json:"yParity" gencodec:"required"`
		R       *hexutil.U256   `json:"r" gencodec:"required"`
		S       *hexutil.U256   `json:"s" gencodec:"required"`
	}
	var dec SetCodeAuthorization
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ChainID == nil {
		return errors.New("missing required field 'chainId' for SetCodeAuthorization")
	}
	s.ChainID = uint256.Int(*dec.ChainID)
	if dec.Address == nil {
		return errors.New("missing required field 'address' for SetCodeAuthorization")
	}
	s.Address = *dec.Address
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' for SetCodeAuthorization")
	}
	s.Nonce = uint64(*dec.Nonce)
	if dec.V == nil {
		return errors.New("missing required field 'yParity' for SetCodeAuthorization")
	}
	s.V = uint8(*dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' for SetCodeAuthorization")
	}
	s.R = uint256.Int(*dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for SetCodeAuthorization")
	}
	s.S = uint256.Int(*dec.S)
	return nil
}
