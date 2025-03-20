run:
	go run main.go

run1:
	go run valid.go

test-without-age-zero:
	curl -X POST http://localhost:8090/user \
          -H "Content-Type: application/json" \
          -d '{"name": "Alice", "age": 1, "age2": 0}'
#{"age":1,"age2":0,"name":"Alice"}
test-with-age-zero:
	curl -X POST http://localhost:8090/user \
          -H "Content-Type: application/json" \
          -d '{"name": "Alice", "age": 0, "age2": 0}'
#{"error":"Key: 'User.Age' Error:Field validation for 'Age' failed on the 'required' tag"}
