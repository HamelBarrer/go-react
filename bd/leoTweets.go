package bd

import (
	"context"
	"log"
	"time"

	"github.com/HamelBarrer/go-react/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LeoTweets lee los tweets de un perfil
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("goreact")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
