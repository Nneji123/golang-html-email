package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

// LoadEnv loads environment variables from a specified .env file.
func LoadEnv(envFilePath string) error {
	if err := godotenv.Load(envFilePath); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

// StripHTML strips HTML tags from a string to produce a plain text version.
func StripHTML(htmlStr string) string {
	var textBuffer bytes.Buffer
	inTag := false
	for _, r := range htmlStr {
		switch {
		case r == '<':
			inTag = true
		case r == '>':
			inTag = false
		case !inTag:
			textBuffer.WriteRune(r)
		}
	}
	return textBuffer.String()
}

// SendEmail sends an email with the specified subject, HTML template path, recipients, context variables, and attachments.
func SendEmail(subject, htmlTemplatePath string, toEmails, ccEmails, bccEmails []string, context map[string]interface{}, attachments []string) error {
	// Parse the base template and the content template
	tmpl, err := template.ParseFiles("base.html", htmlTemplatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Render the template with context variables
	var body bytes.Buffer
	if err := tmpl.ExecuteTemplate(&body, "base", context); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}
	htmlContent := body.String()
	plainContent := StripHTML(htmlContent)

	// Create a new email
	e := email.NewEmail()
	e.From = fmt.Sprintf("FundusAI Inc <%s>", os.Getenv("EMAIL_HOST_USER"))
	e.To = toEmails
	e.Subject = subject
	e.Text = []byte(plainContent)
	e.HTML = []byte(htmlContent)

	// Add CC and BCC if provided
	if len(ccEmails) > 0 {
		e.Cc = ccEmails
	}
	if len(bccEmails) > 0 {
		e.Bcc = bccEmails
	}

	// Attach files if provided
	for _, filePath := range attachments {
		_, fileName := filepath.Split(filePath)
		if _, err := e.AttachFile(filePath); err != nil {
			return fmt.Errorf("failed to attach file %s: %v", fileName, err)
		}
	}

	// Set up the SMTP server
	smtpAddr := fmt.Sprintf("%s:%d", os.Getenv("EMAIL_HOST"), 587)
	smtpAuth := smtp.PlainAuth("", os.Getenv("EMAIL_HOST_USER"), os.Getenv("EMAIL_HOST_PASSWORD"), os.Getenv("EMAIL_HOST"))

	// Send the email
	if err := e.Send(smtpAddr, smtpAuth); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Printf("Successfully sent email to %v", toEmails)
	return nil
}

func main() {
	// Load environment variables from .env file
	if err := LoadEnv(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Example usage
	subject := "Your Subject Here"
	htmlTemplatePath := "content.html"
	toEmails := []string{"recipient@example.com"}
	ccEmails := []string{}  // Optional, can be empty
	bccEmails := []string{} // Optional, can be empty
	context := map[string]interface{}{
		"Name": "John Doe",
		"Link": "http://example.com",
	}
	attachments := []string{"sample.pdf"} // Example attachment

	if err := SendEmail(subject, htmlTemplatePath, toEmails, ccEmails, bccEmails, context, attachments); err != nil {
		log.Fatalf("Error sending email: %v", err)
	}
}
