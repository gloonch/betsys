package mysql

import (
	"betsys/entity"
	"fmt"
)

// We have register for user to support different games
func (d MySQLDB) Register(s entity.Server) (entity.Server, error) {
	response, err := d.db.Exec(`insert into serverservice(name) values(?,)`, s.Name)
	if err != nil {
		return entity.Server{}, fmt.Errorf("cannot register serverservice: %w", err)
	}

	id, _ := response.LastInsertId()
	s.ID = uint(id)
	return s, nil
}

func (d MySQLDB) FindServerByName(name string) (entity.Server, error) {
	var server entity.Server

	row := d.db.QueryRow(`select * from serverservice where name = ?`, name)
	err := row.Scan(&server.ID, &server.Name)

	return server, err
}

func (d MySQLDB) IsServerNameTaken(name string) (bool, error) {
	var server entity.Server

	// TODO: you could also select count() to see if the count is more than zero, then it's taken

	row := d.db.QueryRow(`select * from serverservice where name = ?`, name)
	err := row.Scan(&server.ID)

	isTaken := server.ID != 0
	return isTaken, err
}
