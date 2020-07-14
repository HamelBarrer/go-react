package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/HamelBarrer/go-react/models"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertoTweet graba el tweet en la bd
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("goreact")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
