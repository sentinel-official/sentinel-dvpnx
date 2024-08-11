package models

import (
	"encoding/base64"
	"fmt"
	"time"

	cosmossdkmath "cosmossdk.io/math"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	base "github.com/sentinel-official/hub/v12/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
	sdk "github.com/sentinel-official/sentinel-go-sdk/types"
	"gorm.io/gorm"
)

// Session represents a session record in the database.
type Session struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"` // Timestamp when the record was created
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"` // Timestamp when the record was last updated

	ID            uint64 `gorm:"column:id;not null;primaryKey"`        // Unique identifier for the session
	AccAddr       string `gorm:"column:acc_addr;not null"`             // Account address, cannot be null
	NodeAddr      string `gorm:"column:node_addr;not null"`            // Address of the node associated with the session
	DownloadBytes string `gorm:"column:download_bytes;not null"`       // Download bytes represented as a string
	UploadBytes   string `gorm:"column:upload_bytes;not null"`         // Upload bytes represented as a string
	MaxBytes      string `gorm:"column:max_bytes;not null"`            // Maximum bytes represented as a string
	Duration      int64  `gorm:"column:duration;not null"`             // Duration of the session in nanoseconds
	MaxDuration   int64  `gorm:"column:max_duration;not null"`         // Maximum allowed duration for the session in nanoseconds
	Signature     string `gorm:"column:signature;not null"`            // Signature associated with the session
	PeerKey       string `gorm:"column:peer_key;not null;uniqueIndex"` // Unique key for the peer, indexed and cannot be null
	ServiceType   string `gorm:"column:service_type;not null"`         // Type of service for the session
}

// NewSession creates and returns a new instance of the Session struct with default values.
func NewSession() *Session {
	return &Session{}
}

// WithID sets the ID field and returns the updated Session instance.
func (s *Session) WithID(v uint64) *Session {
	s.ID = v
	return s
}

// WithAccAddr sets the AccAddr field and returns the updated Session instance.
func (s *Session) WithAccAddr(v cosmossdk.AccAddress) *Session {
	s.AccAddr = v.String()
	return s
}

// WithNodeAddr sets the NodeAddr field and returns the updated Session instance.
func (s *Session) WithNodeAddr(v base.NodeAddress) *Session {
	s.NodeAddr = v.String()
	return s
}

// WithDownloadBytes sets the DownloadBytes field from cosmossdkmath.Int and returns the updated Session instance.
func (s *Session) WithDownloadBytes(v cosmossdkmath.Int) *Session {
	s.DownloadBytes = v.String()
	return s
}

// WithUploadBytes sets the UploadBytes field from cosmossdkmath.Int and returns the updated Session instance.
func (s *Session) WithUploadBytes(v cosmossdkmath.Int) *Session {
	s.UploadBytes = v.String()
	return s
}

// WithMaxBytes sets the MaxBytes field from cosmossdkmath.Int and returns the updated Session instance.
func (s *Session) WithMaxBytes(v cosmossdkmath.Int) *Session {
	s.MaxBytes = v.String()
	return s
}

// WithDuration sets the Duration field from time.Duration and returns the updated Session instance.
func (s *Session) WithDuration(v time.Duration) *Session {
	s.Duration = v.Nanoseconds()
	return s
}

// WithMaxDuration sets the MaxDuration field from time.Duration and returns the updated Session instance.
func (s *Session) WithMaxDuration(v time.Duration) *Session {
	s.MaxDuration = v.Nanoseconds()
	return s
}

// WithSignature sets the Signature field and returns the updated Session instance.
func (s *Session) WithSignature(v []byte) *Session {
	s.Signature = base64.StdEncoding.EncodeToString(v)
	return s
}

// WithPeerKey sets the PeerKey field and returns the updated Session instance.
func (s *Session) WithPeerKey(v string) *Session {
	s.PeerKey = v
	return s
}

