package api

import (
	"context"
	"encoding/json"
	"errors"
	"narad/app"
	"narad/schema"
	"narad/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// # Send Notification
func SendNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var err error

	var errMsg string

	ctx := context.Background()

	var data *schema.EmailRequest

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&data)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	camIDStr := data.CamID

	if camIDStr == "" {
		errMsg = "error: 'cam_id' query parameter is required"
		err = errors.New(errMsg)
		utils.SendError(w, err)
		return
	}

	camID, err := primitive.ObjectIDFromHex(camIDStr)
	if err != nil {
		errMsg = "error: Invalid cam_id format"
		err = errors.New(errMsg)
		utils.SendError(w, err)
		return
	}

	person := data.IsPerson
	if person == "" {
		errMsg = "error: 'is_person' query parameter is required"
		err = errors.New(errMsg)
		utils.SendError(w, err)
		return
	} else if person != "true" && person != "false" {
		errMsg = "error: Invalid is_person value (true/false)"
		err = errors.New(errMsg)
		utils.SendError(w, err)
		return
	}

	res, err := app.SendNotification(data, camID, ctx)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.SendSuccess(w, res)
}

// # Get Live Card
func GetLiveCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var err error

	ctx := context.Background()

	res, err := app.GetLiveCard(ctx)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.SendPayload(w, res)
}

// # Get History Card
func GetHistoryCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var err error

	ctx := context.Background()

	res, err := app.GetHistoryCard(ctx)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.SendPayload(w, res)
}

// # Mark Organisation Inactive
func MarkOrgInActive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var err error

	var errMsg string

	params := mux.Vars(r)

	orgID := params["org_id"]

	if orgID == "" {
		errMsg = "error: 'org_id' query parameter is required"
		err = errors.New(errMsg)
		utils.SendError(w, err)
		return
	}

	ctx := context.Background()

	res, err := app.MarkOrgInActive(orgID, ctx)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.SendPayload(w, res)
}

// # Upload Image
func UploadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var err error

	var errMsg string

	var image *schema.Image

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&image)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	base64ImgStr := image.Image

	folder := "public"

	res, err := utils.SaveBase64Image(base64ImgStr, folder)
	if err != nil {
		errMsg = "error: failed to save image: " + err.Error()
		err = errors.New(errMsg)
		utils.SendError(w, err)
		return
	}

	utils.SendSuccess(w, res)
}

// # Send Email
func SendEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "https://cresaclub.vercel.app")

	var err error

	var data schema.Email

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&data)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	res, err := app.SendEmail(&data)
	if err != nil {
		utils.SendError(w, err)
		return
	}

	utils.SendSuccess(w, res)
}
