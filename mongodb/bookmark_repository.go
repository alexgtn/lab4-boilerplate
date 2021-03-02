package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// BookmarkRepository stores bookmarks
type BookmarkRepository struct {
	db *mongo.Client
}

// NewBookmarkRepository builds new bookmark repository using a db connection
func NewBookmarkRepository(db *mongo.Client) *BookmarkRepository {
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
