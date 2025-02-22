package models

type ClassifyRequest struct {
    URL string `json:"url"`
}

type ClassifyResponse struct {
    URL        string `json:"url"`
    Category   string `json:"category"`
    Confidence float64 `json:"confidence"`
}