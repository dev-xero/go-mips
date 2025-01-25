import kagglehub
import pandas as pd
import requests.exceptions
from sklearn.preprocessing import LabelEncoder
from sklearn.linear_model import LinearRegression

url = "srinivasav22/sales-transactions-dataset"
prediction_file = "For_Prediction.csv"

# The function used to handle downloading the excel file from the webpage, and converting them to csv files

def download_file(url):
    try:
        # Download the sales transaction dataset from Kaggle using the provided URL
        path = kagglehub.dataset_download(url)
        print("Path to dataset files:", path)
    
        # Read the train and test datasets from the downloaded Excel files using pandas 
        train_file = pd.read_excel(f"{path}\\Train.xlsx")
        test_file = pd.read_excel(f"{path}\\Test.xlsx")
        
        # Save the training and test DataFrames to CSV files for further use
        train_file.to_csv(f"./datasets/Train.csv", index=False)
        test_file.to_csv(f"./datasets/{prediction_file}", index=False)
    
    except FileNotFoundError as e:
        print(f"File not found error: {e}")
        return None

    except requests.exceptions.RequestException as e:
        print(f"Network error occurred: {e}")
        return None
    
    except Exception as e:
        return f"Error occured -> {e}"

    return train_file, test_file

# This function is used to store the first 200 rows in the train dataset to a file called For_Prediction.csv

def store_prediction(train_file):
    try:
        # Variable to hold the amount of rows to select from the train dataset
        num_rows = 200
        
        # Accessing the first 200 rows from the train dataset
        prediction_df = train_file.iloc[:num_rows, [0, 1, 2, 3, 4, 5]]
        prediction_df.to_csv(prediction_file, index=False)
    
    except PermissionError as e:
        print(f"PermissionError: {e}. Check file permissions or the file path.")
        return None
    except Exception as e:
        print(f"Error occurred: {e}.")
    
    return prediction_file

def prepare_data(train_file, test_file):
    train_data = pd.read_csv(prediction_file)  # Load the training data into a DataFrame
    test_data = pd.read_csv(test_file)
    
    le = LabelEncoder()
    train_data["Suspicious"] = le.fit_transform(train_data["Suspicious"])

    return train_data, test_data, le

def train_model(train_data):
    X = train_data[["Quantity", "TotalSalesValues"]].values
    y = train_data["Suspicious"].values

    model = LinearRegression()
    model.fit(X, y)

    return model

def make_predictions(test_data, model, encoder):
    X_test = test_data[["Quantity", "TotalSalesValues"]].values

    predictions = model.predict(X_test)
    rounded = predictions.round()

    labels = encoder.inverse_transform(rounded.astype(int))

    return labels

def analyze_prediction_file(train_file, test_file):
    try:
        # Prepare the data
        train_data, test_data, encoder = prepare_data(train_file, test_file)
        
        # Train the model
        model = train_model(train_data)
        
        # Make predictions
        predictions = make_predictions(test_data, model, encoder)
        print("Predictions:\n", predictions)
    
    except FileNotFoundError as e:
        print(f"FileNotFoundError: {e}. The file {prediction_file} does not exist.")
    except pd.errors.ParserError as e:
        print(f"ParserError: {e}. The file {prediction_file} is not a valid CSV.")
    except PermissionError as e:
        print(f"PermissionError: {e}. Unable to read the file due to permission issues.")
    except ValueError as e:
        print(f"ValueError: {e}")
    except IndexError as e:
        print(f"IndexError: {e}")
    except Exception as e:
        print(f"An unexpected error occurred: {e}")
    return

# Main execution block
train_file, test_file = download_file(url)
prediction_stored = store_prediction(train_file)
analyze_prediction_file(train_file, test_file)