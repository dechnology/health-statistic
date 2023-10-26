package fcm

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

func New(ctx context.Context) (*messaging.Client, error) {

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
