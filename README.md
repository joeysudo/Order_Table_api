# Test_api

This is a REST api created in GoLang, the api is avaliable at https://gotestapi.herokuapp.com/orders

# Parse the csv file
Run data.py to parse the csv files in test_data folder

# Run API
Run api by 'run go *.go', the api would run on port 8000

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
