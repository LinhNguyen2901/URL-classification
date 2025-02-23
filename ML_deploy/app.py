import json
import ML.feature_eng as feature_eng
import sys
def lambda_handler(event, context):
    # Direct Lambda invocation will receive the event directly
    if isinstance(event, str):
        url = event
    # Handle dictionary input
    elif isinstance(event, dict):
        url = event.get('url')
    else:
        url = None
    
    if url is None:
        return {
            "statusCode": 400,
            "error": "URL parameter is required"
        }
    
    message = f"Received url: {url}"
    print(message)  # This will show up in CloudWatch logs
    prediction = feature_eng.get_pred(url, "./model/DL_model.pth")
    # 0: benign, 1: defacement, 2: malware, 3: phishing

    return {
        "statusCode": 200,
        "message": message,
        "url": url,
        "prediction": prediction
    }

'''
url = sys.argv[1]
message = f"Received url: {url}"
print(message)  # This will show up in CloudWatch logs
prediction = feature_eng.get_pred(url, "./model/DL_model.pth")
# 0: benign, 1: defacement, 2: malware, 3: phishing

print(f"statusCode: 200")
print(f"message: {message}")
print(f"url: {url}")
print(f"prediction: {prediction}")
'''