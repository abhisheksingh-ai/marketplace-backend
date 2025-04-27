🚲 Marketplace Backend (Golang + MongoDB + RabbitMQ + Docker)

A microservices-style backend project built with Go (Golang), MongoDB for storage, RabbitMQ for asynchronous order processing, and fully containerized using Docker.

✨ Features

REST APIs for User, Product, and Order management

Asynchronous order processing using RabbitMQ queues

MongoDB as the primary database

Microservices communication via Message Queues

Go Concurrency Patterns (Goroutines and Channels)

Fully containerized (App, MongoDB, RabbitMQ) using Docker Compose

🛆 Tech Stack

Go (Golang)

Gin Web Framework

MongoDB

RabbitMQ

Docker & Docker Compose

🚀 How to Run the Project

1. Clone the Repository

git clone https://github.com/abhisheksingh-ai/marketplace-backend.git
cd marketplace-backend

2. Build and Start the Project using Docker Compose

docker compose up --build

3. Access the Services

Backend API: http://localhost:8080

RabbitMQ Management UI: http://localhost:15672Username: guest | Password: guest

📚 Available API Endpoints

Method

Endpoint

Description

POST

/users

Create a new user

POST

/products

Create a new product

POST

/orders

Place a new order (async queued)

⚙️ Environment Variables

(Already configured inside Docker Compose)

PORT=8080
MONGO_URI=mongodb://mongo:27017
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/

🛠️ Useful Docker Commands

# To stop all containers
docker compose down

# To rebuild and restart the containers
docker compose up --build

# To check running containers
docker ps

📌 Author

GitHub: abhisheksingh-ai

make it readMe.md 
