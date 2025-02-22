package db

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/LinhNguyen2901/url-classification/back-end/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoDB() (*MongoDB, error) {
	log.Println("NewMongoDB")
	// Get MongoDB connection details from environment variables
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://host.docker.internal:27017" // Use host.docker.internal for Docker
	}

	dbName := os.Getenv("MONGODB_DATABASE")
	if dbName == "" {
		dbName = "url_classification"
	}

	collectionName := os.Getenv("MONGODB_COLLECTION")
	if collectionName == "" {
		collectionName = "url_classifications"
	}

	log.Println("Attempting to connect to MongoDB at:", mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}
	log.Println("MongoDB client created successfully")

	// Ping the database to verify connection
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return nil, err
	}
	log.Println("Ping successful")

	return &MongoDB{
		client:     client,
		database:   dbName,
		collection: collectionName,
	}, nil
}

func (d *MongoDB) GetClassification(ctx context.Context, url string) (string, error) {
	log.Println("GetClassification", url)
	collection := d.client.Database(d.database).Collection(d.collection)
	log.Println("collection")

	var result models.ClassifyResponse
	err := collection.FindOne(ctx, bson.M{"url": url}).Decode(&result)
	log.Println("decode")
	if err != nil {
		log.Println("error", err)
		if err == mongo.ErrNoDocuments {
			return "", err
		}
		return "", err
	}
	log.Println("result query found in DB", result)

	responseJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	log.Println("responseJSONinDB", string(responseJSON))

	return string(responseJSON), nil
}

func (d *MongoDB) SaveClassification(ctx context.Context, classification models.ClassifyResponse) error {
	collection := d.client.Database(d.database).Collection(d.collection)

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"url": classification.URL},
		bson.M{"$set": classification},
		options.Update().SetUpsert(true),
	)
	return err
}
