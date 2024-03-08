package controller

import (
	m "UTS_PBP/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM rooms"
	room_name := r.URL.Query()["room_name"]
	id_game := r.URL.Query()["id_game"]

	if room_name != nil {
		fmt.Println(room_name[0])
		query += " WHERE name='" + room_name[0] + "'"

	}
	if id_game != nil {
		if room_name[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.RoomName, &room.ID_Game); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT r.id,r.room_name,p.id,p.id_account,p.id_room,a.id,a.username FROM rooms r
	JOIN participants p ON r.id = p.id_room
	JOIN accounts a ON p.id_account = a.id;`
	DetailRoomRow, err := db.Query(query)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w)
		return
	}

	var room m.RoomDetail
	var rooms []m.RoomDetail
	var detailRoom m.DetailRooms

	for DetailRoomRow.Next() {
		if err := DetailRoomRow.Scan(
			&detailRoom.ID, &detailRoom.RoomName, &detailRoom.Participant.ID, &detailRoom.Participant.ID_Account, &detailRoom.Participant.ID_Room, &detailRoom.Account.ID,
			&detailRoom.Account.Username); err != nil {
			print(err.Error())
			sendErrorResponse(w)
			return
		}

		room.Room = detailRoom
		rooms = append(rooms, room)
	}
	var response m.RoomDetailResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	idRoom, _ := strconv.Atoi(r.Form.Get("idRoom"))
	idAccount, _ := strconv.Atoi(r.Form.Get("idAccount"))

	var maxPlayer int
	err = db.QueryRow("SELECT g.max_player FROM rooms r JOIN games g ON r.id_game = g.id WHERE r.id = ?", idRoom).Scan(&maxPlayer)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch room information", http.StatusInternalServerError)
		return
	}

	var currentParticipantsCount int
	err = db.QueryRow("SELECT COUNT(*) FROM participants WHERE id_room = ?", idRoom).Scan(&currentParticipantsCount)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to count current participants", http.StatusInternalServerError)
		return
	}
	if currentParticipantsCount >= maxPlayer {
		http.Error(w, "Room is full, cannot join", http.StatusBadRequest)
		return
	}

	_, errQuery := db.Exec("INSERT INTO participants(id_room,id_account) values (?,?)",
		idRoom,
		idAccount,
	)

	var response m.RoomResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Insert Failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	participantID, _ := strconv.Atoi(r.Form.Get("id"))

	_, errQuery := db.Exec("DELETE FROM participants WHERE id=?",
		participantID,
	)

	if errQuery == nil {
		sendSuccessResponse(w)
	} else {
		sendErrorResponse(w)
	}
}

func sendErrorResponse(w http.ResponseWriter) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func sendSuccessResponse(w http.ResponseWriter) {
	var response m.SuccessResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
