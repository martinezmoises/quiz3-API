For testing, make sure you create:

The database: CREATE DATABASE users;
The user role: CREATE ROLE users WITH LOGIN PASSWORD 'fishsticks';
The table:    CREATE TABLE users (
                id SERIAL PRIMARY KEY,
                email VARCHAR(255) NOT NULL,
                full_name VARCHAR(255) NOT NULL,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                version INT DEFAULT 1
              );

To test the CRUD operations: 
POST: curl -X POST http://localhost:4040/v1/users \-H "Content-Type: application/json" \-d '{"full_name": "John Doe", "email": "john.doe@example.com"}'
GET: curl -X GET http://localhost:4000/v1/users/1
PATCH/UPDATE: curl -X PATCH http://localhost:4000/v1/users/1 \-H "Content-Type: application/json" \-d '{"full_name": "Updated A. Doe", "email": "john.a.doe@example.com"}'
DELETE: curl -X DELETE http://localhost:4000/v1/users/1
