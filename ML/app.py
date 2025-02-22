import json

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
    
    message = f"Received URL: {url}"
    print(message)  # This will show up in CloudWatch logs
    
    return {
        "statusCode": 200,
        "message": message,
        "url": url
    }