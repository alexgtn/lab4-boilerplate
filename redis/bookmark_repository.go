package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

// BookmarkRepository stores bookmarks
type BookmarkRepository struct {
	db *redis.Client
}

// NewBookmarkRepository builds new bookmark repository using a db connection
func NewBookmarkRepository(db *redis.Client) *BookmarkRepository {
	return &BookmarkRepository{
		db: db,
	}
}

// Create creates bookmark
func (r *BookmarkRepository) Create(bookmark *Bookmark) (*Bookmark, error) {
	// TODO
}

// GetAll gets all bookmarkss
func (r *BookmarkRepository) GetAll() ([]*Bookmark, error) {
	// TODO
}
