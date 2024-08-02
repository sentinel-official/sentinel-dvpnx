package models

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
)

type Session struct {
	CreatedAt time.Time `gorm:"autoCreateTime"` // Timestamp when the record was created
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Timestamp when the record was last updated

	ID            uint64 `gorm:"primaryKey"`           // Unique identifier for the session
	AccAddress    string `gorm:"not null"`             // Account address, cannot be null
	DownloadBytes string `gorm:"not null"`             // Download bytes represented as a string
	UploadBytes   string `gorm:"not null"`             // Upload bytes represented as a string
	MaxBytes      string `gorm:"not null"`             // Maximum bytes represented as a string
	Duration      int64  `gorm:"not null"`             // Duration of the session in seconds
	MaxDuration   int64  `gorm:"not null"`             // Maximum allowed duration for the session in seconds
	PeerKey       string `gorm:"not null;uniqueIndex"` // Unique key for the peer, indexed and cannot be null
}

// NewSession creates and returns a new instance of the Session struct with default values.
func NewSession() *Session {
	return &Session{}
}

// WithID sets the ID field and returns the updated Session instance
func (s *Session) WithID(v uint64) *Session {
	s.ID = v
	return s
}

// WithAccAddress sets the AccAddress field and returns the updated Session instance
func (s *Session) WithAccAddress(v string) *Session {
	s.AccAddress = v
	return s
}

// WithDownloadBytes sets the DownloadBytes field from sdkmath.Int and returns the updated Session instance
func (s *Session) WithDownloadBytes(v sdkmath.Int) *Session {
	s.DownloadBytes = v.String()
	return s
}

// WithUploadBytes sets the UploadBytes field from sdkmath.Int and returns the updated Session instance
func (s *Session) WithUploadBytes(v sdkmath.Int) *Session {
	s.UploadBytes = v.String()
	return s
}

// WithMaxBytes sets the MaxBytes field from sdkmath.Int and returns the updated Session instance
func (s *Session) WithMaxBytes(v sdkmath.Int) *Session {
	s.MaxBytes = v.String()
	return s
}

// WithDuration sets the Duration field from time.Duration and returns the updated Session instance
func (s *Session) WithDuration(v time.Duration) *Session {
	s.Duration = v.Nanoseconds()
	return s
}

// WithMaxDuration sets the MaxDuration field from time.Duration and returns the updated Session instance
func (s *Session) WithMaxDuration(v time.Duration) *Session {
	s.MaxDuration = v.Nanoseconds()
	return s
}

// WithPeerKey sets the PeerKey field and returns the updated Session instance
func (s *Session) WithPeerKey(v string) *Session {
	s.PeerKey = v
	return s
}

// GetID returns the ID field
func (s *Session) GetID() uint64 {
	return s.ID
}

// GetAccAddress returns the AccAddress field
func (s *Session) GetAccAddress() string {
	return s.AccAddress
}

// GetDownloadBytes returns the DownloadBytes field as sdkmath.Int
func (s *Session) GetDownloadBytes() sdkmath.Int {
	v, ok := sdkmath.NewIntFromString(s.DownloadBytes)
	if !ok {
		panic(fmt.Errorf("invalid download_bytes: %s", s.DownloadBytes))
	}

	return v
}

// GetUploadBytes returns the UploadBytes field as sdkmath.Int
func (s *Session) GetUploadBytes() sdkmath.Int {
	v, ok := sdkmath.NewIntFromString(s.UploadBytes)
	if !ok {
		panic(fmt.Errorf("invalid upload_bytes: %s", s.UploadBytes))
	}

	return v
}

// GetMaxBytes returns the MaxBytes field as sdkmath.Int
func (s *Session) GetMaxBytes() sdkmath.Int {
	v, ok := sdkmath.NewIntFromString(s.MaxBytes)
	if !ok {
		panic(fmt.Errorf("invalid max_bytes: %s", s.MaxBytes))
	}

	return v
}

// GetDuration returns the Duration field as time.Duration
func (s *Session) GetDuration() time.Duration {
	return time.Duration(s.Duration)
}

// GetMaxDuration returns the MaxDuration field as time.Duration
func (s *Session) GetMaxDuration() time.Duration {
	return time.Duration(s.MaxDuration)
}
