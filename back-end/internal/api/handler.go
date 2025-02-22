package api

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/LinhNguyen2901/url-classification/back-end/internal/classifier"
	"github.com/LinhNguyen2901/url-classification/back-end/internal/db"
	"github.com/LinhNguyen2901/url-classification/back-end/internal/models"
)

type Handler struct {
	db         *db.MongoDB
	classifier *classifier.URLClassifier
}

func NewHandler() (*Handler, error) {
	log.Println("NewHandler")
	db, err := db.NewMongoDB()
	if err != nil {
		return nil, err
	}

	return &Handler{
		db:         db,
		classifier: classifier.NewURLClassifier(),
	}, nil
}

func (h *Handler) HandleClassify(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("handleClassify")
	/*
	var req models.ClassifyRequest = models.ClassifyRequest{
		URL: "https://www.google.com",
	}
		*/
	var req models.ClassifyRequest
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request",
		}, nil
	}

	// Check if URL is already classified in MongoDB
	result, err := h.db.GetClassification(ctx, req.URL)
	if err == nil {
		log.Println("result found in DB", result)
		// URL found in database, return cached result
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       result,
		}, nil
	}

	// Classify URL using ML model
	category, confidence, err := h.classifier.Classify(ctx, req.URL)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Classification failed",
		}, nil
	}
	log.Println("category", category)
	log.Println("confidence", confidence)

	// Store result in MongoDB
	response := models.ClassifyResponse{
		URL:        req.URL,
		Category:   category,
		Confidence: confidence,
	}

	if err := h.db.SaveClassification(ctx, response); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to save classification",
		}, nil
	}

	responseJSON, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseJSON),
	}, nil
}
