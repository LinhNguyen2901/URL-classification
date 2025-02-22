# Import Libraries
import pandas as pd
import pickle
from sklearn.preprocessing import StandardScaler
import lightgbm as lgb
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

# Train the LightGBM model
lgb_model = lgb.LGBMClassifier(
    num_leaves=50,             # Larger number of leaves for more complex trees
    learning_rate=0.05,        # Learning rate, smaller for more boosting rounds
    n_estimators=1000,         # More boosting rounds
    max_depth=10,              # Limit depth of each tree
    bagging_fraction=0.8,      # Use 80% of data for each iteration (prevent overfitting)
    feature_fraction=0.8,      # Use 80% of features for each tree
    n_jobs=-1                  # Use all CPUs for faster computation
)

lgb_model.fit(x_train_scaled, y_train)  # Fit the model

with open('lgb_model.pkl', 'wb') as lgb_model_file:
    pickle.dump(lgb_model, lgb_model_file)

# Save the scaler
with open('scaler.pkl', 'wb') as scaler_file:
    pickle.dump(scaler, scaler_file)

# Evaluate the LightGBM model
lgb_train_accuracy = accuracy_score(y_train, lgb_model.predict(x_train_scaled))
lgb_test_accuracy = accuracy_score(y_test, lgb_model.predict(x_test_scaled))

print(f'LightGBM Accuracy on training data: {lgb_train_accuracy:.5f}')
print(f'LightGBM Accuracy on test data: {lgb_test_accuracy:.5f}')