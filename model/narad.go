package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// # Narad Model
type Narad struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	OrgID        string             `bson:"org_id,omitempty" json:"org_id,omitempty"`
	OrgName      string             `bson:"org_name,omitempty" json:"org_name,omitempty"`
	Users        []User             `bson:"users,omitempty" json:"users,omitempty"`
	IPAddress    string             `bson:"ip_address,omitempty" json:"ip_address,omitempty"`
	BuildingNo   string             `bson:"building_no,omitempty" json:"building_no,omitempty"`
	BuildingName string             `bson:"building_name,omitempty" json:"building_name,omitempty"`
	FloorNo      string             `bson:"floor_no,omitempty" json:"floor_no,omitempty"`
	FloorName    string             `bson:"floor_name,omitempty" json:"floor_name,omitempty"`
	RoomNo       string             `bson:"room_no,omitempty" json:"room_no,omitempty"`
	RoomName     string             `bson:"room_name,omitempty" json:"room_name,omitempty"`
	AreaName     string             `bson:"area_name,omitempty" json:"area_name,omitempty"`
	Coordinates  Coordinates        `bson:"coordinates,omitempty" json:"coordinates,omitempty"`
	Address      string             `bson:"address,omitempty" json:"address,omitempty"`
	City         string             `bson:"city,omitempty" json:"city,omitempty"`
	State        string             `bson:"state,omitempty" json:"state,omitempty"`
	Country      string             `bson:"country,omitempty" json:"country,omitempty"`
	Pincode      string             `bson:"pincode,omitempty" json:"pincode,omitempty"`
	MapLocation  string             `bson:"map_location,omitempty" json:"map_location,omitempty"`
	Image        string             `bson:"image,omitempty" json:"image,omitempty"`
	IsPerson     bool               `bson:"is_person,omitempty" json:"is_person,omitempty"`
	IsActive     bool               `bson:"is_active,omitempty" json:"is_active,omitempty"`
}

// # User Model
type User struct {
	ID    string `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string `bson:"name,omitempty" json:"name"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}

// # Organization Model
type Organization struct {
	ID    string `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	CamID string `bson:"cam_id,omitempty" json:"cam_id,omitempty"`
}

// # Camera Model
type Camera struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	IPAddress    string             `bson:"ip_address,omitempty" json:"ip_address,omitempty"`
	BuildingNo   string             `bson:"building_no,omitempty" json:"building_no,omitempty"`
	BuildingName string             `bson:"building_name,omitempty" json:"building_name,omitempty"`
	FloorNo      string             `bson:"floor_no,omitempty" json:"floor_no,omitempty"`
	FloorName    string             `bson:"floor_name,omitempty" json:"floor_name,omitempty"`
	RoomNo       string             `bson:"room_no,omitempty" json:"room_no,omitempty"`
	RoomName     string             `bson:"room_name,omitempty" json:"room_name,omitempty"`
	AreaName     string             `bson:"area_name,omitempty" json:"area_name,omitempty"`
	Coordinates  Coordinates        `bson:"coordinates,omitempty" json:"coordinates,omitempty"`
	Address      string             `bson:"address,omitempty" json:"address,omitempty"`
	City         string             `bson:"city,omitempty" json:"city,omitempty"`
	State        string             `bson:"state,omitempty" json:"state,omitempty"`
	Country      string             `bson:"country,omitempty" json:"country,omitempty"`
	Pincode      string             `bson:"pincode,omitempty" json:"pincode,omitempty"`
	MapLocation  string             `bson:"map_location,omitempty" json:"map_location,omitempty"`
}

// # Coordinates Model
type Coordinates struct {
	Latitude  string `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude string `bson:"longitude,omitempty" json:"longitude,omitempty"`
}
