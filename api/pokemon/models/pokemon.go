package models

import (
	"time"

	"github.com/go-pg/pg/v10"
)

// Pokemon Pokemon struct model
type Pokemon struct {
	ID             int64     `json:"id" pg:",pk"`
	Name           *string   `json:"name" binding:"required"`
	Type           *string   `json:"type" binding:"required"`
	Speed          *int      `json:"speed" binding:"required"`
	Attack         *int      `json:"attack" binding:"required"`
	AttackSpecial  *int      `json:"attack_special" binding:"required"`
	Defense        *int      `json:"defense" binding:"required"`
	DefenseSpecial *int      `json:"defense_special" binding:"required"`
	HealthPoint    *int      `json:"health_point" binding:"required"`
	Shiny          *bool     `json:"shiny" binding:"required"`
	CreatedAt      time.Time `json:"created_at" pg:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" pg:"default:CURRENT_TIMESTAMP"`
}

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

func (pokemons *Pokemons) List(conn *pg.DB) error {
	return conn.Model(pokemons).Select()
}
