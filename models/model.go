package models

type Rooms struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	ID_Game  int    `json:"id_game"`
}
type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Rooms  `json:"data"`
}
type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}
type Participants struct {
	ID         int `json:"id"`
	ID_Room    int `json:"id_room"`
	ID_Account int `json:"id_account"`
}
type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
type DetailRooms struct {
	ID          int          `json:"id"`
	RoomName    string       `json:"room_name"`
	Participant Participants `json:"participant"`
	Account     Account      `json:"account"`
}
type RoomDetails struct {
	Room []DetailRooms `json:"rooms"`
}
type RoomDetail struct {
	Room DetailRooms `json:"room"`
}
type RoomDetailResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    []RoomDetail `json:"data"`
}
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
