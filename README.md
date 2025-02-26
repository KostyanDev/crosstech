# **Railway Signal and Track Management Microservice**

## **Project Description**
This microservice is designed to manage railway signals and tracks. It provides a RESTful API that allows users to create, update, and retrieve railway signals and tracks while maintaining relationships between them. The system is built with **Golang**, **PostgreSQL**, and follows a clean architecture approach.

## **Project Goals and Approach**
The goal of this project was to demonstrate my ability to design and develop a scalable microservice that efficiently handles railway track and signal data. The focus was on writing clean, maintainable, and testable code while ensuring robust data integrity and API functionality. The project follows industry best practices, including structured logging, database migrations, and automated testing.

---

## **Task**
Build a microservice that allows users to manage railway signals and tracks with the following capabilities:

### **Requirements**
- Connect to a **PostgreSQL** database
- Implement CRUD operations for **signals** and **tracks**
- Establish relationships between **signals** and **tracks**
- Support soft deletion for data integrity
- Provide Swagger documentation for API endpoints
- Implement integration tests for core functionality
- Use Docker and Docker Compose for easy deployment

---
## **Makefile Commands**
| Command          | Description                                      |
|-----------------|--------------------------------------------------|
| `make all`      | Build the application and start all containers   |
| `make build`    | Build the application binary                     |
| `make up`       | Start the application and database               |
| `make down`     | Stop the application and database                |
| `make restart`  | Restart the application                          |
| `make test`     | Run all integration tests                        |
| `make clean-db` | Clean up test data from the database             |
| `make swagger`  | Generate Swagger documentation                   |

---
## **Usage** 

After starting the application, the API will be available at:
http://0.0.0.0:8080
Swagger documentation can be accessed at:
http://0.0.0.0:8080/swagger/index.html
### **API Endpoints**
#### **Signals**
| Method | Endpoint              | Description                      |
|--------|-----------------------|----------------------------------|
| `POST` | `/signal/create`      | Create a new signal             |
| `PUT`  | `/signal/update`      | Update an existing signal       |
| `GET`  | `/signals`            | Retrieve signals with filters   |

#### **Tracks**
| Method | Endpoint              | Description                      |
|--------|-----------------------|----------------------------------|
| `POST` | `/track/create`       | Create a new track              |
| `PUT`  | `/track/update`       | Update an existing track        |
| `GET`  | `/tracks`             | Retrieve tracks with filters    |

---