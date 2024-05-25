package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type IFirebaseApp interface {
	VerifyIDToken(ctx context.Context, token string) (string, error)
}

type firebaseApp struct {
	*firebase.App
}

func NewFirebaseApp() (*firebaseApp, error) {
	opt := option.WithCredentialsFile("./service_accout_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}

	return &firebaseApp{app}, nil
}

func (fa *firebaseApp) VerifyIDToken(ctx context.Context, token string) (string, error) {
	auth, err := fa.App.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return "", err
	}

	t, err := auth.VerifyIDToken(ctx, token)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
		return "", err
	}

	return t.UID, nil
}
