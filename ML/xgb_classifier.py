# Import Libraries
import pandas as pd
import pickle
from sklearn.preprocessing import StandardScaler
import xgboost as xgb
from sklearn.metrics import accuracy_score

# Load the datasets
x_train = pd.read_csv('x-train.csv')
x_test = pd.read_csv('x-test.csv')
y_train = pd.read_csv('y-train.csv')
y_test = pd.read_csv('y-test.csv')

x_train = x_train.iloc[:, 1:]
x_test = x_test.iloc[:, 1:]
y_train = y_train['labels']
y_test = y_test['labels']

# Preprocessing
x_train.fillna(0, inplace=True)
x_test.fillna(0, inplace=True)

scaler = StandardScaler()
x_train_scaled = scaler.fit_transform(x_train)
x_test_scaled = scaler.transform(x_test)

print(x_train_scaled.shape)
print(x_test.shape)
print(y_train.shape)
print(y_test.shape)

# Train the XGBoost model
xgb_model = xgb.XGBClassifier(
    eval_metric='mlogloss', 
    max_depth=6,              # Deeper trees for more complex patterns
    learning_rate=0.05,       # Lower learning rate for more training iterations
    n_estimators=1000,        # More boosting rounds
    subsample=0.8,            # Sample 80% of the data
    colsample_bytree=0.8,     # Use 80% of features for each tree
    n_jobs=-1                 # Use all CPUs for faster computation
)

xgb_model.fit(x_train_scaled, y_train)  # Fit the model

with open('xgb_model.pkl', 'wb') as xgb_model_file:
    pickle.dump(xgb_model, xgb_model_file)

# Save the scaler
with open('scaler.pkl', 'wb') as scaler_file:
    pickle.dump(scaler, scaler_file)

# Evaluate the XGBoost model
xgb_train_accuracy = accuracy_score(y_train, xgb_model.predict(x_train_scaled))
xgb_test_accuracy = accuracy_score(y_test, xgb_model.predict(x_test_scaled))

print(f'XGBoost Accuracy on training data: {xgb_train_accuracy:.5f}')
print(f'XGBoost Accuracy on test data: {xgb_test_accuracy:.5f}')