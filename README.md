# Mini Inventory (Go)

A clean and efficient Inventory Management System built with **Go**, **Gin framework**, and **PostgreSQL**. This project follows modular structure and provides a comprehensive API for managing products, locations, and inventory transactions with reporting capabilities.

## 🚀 Features

- **Authentication:** Secure registration and login using Bearer Token (Stored in database).
- **User Management:** Full CRUD for users with authentication middleware.
- **Product Management:** Manage products with codes, sizes, colors, and UOM.
- **Location Management:** Define multiple storage locations/warehouses.
- **Inventory Transactions:**
  - **Receive Stock:** Record incoming items to a specific location.
  - **Deliver Stock:** Record outgoing items from a specific location.
- **Reporting System:**
  - **Stock Register:** Comprehensive view of opening, in, out, and current balance.
  - **Receive Report:** Detailed historical logs of all stock receipts.
  - **Delivery Report:** Detailed historical logs of all stock deliveries.

## 🛠 Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database:** PostgreSQL
- **SQL Library:** [Sqlx](https://github.com/jmoiron/sqlx) (for easier database mapping)
- **Hot Reload:** [Air](https://github.com/cosmtrek/air)
- **Containerization:** Docker & Docker Compose

## 🏁 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) installed.

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Rabbi728/Mini-Inventory.git
   cd Mini-Inventory
   ```

2. **Start the application using Docker Compose:**
   ```bash
   docker-compose up -d
   ```

3. **Database Migration:**
   The application handles database connections automatically. Make sure the database schema is initialized (you can check the `database/` folder for SQL scripts if any).

4. **Access the API:**
   - Base URL: `http://localhost:8080/api`
   - Database Admin (Adminer): `http://localhost:8888`

## 📖 API Documentation

### Auth Module
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/register` | Register a new user |
| POST | `/api/login` | Login and get token |
| POST | `/api/logout` | Logout (Requires Token) |
| GET | `/api/me` | Get current user profile |

### Products
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/products` | List all products |
| POST | `/api/products` | Create a new product |
| GET | `/api/products/:id` | Get product details |
| PUT | `/api/products/:id` | Update product |
| DELETE | `/api/products/:id` | Delete product |

### Locations
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/locations` | List all locations |
| POST | `/api/locations` | Create a new location |
| GET | `/api/locations/:id` | Get location details |
| PUT | `/api/locations/:id` | Update location |
| DELETE | `/api/locations/:id` | Delete location |

### Inventory
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/inventories/receive`| Add stock to a location |
| POST | `/api/inventories/delivery`| Deduct stock from a location |

### Reports
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/reports/stock-register`| Stock register with balance |
| GET | `/api/reports/receive` | Detailed receive logs |
| GET | `/api/reports/delivery` | Detailed delivery logs |

> **Note:** Most endpoints require an `Authorization: Bearer <token>` header.

## 🧪 Postman Collection
A Postman collection is included in the root directory: `Mini Inventory.postman_collection.json`. You can import it into Postman to quickly test all endpoints.

## ⚙️ Environment Variables
The following variables can be configured in `docker-compose.yml`:
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `DB_HOST`: Database host
- `DB_PORT`: Database port (default: 5432)
