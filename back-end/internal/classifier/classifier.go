package classifier

import (
    "context"
    "encoding/json"
    "bytes"
    "net/http"
    "log"
    
    "github.com/aws/aws-sdk-go-v2/service/lambda"
    "github.com/aws/aws-sdk-go-v2/config"
)

type URLClassifier struct {
    lambdaClient *lambda.Client
    functionName string
}

func NewURLClassifier() *URLClassifier {
    // Create new AWS Lambda client
    cfg, err := config.LoadDefaultConfig(context.Background())
    if err != nil {
        // In a real application, you might want to handle this error differently
        panic(err)
    }
    lambdaClient := lambda.NewFromConfig(cfg)

    return &URLClassifier{
        lambdaClient: lambdaClient,
        functionName: "MLFunction",
    }
}

func (c *URLClassifier) Classify(ctx context.Context, url string) (category string, confidence float64, err error) {
    // Trim http:// or https:// from the URL
    trimmedURL := url
    if prefix := "https://"; len(url) >= len(prefix) && url[:len(prefix)] == prefix {
        trimmedURL = url[len(prefix):]
    } else if prefix := "http://"; len(url) >= len(prefix) && url[:len(prefix)] == prefix {
        trimmedURL = url[len(prefix):]
    }
    
    // Prepare the request payload
    payload := struct {
        URL string `json:"url"`
    }{
        URL: trimmedURL,
    }

    // Convert payload to JSON
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", 0.0, err
    }
    log.Println("jsonData", string(jsonData))
    // Make HTTP request to the Flask API
    resp, err := http.Post("http://host.docker.internal:5000/classifynn", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Println("error posting to flask api", err)
        return "", 0.0, err
    }
    log.Println("response", resp.Body)
    defer resp.Body.Close()

    // Parse the response
    var response struct {
        URL      string `json:"url"`
        Category int    `json:"category"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        log.Println("error decoding response", err)
        return "", 0.0, err
    }
    log.Println("response decode", response)
    // Since the Flask API doesn't return confidence, we'll return 1.0 as default
    if response.Category == 0 {
        return "benign", 1.0, nil
    } else if response.Category == 1 {
        return "defacement", 1.0, nil
    } else if response.Category == 2 {
        return "malware", 1.0, nil
    } else if response.Category == 3 {
        return "phishing", 1.0, nil
    }
    return "unknown", 0.0, nil
}