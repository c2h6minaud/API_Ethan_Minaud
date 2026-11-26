package store

import (
	"API/model"
	"errors"
	"sync"
	"encoding/json"
	"os"
)

var ErrArtistNotFound = errors.New("artist not found")

type Store struct {
	artists []model.Artist
	mu      sync.RWMutex
	nextID  int
	path    string
}

func NewStore(path string) *Store {
	s := &Store{
		artists: []model.Artist{},
		nextID:  1,
		path:    path,
	}
	s.load()
	return s
}

func (s *Store) load() {
	maxID := 0
	data, err := os.ReadFile(s.path)
	if err != nil {
		return
	}
	if len(data) == 0 {
		return
	}
	err = json.Unmarshal(data, &s.artists)
	if  err != nil {
		s.artists = []model.Artist{}
		return
	}
	if s.artists == nil {
		s.artists = []model.Artist{}
	}
	for _, artist := range s.artists {
		if artist.ID > maxID {
			maxID = artist.ID
		}
	}
	s.nextID = maxID + 1
}

func (s *Store) save() {
	data, err := json.MarshalIndent(s.artists, "", "  ")
	if err != nil {
		return
	}
	err = os.WriteFile(s.path, data, 0644)
	if err != nil {
		return
	}
}

func (s *Store) Artists() []model.Artist {
	s.mu.RLock()
	defer s.mu.RUnlock()
	artists := make([]model.Artist, len(s.artists))
	copy(artists, s.artists)
	return artists
}

func (s *Store) GetByID(id int) (model.Artist, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, artist := range s.artists {
		if artist.ID == id {
			return artist, nil
		}
	}
	return model.Artist{}, ErrArtistNotFound
}

func (s *Store) Create(artist model.Artist) model.Artist {
	s.mu.Lock()
	defer s.mu.Unlock()
	artist.ID = s.nextID
	s.nextID++
	s.artists = append(s.artists, artist)
	s.save()
	return artist
}

func (s *Store) Update(id int, updatedArtist model.Artist) (model.Artist, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, artist := range s.artists {
		if artist.ID == id {
			updatedArtist.ID = id
			s.artists[i] = updatedArtist
			s.save()
			return updatedArtist, nil
		}
	}
	return model.Artist{}, ErrArtistNotFound
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, artist := range s.artists {
		if artist.ID == id {
			s.artists = append(s.artists[:i], s.artists[i+1:]...)
			s.save()
			return nil
		}
	}
	return ErrArtistNotFound
}
