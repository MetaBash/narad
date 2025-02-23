package schema

// # Email Request Model
type EmailRequest struct {
	CamID    string `bson:"cam_id,omitempty" json:"cam_id,omitempty" validate:"required"`
	Image    string `bson:"image,omitempty" json:"image,omitempty" validate:"required"`
	IsPerson string `bson:"is_person,omitempty" json:"is_person,omitempty", validate:"required"`
}

// # Email Response Model
type EmailResponse struct {
	CamID        string `bson:"cam_id,omitempty" json:"cam_id,omitempty"`
	CamName      string `bson:"cam_name,omitempty" json:"cam_name,omitempty"`
	OrgID        string `bson:"org_id,omitempty" json:"org_id,omitempty"`
	OrgName      string `bson:"org_name,omitempty" json:"org_name,omitempty"`
	IPAddress    string `bson:"ip_address,omitempty" json:"ip_address,omitempty"`
	BuildingNo   string `bson:"building_no,omitempty" json:"building_no,omitempty"`
	BuildingName string `bson:"building_name,omitempty" json:"building_name,omitempty"`
	FloorNo      string `bson:"floor_no,omitempty" json:"floor_no,omitempty"`
	FloorName    string `bson:"floor_name,omitempty" json:"floor_name,omitempty"`
	RoomNo       string `bson:"room_no,omitempty" json:"room_no,omitempty"`
	RoomName     string `bson:"room_name,omitempty" json:"room_name,omitempty"`
	AreaName     string `bson:"area_name,omitempty" json:"area_name,omitempty"`
	Latitude     string `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude    string `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Address      string `bson:"address,omitempty" json:"address,omitempty"`
	City         string `bson:"city,omitempty" json:"city,omitempty"`
	State        string `bson:"state,omitempty" json:"state,omitempty"`
	Country      string `bson:"country,omitempty" json:"country,omitempty"`
	Pincode      string `bson:"pincode,omitempty" json:"pincode,omitempty"`
	MapLocation  string `bson:"map_location,omitempty" json:"map_location,omitempty"`
	AlertMsg     string `bson:"alert_msg,omitempty" json:"alert_msg,omitempty"`
	AlertClass   string `bson:"alert_class,omitempty" json:"alert_class,omitempty"`
}

// # Live Card Model
type LiveCard struct {
	OrgID        string `bson:"org_id,omitempty" json:"org_id,omitempty"`
	OrgName      string `bson:"org_name,omitempty" json:"org_name,omitempty"`
	BuildingNo   string `bson:"building_no,omitempty" json:"building_no,omitempty"`
	BuildingName string `bson:"building_name,omitempty" json:"building_name,omitempty"`
	FloorNo      string `bson:"floor_no,omitempty" json:"floor_no,omitempty"`
	FloorName    string `bson:"floor_name,omitempty" json:"floor_name,omitempty"`
	RoomNo       string `bson:"room_no,omitempty" json:"room_no,omitempty"`
	RoomName     string `bson:"room_name,omitempty" json:"room_name,omitempty"`
	AreaName     string `bson:"area_name,omitempty" json:"area_name,omitempty"`
	Address      string `bson:"address,omitempty" json:"address,omitempty"`
	City         string `bson:"city,omitempty" json:"city,omitempty"`
	State        string `bson:"state,omitempty" json:"state,omitempty"`
	Country      string `bson:"country,omitempty" json:"country,omitempty"`
	Pincode      string `bson:"pincode,omitempty" json:"pincode,omitempty"`
	MapLocation  string `bson:"map_location,omitempty" json:"map_location,omitempty"`
	Image        string `bson:"image,omitempty" json:"image,omitempty"`
	IsPerson     bool   `bson:"is_person,omitempty" json:"is_person,omitempty"`
	IsActive     bool   `bson:"is_active,omitempty" json:"is_active,omitempty"`
}

// # History Card Model
type HistoryCard struct {
	OrgID        string `bson:"org_id,omitempty" json:"org_id,omitempty"`
	OrgName      string `bson:"org_name,omitempty" json:"org_name,omitempty"`
	BuildingNo   string `bson:"building_no,omitempty" json:"building_no,omitempty"`
	BuildingName string `bson:"building_name,omitempty" json:"building_name,omitempty"`
	FloorNo      string `bson:"floor_no,omitempty" json:"floor_no,omitempty"`
	FloorName    string `bson:"floor_name,omitempty" json:"floor_name,omitempty"`
	RoomNo       string `bson:"room_no,omitempty" json:"room_no,omitempty"`
	RoomName     string `bson:"room_name,omitempty" json:"room_name,omitempty"`
	AreaName     string `bson:"area_name,omitempty" json:"area_name,omitempty"`
	Address      string `bson:"address,omitempty" json:"address,omitempty"`
	City         string `bson:"city,omitempty" json:"city,omitempty"`
	State        string `bson:"state,omitempty" json:"state,omitempty"`
	Country      string `bson:"country,omitempty" json:"country,omitempty"`
	Pincode      string `bson:"pincode,omitempty" json:"pincode,omitempty"`
	MapLocation  string `bson:"map_location,omitempty" json:"map_location,omitempty"`
	Cause        string `bson:"cause,omitempty" json:"cause,omitempty"`
	Image        string `bson:"image,omitempty" json:"image,omitempty"`
	IsPerson     bool   `bson:"is_person,omitempty" json:"is_person,omitempty"`
	IsActive     bool   `bson:"is_active,omitempty" json:"is_active,omitempty"`
}
