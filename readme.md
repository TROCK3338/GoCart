# GoCart

This is a full-stack e-commerce shopping cart application built as a web service. The application allows users to sign up, log in, add items to a cart, and convert that cart into an order. The project consists of a backend API built with Go and a frontend web application built with React.

## Project Flow

The user flow is as follows:
1.  **User Signup/Login**: Users can create an account or log in to access the shopping features.
2.  **Add Items to Cart**: Logged-in users can view available items and add them to their single, dedicated shopping cart.
3.  **Place Order**: The user's cart can be converted into an official order.

## Technologies Used

### Backend
* **Language**: Go
* **Web Framework**: [gin-gonic/gin](https://github.com/gin-gonic/gin)
* **ORM**: [gorm.io/gorm](https://github.com/go-gorm/gorm)
* **Database**: SQLite (for simplicity)
* **Testing**: [onsi/ginkgo](https://github.com/onsi/ginkgo) (recommended, not implemented in this version)
* **CORS**: [gin-contrib/cors](https://github.com/gin-contrib/cors)

### Frontend
* **Framework**: [React](https://reactjs.org/)
* **State Management**: React Hooks (`useState`, `useEffect`)
* **HTTP Client**: [Axios](https://axios-http.com/)

## Getting Started

### Prerequisites
* [Go](https://go.dev/doc/install) (1.18 or higher)
* [Node.js](https://nodejs.org/) and [npm](https://www.npmjs.com/)
* [Postman](https://www.postman.com/) (for API testing)

### Backend Setup
1.  Navigate to the `backend` directory.
    ```sh
    cd GoCart/backend
    ```
2.  Install Go dependencies.
    ```sh
    go mod tidy
    ```
3.  Run the server. This will also create the `gocart.db` file and migrate the database schema.
    ```sh
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

### Frontend Setup
1.  Navigate to the `frontend` directory.
    ```sh
    cd GoCart/frontend
    ```
2.  Install Node.js dependencies.
    ```sh
    npm install
    ```
3.  Start the development server.
    ```sh
    npm start
    ```
    The React application will open in your browser on `http://localhost:3000`.

## API Endpoints

The backend provides the following REST API endpoints:

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/users` | [cite_start]Creates a new user [cite: 20] |
| `POST` | `/users/login` | [cite_start]Logs in an existing user and returns a token [cite: 20] |
| `POST` | `/items` | [cite_start]Creates a new item [cite: 20] |
| `GET` | `/items` | [cite_start]Lists all items [cite: 20] |
| `POST` | `/carts` | [cite_start]Creates a cart and adds items to it (requires a token) [cite: 20, 21] |
| `GET` | `/carts` | [cite_start]Lists all carts (requires a token) [cite: 20, 21] |
| `POST` | `/orders` | [cite_start]Converts a cart to an order (requires a token) [cite: 20, 21] |
| `GET` | `/orders` | [cite_start]Lists all orders for the authenticated user (requires a token) [cite: 20, 21] |

**Note**: The token obtained from the `/users/login` endpoint must be included in the `Authorization` header for all protected endpoints.

Made with ❤️ by Aman Singhal