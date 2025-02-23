package utils

import (
	"bytes"
	"errors"
	"fmt"
	"narad/env"
	"narad/model"
	"narad/schema"
	"net/smtp"
	"text/template"
)

// # Send Email
func SendEmail(narad model.Narad, alertMsg string) error {
	smtpUser := env.GetEnv("SMTP_USER", "email")

	smtpPass := env.GetEnv("SMTP_PASS", "password")

	smtpHost := env.GetEnv("SMTP_HOST", "host")

	smtpPort := env.GetEnv("SMTP_PORT", "port")

	smtpAuth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	smtpAddress := smtpHost + ":" + smtpPort

	var err error

	var errMsg string

	var subject, alertClass string

	if alertMsg == "Alert" {
		subject = "🔥 Fire Alert: Immediate Action Required"
		// # Yellow Background
		alertClass = "background-color: #ffeb3b; color: #000;"
	} else {
		subject = "🚨 Intrusion Detected: Human or Animal Presence"
		// # Red Background
		alertClass = "background-color: #ff0000; color: #fff;"
	}

	const emailTemplate = `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title> Fire Alert Notification</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f8f9fa;
					padding: 20px;
				}

				.container {
					max-width: 600px;
					background: #ffffff;
					padding: 20px;
					border-radius: 8px;
					box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
				}

				.alert-box {
					font-size: 18px;
					font-weight: bold;
					padding: 10px;
					text-align: center;
					border-radius: 5px;
					margin-bottom: 20px;
				}

				.high-alert {
					background-color: #ff4d4d;
					color: white;
				}

				.normal-alert {
					background-color: #ffcc00;
					color: black;
				}

				table {
					width: 100%;
					border-collapse: collapse;
					margin-top: 20px;
				}

				th,
				td {
					border: 1px solid #ddd;
					padding: 8px;
					text-align: left;
				}

				th {
					background-color: #f2f2f2;
				}

				.footer {
					margin-top: 20px;
					font-size: 14px;
					color: #555;
					text-align: center;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="alert-box {{.AlertClass}}">
					🚨 {{.AlertMsg}}
				</div>

				<p>Dear User,</p>
				<p>An alert has been triggered by one of your camera. Please find the details below:</p>

				<table>
					<tr>
						<th>Column</th>
						<th>Value</th>
					</tr>
					<tr>
						<td>Camera ID</td>
						<td> {{.CamID}} </td>
					</tr>
					<tr>
						<td>Camera Name</td>
						<td> CCTV Camera </td>
					</tr>
					<tr>
						<td>Organization ID</td>
						<td> {{.OrgID}} </td>
					</tr>
					<tr>
						<td>Organization Name</td>
						<td> {{.OrgName}} </td>
					</tr>
					<tr>
						<td>IP Address</td>
						<td> {{.IPAddress}} </td>
					</tr>
					<tr>
						<td>Building No.</td>
						<td> {{.BuildingNo}} </td>
					</tr>
					<tr>
						<td>Building Name</td>
						<td> {{.BuildingName}} </td>
					</tr>
					<tr>
						<td>Floor No.</td>
						<td> {{.FloorNo}} </td>
					</tr>
					<tr>
						<td>Floor Name</td>
						<td> {{.FloorName}} </td>
					</tr>
					<tr>
						<td>Room No</td>
						<td> {{.RoomNo}} </td>
					</tr>
					<tr>
						<td>Room Name</td>
						<td> {{.RoomName}} </td>
					</tr>
					<tr>
						<td>Area Name</td>
						<td> {{.AreaName}} </td>
					</tr>
					<tr>
						<td>Latitude</td>
						<td> {{.Latitude}} </td>
					</tr>
					<tr>
						<td>Longitude</td>
						<td> {{.Longitude}} </td>
					</tr>
					<tr>
						<td>Address</td>
						<td> {{.Address}} </td>
					</tr>
					<tr>
						<td>City</td>
						<td> {{.City}} </td>
					</tr>
					<tr>
						<td>State</td>
						<td> {{.State}} </td>
					</tr>
					<tr>
						<td>Country</td>
						<td> {{.Country}} </td>
					</tr>
					<tr>
						<td>Pincode</td>
						<td> {{.Pincode}} </td>
					</tr>
					<tr>
						<td>Map Location</td>
						<td> <a href="{{.MapLocation}}"> View on Map </a> </td>
					</tr>
				</table>

				<p>We strongly recommend taking immediate action based on the alert type:</p>
				<ul>
					<li><strong style="color: #ff4d4d;">High Alert:</strong> A human or pet animal has been detected in a
						restricted zone. Please check your camera feed urgently.</li>
					<li><strong style="color: #ffcc00;">Alert:</strong> A fire has been detected in the monitored area. Please
						verify and ensure safety measures are in place.</li>
				</ul>

				<p>For further assistance, please contact support immediately.</p>

				<div class="footer">
					<div class="footer">
						<p>Best regards,<br>Security Team - <span
								style="color: darkblue; font-weight: bold; font-style: italic;">Homebrew</span></p>
					</div>
				</div>
			</div>
		</body>
		</html>`

	emailBody := &schema.EmailResponse{
		AlertClass: alertClass,
		AlertMsg:   alertMsg,
		CamID:      narad.ID.Hex(),
		// CamName:      narad.Name,
		OrgID:        narad.OrgID,
		OrgName:      narad.OrgName,
		IPAddress:    narad.IPAddress,
		BuildingNo:   narad.BuildingNo,
		BuildingName: narad.BuildingName,
		FloorNo:      narad.FloorNo,
		FloorName:    narad.FloorName,
		RoomNo:       narad.RoomNo,
		RoomName:     narad.RoomName,
		AreaName:     narad.AreaName,
		Latitude:     narad.Coordinates.Latitude,
		Longitude:    narad.Coordinates.Longitude,
		Address:      narad.Address,
		City:         narad.City,
		State:        narad.State,
		Country:      narad.Country,
		Pincode:      narad.Pincode,
		MapLocation:  narad.MapLocation,
	}

	template, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		errMsg = "error: error parsing email template: " + err.Error()
		err = errors.New(errMsg)
		return err
	}

	var body bytes.Buffer

	err = template.Execute(&body, emailBody)
	if err != nil {
		errMsg = "error: error executing email template: " + err.Error()
		err = errors.New(errMsg)
		return err
	}

	var emails []string

	for _, user := range narad.Users {
		emails = append(emails, user.Email)
	}

	message := fmt.Sprintf("Subject: %s\nMIME-Version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s", subject, body.String())

	msg := []byte(message)

	err = smtp.SendMail(smtpAddress, smtpAuth, smtpUser, emails, msg)

	if err != nil {
		return err
	}

	return nil
}
