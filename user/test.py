import requests

url = "http://localhost:5000"
response = requests.get(url)

print(response.text)