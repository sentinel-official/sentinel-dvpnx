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

	AccAddr       string `gorm:"column:acc_addr;not null"`             // Account address, cannot be null
	DownloadBytes string `gorm:"column:download_bytes;not null"`       // Download bytes represented as a string
	Duration      int64  `gorm:"column:duration;not null"`             // Duration of the session in nanoseconds
	ID            uint64 `gorm:"column:id;not null;primaryKey"`        // Unique identifier for the session
	MaxBytes      string `gorm:"column:max_bytes;not null"`            // Maximum bytes represented as a string
	MaxDuration   int64  `gorm:"column:max_duration;not null"`         // Maximum allowed duration for the session in nanoseconds
	NodeAddr      string `gorm:"column:node_addr;not null"`            // Address of the node associated with the session
	PeerKey       string `gorm:"column:peer_key;not null;uniqueIndex"` // Unique key for the peer, indexed and cannot be null
	ServiceType   string `gorm:"column:service_type;not null"`         // Type of service for the session
	Signature     string `gorm:"column:signature;not null"`            // Signature associated with the session
	UploadBytes   string `gorm:"column:upload_bytes;not null"`         // Upload bytes represented as a string
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

// WithDownloadBytes sets the DownloadBytes field from math.Int and returns the updated Session instance.
func (s *Session) WithDownloadBytes(v math.Int) *Session {
	s.DownloadBytes = v.String()
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

// WithPeerKey sets the PeerKey field and returns the updated Session instance.
func (s *Session) WithPeerKey(v string) *Session {
	s.PeerKey = v
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

// WithUploadBytes sets the UploadBytes field from math.Int and returns the updated Session instance.
func (s *Session) WithUploadBytes(v math.Int) *Session {
	s.UploadBytes = v.String()
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

// GetBytes returns the total number of bytes (download + upload) as math.Int.
func (s *Session) GetBytes() math.Int {
	downloadBytes := s.GetDownloadBytes()
	uploadBytes := s.GetUploadBytes()

	return downloadBytes.Add(uploadBytes)
}

// GetDownloadBytes returns the DownloadBytes field as math.Int.
func (s *Session) GetDownloadBytes() math.Int {
	v, ok := math.NewIntFromString(s.DownloadBytes)
	if !ok {
		panic(fmt.Errorf("invalid download_bytes %s", s.DownloadBytes))
	}

	return v
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

// GetPeerKey returns the PeerKey field.
func (s *Session) GetPeerKey() string {
	return s.PeerKey
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

// GetUploadBytes returns the UploadBytes field as math.Int.
func (s *Session) GetUploadBytes() math.Int {
	v, ok := math.NewIntFromString(s.UploadBytes)
	if !ok {
		panic(fmt.Errorf("invalid upload_bytes %s", s.UploadBytes))
	}

	return v
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
func (s *Session) MsgUpdateSessionRequest() *v3.MsgUpdateSessionRequest {
	return v3.NewMsgUpdateSessionRequest(
		s.GetNodeAddr(),
		s.GetID(),
		s.GetDownloadBytes(),
		s.GetUploadBytes(),
		s.GetDuration(),
		s.GetSignature(),
	)
}
