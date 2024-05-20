package utils

import (
    "net/smtp"
    "os"
    "github.com/joho/godotenv"
    "fmt"
    "log"
)

func SendEmail(recipientEmail, subject, message string) error {
    errEnv := godotenv.Load()
    if errEnv != nil {
        log.Fatalf("Error loading .env file: %v", errEnv)
    }

    // Set up authentication information for the SMTP server.
    auth := smtp.PlainAuth("", os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("SMTP_HOST"))

    // Convert message to bytes.
    msg := []byte("From: " + os.Getenv("EMAIL_ADDRESS") + "\r\n" +
        "To: " + recipientEmail + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" + message)

    // Send the email using the SMTP server.
    err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("EMAIL_ADDRESS"), []string{recipientEmail}, msg)
    if err != nil {
        return err
    }
    fmt.Println("\nEmail sent successfully")

    return nil
}
