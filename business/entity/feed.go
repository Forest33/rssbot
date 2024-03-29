// Package entity provides entities for business logic.
package entity

import (
	"time"
)

// Feed is feed
type Feed struct {
	ID           string
	Title        string
	SiteURL      string
	FeedURL      string
	LastItemHash string
	ErrorCount   int
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

// RSSFeedItem is feed item
type RSSFeedItem struct {
	GUID        string
	Title       string
	Description string
	Content     string
	Link        string
	Updated     string
	Published   string
	Hash        string
}

// FeedFilter is filter for getting feeds
type FeedFilter struct {
	ID []string
}

// FeedItem is feed
type FeedItem struct {
	ID        string
	FeedID    string
	ItemHash  string
	CreatedAt time.Time
}
