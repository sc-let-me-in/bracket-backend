package models

// Note: all of these models map to a table in the database

// Song Represents a single spotify song.
type Song struct {
	SongID string // primary key for this table
	Title string
	AlbumID string
	PreviewClip string // don't know how to represent this yet
	ArtistID string
}

// Album maps to an album by an artist containing multiple songs.
type Album struct {
	AlbumID string // primary key
	Title string
	ArtistID string
	Year int // TODO: change this to a date-based structure or do some validation or something
}

// Artist represents a single artist as defined by spotify
type Artist struct {
	ArtistID string // primary key
	Name string
	// TODO: Artist picture, bio, etc. etc.
}