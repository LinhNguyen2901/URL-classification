#!/bin/bash

# Wait for MongoDB to be ready
sleep 5

# Create the url_classifications collection with basic schema validation
docker exec mongodb mongosh --eval '
  db = db.getSiblingDB("urldb");
  db.createCollection("url_classifications", {
    validator: {
      $jsonSchema: {
        bsonType: "object",
        required: ["url", "classification"],
        properties: {
          url: {
            bsonType: "string",
            description: "URL to be classified - required"
          },
          classification: {
            bsonType: "string",
            description: "Classification result - required"
          },
          timestamp: {
            bsonType: "date",
            description: "When the classification was made"
          }
        }
      }
    }
  });
  print("url_classifications collection created successfully");
'
