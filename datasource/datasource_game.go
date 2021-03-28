package datasource

import (
	"database/sql"
	"minesweeper-API/minesweeper-service/model"
)

func (ds *Datasource) FindGame(id int) (*model.Game, error) {
	var game model.Game
	switch err := ds.db.Get(&game, `SELECT * FROM minesweeper.games WHERE game_id = $1`, id); err {
	case nil, sql.ErrNoRows:
		return &game, nil
	default:
		return nil, err
	}
}

func (ds *Datasource) SaveGame(g *model.Game) (int, error) {
	res, err := ds.db.NamedQuery(
		`INSERT INTO minesweeper.games (state, columns, rows, mine_amount, flag_amount, board)
 		VALUES (:state, :columns, :rows, :mine_amount, :flag_amount, :board) returning game_id`,
		&g)

	if err != nil {
		return 0, err
	}

	res.Next()
	var id int
	err = res.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ds *Datasource) UpdateGame(g *model.Game) error {
	_, err := ds.db.NamedQuery(
		`UPDATE  minesweeper.games  SET state = :state, 
                               columns = :columns, rows = :rows, mine_amount = :mine_amount, flag_amount = :flag_amount, 
                               board = :board WHERE game_id = :game_id`,
		&g)

	return err
}