#  Test-Order-Table-api

This repo houses the assets used to build the Rest api for Order-Table
This project was powered by Golang

To check the front end repository see https://github.com/joeysudo/Test_app

You can access the API https://gotestapi.herokuapp.com/

# Parse the csv file

```
Run 'data.py' to parse the csv files in test_data folder
```


# Run API

```
go run *.go
```
the api would run on port 8000

Get all orders[method:Get]:
http://localhost:8000/orders

Get order by ID [method:Get]:
http://localhost:8000/orders/{id}

Create new order [method:Post]:
http://localhost:8000/orders

Update order [method:Post]:
http://localhost:8000/orders/{id}

Delete order by ID [method:Delete]:
http://localhost:8000/orders/{id}
