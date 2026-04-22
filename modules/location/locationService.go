package location

import (
	"mini-inventory/config"
	"time"
)

type LocationService struct{}

func (s *LocationService) GetAllLocations() ([]Location, error) {
	var locations []Location
	query := `SELECT id, title, description, created_by, created_at, updated_at FROM locations`
	err := config.DB.Select(&locations, query)
	return locations, err
}

func (s *LocationService) GetLocationByID(id string) (Location, error) {
	var l Location
	query := `SELECT id, title, description, created_by, created_at, updated_at FROM locations WHERE id = $1`
	err := config.DB.Get(&l, query, id)
	return l, err
}

func (s *LocationService) CreateLocation(l *Location) error {
	query := `INSERT INTO locations (title, description, created_by, created_at, updated_at) 
			  VALUES (:title, :description, :created_by, :created_at, :updated_at) RETURNING id`

	now := time.Now()
	l.CreatedAt = now
	l.UpdatedAt = now

	nstmt, err := config.DB.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	err = nstmt.Get(&l.ID, l)
	return err
}

func (s *LocationService) UpdateLocation(l *Location) error {
	query := `UPDATE locations SET title = :title, description = :description, updated_at = :updated_at WHERE id = :id`

	l.UpdatedAt = time.Now()
	_, err := config.DB.NamedExec(query, l)
	return err
}

func (s *LocationService) DeleteLocation(id string) error {
	query := `DELETE FROM locations WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	return err
}
