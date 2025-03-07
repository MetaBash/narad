package app

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"narad/db"
	"narad/env"
	"narad/model"
	"narad/schema"
	"narad/utils"
	"net/smtp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// # Send Notification
func SendNotification(data *schema.EmailRequest, camID primitive.ObjectID, ctx context.Context) (string, error) {
	var err error

	var errMsg string

	narad := env.GetEnv("NARAD", "narad")

	naradColl := db.GetCollectionByName(narad)

	var naradModel model.Narad

	filter := bson.M{"_id": camID}

	var personBool bool

	person := data.IsPerson

	image := data.Image

	imageURL, err := utils.UploadImage(image)
	if err != nil {
		return "", err
	}

	if person == "true" {
		personBool = true
	} else {
		personBool = false
	}

	set := bson.M{"image": imageURL, "is_person": personBool}

	update := bson.M{"$set": set}

	_, err = naradColl.UpdateOne(ctx, filter, update)
	if err != nil {
		errMsg = "error: failed to update narad document for camera ID: " + camID.Hex()
		err = errors.New(errMsg)
		return "", err
	}

	err = naradColl.FindOne(ctx, filter).Decode(&naradModel)
	if err != nil {
		errMsg = "error: camera not found for camera ID: " + camID.Hex()
		err = errors.New(errMsg)
		return "", err
	}

	if !naradModel.IsActive {
		if !naradModel.IsActive {
			err = fmt.Errorf("fire alert already sent or inactive for %s (Cam ID: %s)", naradModel.OrgName, camID.Hex())
			return "", err
		}

		errMsg = "info: fire alert already sent or fire is inactive for organisation '" + naradModel.OrgName + "' with camera ID: " + camID.Hex()
		err = errors.New(errMsg)
		return "", err
	}

	var alertMsg string

	if personBool {
		alertMsg = "High Alert: Human Detected"
	} else {
		alertMsg = "Alert: Fire Detected"
	}

	err = utils.SendEmail(naradModel, alertMsg)
	if err != nil {
		errMsg = "error: failed to send email alert notification: " + err.Error()
		err = errors.New(errMsg)
		return "", err
	}

	return "Email alert notification sent successfully!", nil
}

// # Get Live Card
func GetLiveCard(ctx context.Context) (*[]schema.LiveCard, error) {
	var err error

	var errMsg string

	narad := env.GetEnv("NARAD", "narad")

	naradColl := db.GetCollectionByName(narad)

	var liveCard []schema.LiveCard

	active := true

	filter := bson.M{"is_active": active}

	cur, err := naradColl.Find(ctx, filter)
	if err != nil {
		errMsg = "error: failed to get live card"
		err = errors.New(errMsg)
		return nil, err
	}

	defer cur.Close(ctx)

	err = cur.All(ctx, &liveCard)
	if err != nil {
		errMsg = "error: failed to decode live card"
		err = errors.New(errMsg)
		return nil, err
	}

	return &liveCard, nil
}

// # Get History Card
func GetHistoryCard(ctx context.Context) (*[]schema.HistoryCard, error) {
	var err error

	var errMsg string

	narad := env.GetEnv("NARAD", "narad")

	naradColl := db.GetCollectionByName(narad)

	var historyCard []schema.HistoryCard

	active := false

	filter := bson.M{"is_active": active}

	cur, err := naradColl.Find(ctx, filter)
	if err != nil {
		errMsg = "error: failed to get history card"
		err = errors.New(errMsg)
		return nil, err
	}

	defer cur.Close(ctx)

	err = cur.All(ctx, &historyCard)
	if err != nil {
		errMsg = "error: failed to decode history card"
		err = errors.New(errMsg)
		return nil, err
	}

	return &historyCard, nil
}

// # Mark Organisation Inactive
func MarkOrgInActive(orgID string, ctx context.Context) (string, error) {
	var err error

	var errMsg string

	narad := env.GetEnv("NARAD", "narad")

	naradColl := db.GetCollectionByName(narad)

	active := false

	filter := bson.M{"org_id": orgID}

	set := bson.M{"is_active": active}

	update := bson.M{"$set": set}

	res, err := naradColl.UpdateMany(ctx, filter, update)
	if err != nil {
		errMsg = "error: failed to mark organisation inactive"
		err = errors.New(errMsg)
		return "", err
	}

	count := int(res.ModifiedCount)

	successMsg := fmt.Sprintf("Marked organisations 'Inactive' successfully with object modified count of %d", count)

	return successMsg, nil
}

// // # Send Notification
// func SendNotification(camID primitive.ObjectID, alertMsg string, ctx context.Context) (string, error) {
// 	var err error

// 	var errMsg string

// 	cam := env.GetEnv("CAM", "camera")

// 	camColl := db.GetCollectionByName(cam)

// 	var camModel model.Camera

// 	filter := bson.M{"_id": camID}

// 	err = camColl.FindOne(ctx, filter).Decode(&camModel)
// 	if err != nil {
// 		errMsg = "error: camera not found for camera ID: " + camID.Hex()
// 		err = errors.New(errMsg)
// 		return "", err
// 	}

// 	org := env.GetEnv("ORG", "org")

// 	orgColl := db.GetCollectionByName(org)

// 	var orgModel model.Organization

// 	filter = bson.M{"cam_id": camID.Hex()}

