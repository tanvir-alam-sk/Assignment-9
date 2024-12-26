# Cat Image Viewer

This project replicates the features available on **[The Cat API](https://thecatapi.com)**, allowing users to view random cat images, search for cats based on categories, and view details such as breed and image. It is built using **Beego** for the backend and **Vanilla JavaScript** for the frontend interaction.

## Features

- Display random cat images from The Cat API.
- Search for cat images by breed or category.
- Display cat breed details such as name and origin.
- Interactive and responsive frontend UI.
- Utilizes Go channels for concurrent API calls.
- Configuration management using Beego's config system.

## Technologies Used

- **Backend:** Beego (Go)
- **Frontend:** Vanilla JavaScript
- **API:** The Cat API (https://thecatapi.com)
- **Configuration Management:** Beego config
- **Concurrency:** Go Channels
- **Testing:** Unit tests for backend and frontend
- **Version Control:** Git

## Setup Instructions

### Prerequisites

- Go 1.18 or higher
- Beego framework
- Access to The Cat API key (can be obtained from [The Cat API](https://thecatapi.com/signup))
- Node.js and npm (if using JavaScript packages)

### 1. Clone the repository

```bash
git clone https://github.com/tanvir-alam-sk Cat-Image-Viewer
cd Cat-Image-Viewer
```

## Setup Instructions

### 2. Install Go dependencies

```bash
go mod tidy
```

### 3. Set up Beego Configuration

```
appname = example-beego
httpport = 8080
runmode = dev
[catapi]
apikey = <your-api-key-here>
apiurl=https://api.thecatapi.com/v1
```

### 4. Run the Application

```
bee run

```

The server will start on [http://localhost:8080](http://localhost:8080).

### 5. Frontend Interaction

The frontend is built using vanilla JavaScript. You can interact with the application through the browser by searching for cat breeds, viewing random cat images, and browsing through categories.

### 6. Testing

To run unit tests and ensure 80% code coverage:

```
go test -cover

```
