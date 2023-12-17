import requests
import json
import time

# Post
# url = "http://127.0.0.1:1373/v1/basket"
# headers = {'Content-Type': 'application/json'}
# data = {"data":"12345","state":False}
# response = requests.post(url, headers=headers, data=json.dumps(data))
# print(response.text)


# Get by id
# id = 900841
# url = f"http://127.0.0.1:1373/v1/basket/{id}"
# response = requests.get(url)
# print(response.text)

# Get All
# url = f"http://127.0.0.1:1373/v1/basket"
# response = requests.get(url)
# print(response.text)

# Patch
# id = 900841
# url = f"http://127.0.0.1:1373/v1/basket/{id}"
# headers = {'Content-Type': 'application/json'}
# data = {"data":"123456","state":True}
# response = requests.patch(url, headers=headers, data=json.dumps(data))
# print(response.text)

# Delete
# id = 900841
# url = f"http://127.0.0.1:1373/v1/basket/{id}"
# response = requests.delete(url)
# print(response.text)

# full test
print("Create a basket ...")
url = "http://127.0.0.1:1373/v1/basket"
headers = {'Content-Type': 'application/json'}
data = {"data":"12345","state":False}
response = requests.post(url, headers=headers, data=json.dumps(data))
id = int(response.text)
print("Response:")
print(response.text)


print("Get basket by id ...")
url = f"http://127.0.0.1:1373/v1/basket/{id}"
response = requests.get(url)
print("Response:")
print(response.text)

print("Get all baskets ...")
url = f"http://127.0.0.1:1373/v1/basket"
response = requests.get(url)
print("Response:")
print(response.text)

print("Update basket on pending ...")
url = f"http://127.0.0.1:1373/v1/basket/{id}"
headers = {'Content-Type': 'application/json'}
data = {"data":"123456789","state":True}
response = requests.patch(url, headers=headers, data=json.dumps(data))
print("Response:")
print(response.text)

print("Get basket after update ...")
url = f"http://127.0.0.1:1373/v1/basket/{id}"
response = requests.get(url)
print("Response:")
print(response.text)

print("Update basket after completed ...")
url = f"http://127.0.0.1:1373/v1/basket/{id}"
headers = {'Content-Type': 'application/json'}
data = {"data":"123456","state":False}
response = requests.patch(url, headers=headers, data=json.dumps(data))
print("Response:")
print(response.text)

print("Delete a basket by id ...")
url = f"http://127.0.0.1:1373/v1/basket/{id}"
response = requests.delete(url)
print("Response:")
print(response.text)

print("Delete a basket after deleted ...")
url = f"http://127.0.0.1:1373/v1/basket/{id}"
response = requests.delete(url)
print("Response:")
print(response.text)