package repository

import (
	"database/sql"
	"fmt"
	"travelnest/internal/model"
)

type RoomRepository struct {
	connection *sql.DB
}

func NewRoomRepository(connection *sql.DB) RoomRepository {
	return RoomRepository{
		connection: connection,
	}
}

func (rr * RoomRepository) GetRooms() ([]model.Room, error) {
	query := "SELECT * FROM rooms"
	rows, err := rr.connection.Query(query)
	
	if err != nil {
		fmt.Println(err)
		return []model.Room{}, err
	}
	
	var roomList []model.Room

	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.ID, &room.Name, &room.Description, &room.Price, &room.Status)
		if err != nil {
			fmt.Println(err)
			return []model.Room{}, err
		}
		roomList = append(roomList, room)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return []model.Room{}, err
	}

	return roomList, nil
}