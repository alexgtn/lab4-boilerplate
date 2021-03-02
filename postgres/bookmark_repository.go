package main

import (
	"context"
	"database/sql"
	"fmt"
)

// BookmarkRepository stores bookmarks
type BookmarkRepository struct {
	db *sql.DB
}

// NewBookmarkRepository builds new bookmark repository using a db connection
func NewBookmarkRepository(db *sql.DB) *BookmarkRepository {
	return &BookmarkRepository{
		db: db,
	}
}

// Create creates bookmark
func (r *BookmarkRepository) Create(bookmark *Bookmark) (*Bookmark, error) {
	// TODO
}

// GetAll gets all bookmarks
func (r *BookmarkRepository) GetAll() ([]*Bookmark, error) {
	// TODO
}