// WithServiceType sets the ServiceType field and returns the updated Session instance.
func (s *Session) WithServiceType(v sdk.ServiceType) *Session {
	s.ServiceType = v.String()
	return s
}

// GetID returns the ID field.
func (s *Session) GetID() uint64 {
	return s.ID
}

// GetAccAddr returns the AccAddr field as cosmossdk.AccAddress.
func (s *Session) GetAccAddr() cosmossdk.AccAddress {
	addr, err := cosmossdk.AccAddressFromBech32(s.AccAddr)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetNodeAddr returns the NodeAddr field as base.NodeAddress.
func (s *Session) GetNodeAddr() base.NodeAddress {
	addr, err := base.NodeAddressFromBech32(s.NodeAddr)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetBytes returns the total number of bytes (download + upload) as cosmossdkmath.Int.
func (s *Session) GetBytes() cosmossdkmath.Int {
	downloadBytes := s.GetDownloadBytes()
	uploadBytes := s.GetUploadBytes()

	return downloadBytes.Add(uploadBytes)
}

// GetDownloadBytes returns the DownloadBytes field as cosmossdkmath.Int.
func (s *Session) GetDownloadBytes() cosmossdkmath.Int {
	v, ok := cosmossdkmath.NewIntFromString(s.DownloadBytes)
	if !ok {
		panic(fmt.Errorf("invalid download_bytes: %s", s.DownloadBytes))
	}

	return v
}

// GetUploadBytes returns the UploadBytes field as cosmossdkmath.Int.
func (s *Session) GetUploadBytes() cosmossdkmath.Int {
	v, ok := cosmossdkmath.NewIntFromString(s.UploadBytes)
	if !ok {
		panic(fmt.Errorf("invalid upload_bytes: %s", s.UploadBytes))
	}

	return v
}

// GetMaxBytes returns the MaxBytes field as cosmossdkmath.Int.
func (s *Session) GetMaxBytes() cosmossdkmath.Int {
	v, ok := cosmossdkmath.NewIntFromString(s.MaxBytes)
	if !ok {
		panic(fmt.Errorf("invalid max_bytes: %s", s.MaxBytes))
	}

	return v
}

// GetDuration returns the Duration field as time.Duration.
func (s *Session) GetDuration() time.Duration {
	return time.Duration(s.Duration)
}

// GetMaxDuration returns the MaxDuration field as time.Duration.
func (s *Session) GetMaxDuration() time.Duration {
	return time.Duration(s.MaxDuration)
}

// GetSignature returns the Signature field as a byte slice.
func (s *Session) GetSignature() []byte {
	buf, err := base64.StdEncoding.DecodeString(s.Signature)
	if err != nil {
		panic(err)
	}

	return buf
}

// GetPeerKey returns the PeerKey field.
func (s *Session) GetPeerKey() string {
	return s.PeerKey
}

// GetServiceType returns the ServiceType field as sdk.ServiceType.
func (s *Session) GetServiceType() sdk.ServiceType {
	return sdk.ServiceTypeFromString(s.ServiceType)
}

// BeforeUpdate is a GORM hook that updates the Duration field if relevant fields change.
func (s *Session) BeforeUpdate(db *gorm.DB) (err error) {
	if s.ID == 0 {
		return nil
	}

	if db.Statement.Changed("download_bytes", "upload_bytes") {
		duration := time.Since(s.CreatedAt).Nanoseconds()
		db.Statement.SetColumn("duration", duration)
	}

	return nil
}

// MsgUpdateSessionRequest creates a MsgUpdateSessionRequest for the session.
func (s *Session) MsgUpdateSessionRequest() *sessiontypes.MsgUpdateSessionRequest {
	return sessiontypes.NewMsgUpdateSessionRequest(
		s.GetNodeAddr(),
		s.GetID(),
		s.GetDownloadBytes(),
		s.GetUploadBytes(),
		s.GetDuration(),
		s.GetSignature(),
	)
}
