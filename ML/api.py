from flask import Flask, request, jsonify
import feature_eng

app = Flask(__name__)

@app.route('/classifynn', methods=['POST'])
def classify():
    data = request.get_json()
    url = data.get('url')
    prediction = feature_eng.get_pred(url, "DL_model.pth")
    
    if not url:
        return jsonify({'error': 'URL is required'}), 400
    
    return jsonify({'url': url, 'category': prediction})

if __name__ == '__main__':
    app.run(debug=True)
