# Weather Application

A simple weather application that demonstrates the use of Go backend, Redis caching, and a modern frontend.

## Architecture

- Backend: Go REST API with Gin framework
- Frontend: HTML + JavaScript with Tailwind CSS
- Cache: Redis for weather data caching
- Containerization: Docker and docker-compose
- CI/CD: GitHub Actions

## Prerequisites

- Docker and docker-compose
- Go 1.20 or later (for local development)
- Redis (handled by docker-compose)

## Setup and Running

1. Clone the repository:
```bash
git clone <repository-url>
cd weather-app
```

2. Start the application using docker-compose:
```bash
docker-compose up --build
```

The application will be available at:
- Frontend: http://localhost
- Backend API: http://localhost:8080

## API Endpoints

### GET /weather
Get weather information for given coordinates.

Query Parameters:
- `lat`: Latitude
- `lon`: Longitude

Example:
```
GET http://localhost:8080/weather?lat=40.7128&lon=-74.0060
```

Response:
```json
{
    "temperature": 20.5,
    "description": "Sunny",
    "city": "Example City"
}
```

## Development

### Backend
The backend is written in Go using the Gin framework. It includes:
- REST API endpoints
- Redis caching for weather data
- CORS support
- Environment variable configuration

### Frontend
The frontend is a simple HTML/JavaScript application with:
- Modern UI using Tailwind CSS
- Responsive design
- Error handling
- Real-time weather updates

## CI/CD
The project includes GitHub Actions workflow that:
- Runs tests on pull requests and pushes to main
- Builds Docker images
- (Optional) Pushes to Docker registry

## License
MIT 