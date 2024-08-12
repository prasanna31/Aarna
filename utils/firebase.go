package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	firebaseApp       *firebase.App
	firebaseStorage   *storage.Client
	firebaseMessaging *messaging.Client
)

func InitFirebase() {
	ctx := context.Background()

	// Initialize Firebase App
	opt := option.WithCredentialsFile("path/to/your/firebase-service-account-key.json")
	var err error
	firebaseApp, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Initialize Firebase Storage
	firebaseStorage, err = firebaseApp.Storage(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase storage: %v", err)
	}

	// Initialize Firebase Messaging
	firebaseMessaging, err = firebaseApp.Messaging(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase messaging: %v", err)
	}
}

// UploadFile uploads a file to Firebase Storage
func UploadFile(filePath string, fileName string) (string, error) {
	ctx := context.Background()

	// Get a reference to the Firebase Storage bucket
	bucket, err := firebaseStorage.DefaultBucket(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get bucket: %v", err)
	}

	// Upload file to Firebase Storage
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	object := bucket.Object(fileName)
	writer := object.NewWriter(ctx)
	if _, err := writer.ReadFrom(file); err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// Get file URL
	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get file attributes: %v", err)
	}
	return attrs.MediaLink, nil
}

// SendNotification sends a notification using Firebase Cloud Messaging
func SendNotification(token string, title string, body string) error {
	ctx := context.Background()

	// Create a message to send
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	// Send the message
	_, err := firebaseMessaging.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send notification: %v", err)
	}
	return nil
}
