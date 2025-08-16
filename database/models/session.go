package models

import (
	"encoding/base64"
	"fmt"
	"time"

	"cosmossdk.io/math"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	sentinelsdk "github.com/sentinel-official/sentinel-go-sdk/types"
	sentinelhub "github.com/sentinel-official/sentinelhub/v12/types"
	"github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
	"gorm.io/gorm"
)

// Session represents a session record in the database.
type Session struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"` // Timestamp when the record was created
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"` // Timestamp when the record was last updated

	AccAddr     string `gorm:"column:acc_addr;not null"`                 // Account address, cannot be null
	Duration    int64  `gorm:"column:duration;not null"`                 // Duration of the session in nanoseconds
	ID          uint64 `gorm:"column:id;not null;primaryKey"`            // Unique identifier for the session
	MaxBytes    string `gorm:"column:max_bytes;not null"`                // Maximum bytes represented as a string
	MaxDuration int64  `gorm:"column:max_duration;not null"`             // Maximum allowed duration for the session in nanoseconds
	NodeAddr    string `gorm:"column:node_addr;not null"`                // Address of the node associated with the session
	PeerID      string `gorm:"column:peer_id;not null;uniqueIndex"`      // Unique identifier for the peer (e.g., public key, email, or name depending on protocol)
	PeerRequest string `gorm:"column:peer_request;not null;uniqueIndex"` // Unique peer request for the session, indexed and cannot be null
	RxBytes     string `gorm:"column:rx_bytes;not null"`                 // Rx bytes represented as a string
	ServiceType string `gorm:"column:service_type;not null"`             // Type of service for the session
	Signature   string `gorm:"column:signature;not null"`                // Signature associated with the session
	TxBytes     string `gorm:"column:tx_bytes;not null"`                 // Tx bytes represented as a string
}

// NewSession creates and returns a new instance of the Session struct with default values.
func NewSession() *Session {
	return &Session{}
}

// WithAccAddr sets the AccAddr field and returns the updated Session instance.
func (s *Session) WithAccAddr(v cosmossdk.AccAddress) *Session {
	s.AccAddr = v.String()
	return s
}

// WithDuration sets the Duration field from time.Duration and returns the updated Session instance.
func (s *Session) WithDuration(v time.Duration) *Session {
	s.Duration = v.Nanoseconds()
	return s
}

// WithID sets the ID field and returns the updated Session instance.
func (s *Session) WithID(v uint64) *Session {
	s.ID = v
	return s
}

// WithMaxBytes sets the MaxBytes field from math.Int and returns the updated Session instance.
func (s *Session) WithMaxBytes(v math.Int) *Session {
	s.MaxBytes = v.String()
	return s
}

// WithMaxDuration sets the MaxDuration field from time.Duration and returns the updated Session instance.
func (s *Session) WithMaxDuration(v time.Duration) *Session {
	s.MaxDuration = v.Nanoseconds()
	return s
}

// WithNodeAddr sets the NodeAddr field and returns the updated Session instance.
func (s *Session) WithNodeAddr(v sentinelhub.NodeAddress) *Session {
	s.NodeAddr = v.String()
	return s
}

// WithPeerID sets the PeerID field and returns the updated Session instance.
func (s *Session) WithPeerID(v string) *Session {
	s.PeerID = v
	return s
}

// WithPeerRequest sets the PeerRequest field and returns the updated Session instance.
func (s *Session) WithPeerRequest(v []byte) *Session {
	s.PeerRequest = base64.StdEncoding.EncodeToString(v)
	return s
}

// WithRxBytes sets the RxBytes field from math.Int and returns the updated Session instance.
func (s *Session) WithRxBytes(v math.Int) *Session {
	s.RxBytes = v.String()
	return s
}

// WithServiceType sets the ServiceType field and returns the updated Session instance.
func (s *Session) WithServiceType(v sentinelsdk.ServiceType) *Session {
	s.ServiceType = v.String()
	return s
}

// WithSignature sets the Signature field and returns the updated Session instance.
func (s *Session) WithSignature(v []byte) *Session {
	s.Signature = base64.StdEncoding.EncodeToString(v)
	return s
}

