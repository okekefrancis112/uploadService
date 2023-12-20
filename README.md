# Your Project Name

Brief description of your project.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features

- File Upload Service with Go and GoFiber
- Video upload to Cloudinary
- List uploaded files
- Download specific files
- SQLite database persistence using GORM
- ...

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go installed on your machine.
- Cloudinary account for video uploads.
- SQLite3 installed (for local database).

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/okekefrancis112/uploadService.git

2. Navigate to the project directory:

   cd uploadService

3. Install dependencies:

   go mod tidy

## Configuration
4. Create a .env file in the project root:

    DB_URL=sqlite3:/path/to/your/database.db
    CLOUDINARY_CLOUD_NAME=your_cloud_name
    CLOUDINARY_API_KEY=your_api_key
    CLOUDINARY_API_SECRET=your_api_secret

## Usage
5. To run the application, execute:

   go run main.go

    The application will be accessible at http://localhost:3000.

## Endpoints

    POST /api/v1/file - Upload files to Cloudinary
    POST /api/v1/remote - Upload remote videos to Cloudinary
    GET /api/v1/files - List uploaded files
    GET /api/v1/file/{fileID} - Download a specific file
    DELETE /api/v1/file/{fileID} - Delete a specific file


## Testing
    To run tests, use the following command:

    go test ./...

## License
    This project is licensed under the MIT License.