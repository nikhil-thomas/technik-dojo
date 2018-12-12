package main

import (
	"io"
)

// FileSystemPlayerStore defines a FileSystemPlayerStore
type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// GetLeague returns league
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
