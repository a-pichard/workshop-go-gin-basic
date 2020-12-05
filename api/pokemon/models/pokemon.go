package models

import (
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
)

// Pokemon Pokemon struct model
type Pokemon struct {
	ID             int64     `json:"id" pg:",pk"`
	Name           *string   `json:"name" validate:"nonnil"`
	Type           *string   `json:"type" validate:"nonnil"`
	Speed          *int      `json:"speed" validate:"nonnil"`
	Attack         *int      `json:"attack" validate:"nonnil"`
	AttackSpecial  *int      `json:"attack_special" validate:"nonnil"`
	Defense        *int      `json:"defense" validate:"nonnil"`
	DefenseSpecial *int      `json:"defense_special" validate:"nonnil"`
	HealthPoint    *int      `json:"health_point" validate:"nonnil"`
	Shiny          *bool     `json:"shiny" validate:"nonnil"`
	CreatedAt      time.Time `json:"created_at" pg:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" pg:"default:CURRENT_TIMESTAMP"`
}

// Pokemons array of Pokemon type
type Pokemons []Pokemon

// InsertMe insert new pokemon in database
func (pokemon *Pokemon) InsertMe(conn *pg.DB) error {
	_, err := conn.Model(pokemon).Insert()
	return err
}

// GetByPk Get pokemon by it's primary key
func (pokemon *Pokemon) GetByPk(conn *pg.DB) error {
	return conn.Model(pokemon).WherePK().Select()
}

// List list all pokemons
func (pokemons *Pokemons) List(conn *pg.DB) error {
	return conn.Model(pokemons).Select()
}

// UpdateMe update a pokemon
func (pokemon *Pokemon) UpdateMe(conn *pg.DB) error {
	_, err := conn.Model(pokemon).WherePK().UpdateNotZero()
	return err
}

// DeleteByPk delete a pokemon
func (pokemon *Pokemon) DeleteByPk(conn *pg.DB) error {
	result, err := conn.Model(pokemon).WherePK().Delete()
	if result.RowsAffected() == 0 {
		return errors.New("Element not found")
	}
	return err
}