// 	err = orgColl.FindOne(ctx, filter).Decode(&orgModel)
// 	if err != nil {
// 		errMsg = "error: organization not found for camera ID: " + camID.Hex()
// 		err = errors.New(errMsg)
// 		return "", err
// 	}

// 	user := env.GetEnv("USER", "user")

// 	userColl := db.GetCollectionByName(user)

// 	orgID := orgModel.ID

// 	var userModel model.User

// 	filter = bson.M{"org_id": orgID.Hex()}

// 	err = userColl.FindOne(ctx, filter).Decode(&userModel)
// 	if err != nil {
// 		errMsg = "error: error finding users for organization ID: " + orgID.Hex()
// 		err = errors.New(errMsg)
// 		return "", err
// 	}

// 	var userEmails []string = make([]string, 0)

// 	// for cur.Next(ctx) {
// 	// 	var user model.User

// 	// 	err = cur.Decode(&user)
// 	// 	if err != nil {
// 	// 		errMsg = "error: error decoding user for organization ID: " + orgID.Hex()
// 	// 		err = errors.New(errMsg)
// 	// 		return "", err
// 	// 	}

// 	// 	userEmails = append(userEmails, user.Email)
// 	// }

// 	userEmails = append(userEmails, userModel.Email)

// 	if len(userEmails) == 0 {
// 		errMsg = "error: no users found for organization ID: " + orgID.Hex()
// 		err = errors.New(errMsg)
// 		return "", err
// 	}

// 	var emailBody *schema.EmailResponse

// 	emailBody = &schema.EmailResponse{
// 		CamID:        camModel.ID.Hex(),
// 		CamName:      camModel.Name,
// 		OrgID:        orgModel.ID.Hex(),
// 		OrgName:      orgModel.Name,
// 		IPAddress:    camModel.IPAddress,
// 		BuildingNo:   camModel.BuildingNo,
// 		BuildingName: camModel.BuildingName,
// 		FloorNo:      camModel.FloorNo,
// 		FloorName:    camModel.FloorName,
// 		RoomNo:       camModel.RoomNo,
// 		RoomName:     camModel.RoomName,
// 		AreaName:     camModel.AreaName,
// 		Coordinates:  (schema.Coordinates)(camModel.Coordinates),
// 		Address:      camModel.Address,
// 		City:         camModel.City,
// 		State:        camModel.State,
// 		Country:      camModel.Country,
// 		Pincode:      camModel.Pincode,
// 	}

// 	err = utils.SendEmail(userEmails, alertMsg, emailBody)
// 	if err != nil {
// 		errMsg = "error: failed to send email alert notification: " + err.Error()
// 		err = errors.New(errMsg)
// 		return "", err
// 	}

// 	return "Email alert notification sent successfully", nil
// }

// # Send Email
func SendEmail(data *schema.Email) (string, error) {
	smtpUser := env.GetEnv("SMTP_USER", "email")

	smtpPass := env.GetEnv("SMTP_PASS", "password")

	smtpHost := env.GetEnv("SMTP_HOST", "host")

	smtpPort := env.GetEnv("SMTP_PORT", "port")

	smtpAuth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	smtpAddress := smtpHost + ":" + smtpPort

	var err error

	var errMsg string

	var emailBody bytes.Buffer

	emailTemplate := `<h2>New Message from {{.Name}}</h2>
	<p><strong>Email:</strong> {{.Email}}</p>
	<p><strong>Message:</strong></p>
	<p>{{.Message}}</p>`

	template, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		errMsg = "error: error parsing email template: " + err.Error()
		err = errors.New(errMsg)
		return "", err
	}

	err = template.Execute(&emailBody, data)
	if err != nil {
		errMsg = "error: error executing email template: " + err.Error()
		err = errors.New(errMsg)
		return "", err
	}

	var emails []string = make([]string, 0)
	emails = append(emails, data.Email)

	var adminEmails []string = make([]string, 0)
	adminEmails = append(adminEmails, smtpUser)

	subject := "New Message from [ " + data.Name + " ]"

	email := emailBody.String()

	message := fmt.Sprintf("Subject: %s\nMIME-Version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s", subject, email)

	msg := []byte(message)

	err = smtp.SendMail(smtpAddress, smtpAuth, smtpUser, adminEmails, msg)

	if err != nil {
		errMsg = "error: failed to send email: " + err.Error()
		err = errors.New(errMsg)
		return "", err
	}

	emailTemplate = `<h2>Thank You {{.Name}}!</h2>
<p>We have received your message and will get back to you soon.</p>
<p>Best Regards,<br>Cresa Club</p>`

	template, err = template.New("thank").Parse(emailTemplate)
	if err != nil {
		errMsg = "error: error parsing thank you email template: " + err.Error()
		err = errors.New(errMsg)
		return "", err
	}

	if err = template.Execute(&emailBody, data); err != nil {
		errMsg = "error: error executing thank you email template: " + err.Error()
		err = errors.New(errMsg)
		return "", err
	}

	subject = "Thank you for contacting us!"

	email = emailBody.String()

	message = fmt.Sprintf("Subject: %s\nMIME-Version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s", subject, email)

	msg = []byte(message)

	err = smtp.SendMail(smtpAddress, smtpAuth, smtpUser, emails, msg)

	if err != nil {
		errMsg = "error: failed to send thank you email: " + err.Error()
		err = errors.New(errMsg)
		return "", err
	}

	return "Email sent successfully!", nil
}