// WithTxBytes sets the TxBytes field from math.Int and returns the updated Session instance.
func (s *Session) WithTxBytes(v math.Int) *Session {
	s.TxBytes = v.String()
	return s
}

// GetAccAddr returns the AccAddr field as cosmossdk.AccAddress.
func (s *Session) GetAccAddr() cosmossdk.AccAddress {
	addr, err := cosmossdk.AccAddressFromBech32(s.AccAddr)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetTotalBytes returns the total number of bytes (rx + tx) as math.Int.
func (s *Session) GetTotalBytes() math.Int {
	rxBytes := s.GetRxBytes()
	txBytes := s.GetTxBytes()

	return rxBytes.Add(txBytes)
}

// GetDuration returns the Duration field as time.Duration.
func (s *Session) GetDuration() time.Duration {
	return time.Duration(s.Duration)
}

// GetID returns the ID field.
func (s *Session) GetID() uint64 {
	return s.ID
}

// GetMaxBytes returns the MaxBytes field as math.Int.
func (s *Session) GetMaxBytes() math.Int {
	v, ok := math.NewIntFromString(s.MaxBytes)
	if !ok {
		panic(fmt.Errorf("invalid max_bytes %s", s.MaxBytes))
	}

	return v
}

// GetMaxDuration returns the MaxDuration field as time.Duration.
func (s *Session) GetMaxDuration() time.Duration {
	return time.Duration(s.MaxDuration)
}

// GetNodeAddr returns the NodeAddr field as sentinelhub.NodeAddress.
func (s *Session) GetNodeAddr() sentinelhub.NodeAddress {
	addr, err := sentinelhub.NodeAddressFromBech32(s.NodeAddr)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetPeerID returns the PeerID field.
func (s *Session) GetPeerID() string {
	return s.PeerID
}

// GetPeerRequest returns the PeerRequest field as a decoded byte slice.
func (s *Session) GetPeerRequest() []byte {
	buf, err := base64.StdEncoding.DecodeString(s.PeerRequest)
	if err != nil {
		panic(err)
	}

	return buf
}

// GetRxBytes returns the RxBytes field as math.Int.
func (s *Session) GetRxBytes() math.Int {
	v, ok := math.NewIntFromString(s.RxBytes)
	if !ok {
		panic(fmt.Errorf("invalid rx_bytes %s", s.RxBytes))
	}

	return v
}

// GetServiceType returns the ServiceType field as sentinelsdk.ServiceType.
func (s *Session) GetServiceType() sentinelsdk.ServiceType {
	return sentinelsdk.ServiceTypeFromString(s.ServiceType)
}

// GetSignature returns the Signature field as a byte slice.
func (s *Session) GetSignature() []byte {
	if s.Signature == "" {
		return nil
	}

	buf, err := base64.StdEncoding.DecodeString(s.Signature)
	if err != nil {
		panic(err)
	}

	return buf
}

// GetTxBytes returns the TxBytes field as math.Int.
func (s *Session) GetTxBytes() math.Int {
	v, ok := math.NewIntFromString(s.TxBytes)
	if !ok {
		panic(fmt.Errorf("invalid tx_bytes %s", s.TxBytes))
	}

	return v
}

// BeforeUpdate is a GORM hook that updates the Duration field if relevant fields change.
func (s *Session) BeforeUpdate(db *gorm.DB) (err error) {
	if s.ID == 0 {
		return nil
	}

	if db.Statement.Changed("rx_bytes", "tx_bytes") {
		duration := time.Since(s.CreatedAt).Nanoseconds()
		db.Statement.SetColumn("duration", duration)
	}

	return nil
}

// MsgUpdateSessionRequest creates a MsgUpdateSessionRequest for the session.
func (s *Session) MsgUpdateSessionRequest() *v3.MsgUpdateSessionRequest {
	return v3.NewMsgUpdateSessionRequest(
		s.GetNodeAddr(),
		s.GetID(),
		s.GetTxBytes(),
		s.GetRxBytes(),
		s.GetDuration(),
		s.GetSignature(),
	)
}
