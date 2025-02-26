# FXC Go Take-Home Technical Task


## Project Description
Frogo Baddins, a hobbit from the Shire, has $10,000 and wants to send money to his friend in the Philippines. He has been gathering a database of the different FX rates offered by various FX payment companies, on different days, for different transfer amounts.

## Project Goals and Approach
The main goal of this project was to demonstrate my capabilities in designing and developing a microservice that meets specific requirements. Throughout the development process, I focused on showcasing my skills in project architecture, coding best practices, and efficient problem-solving. While striving to deliver a robust and scalable solution, I also aimed to balance the time spent on the project with the quality of the results achieved. This involved making thoughtful decisions about the scope and depth of the implementation, ensuring that the final product is both functional and maintainable within the given timeframe. This project not only reflects my technical expertise but also my ability to manage and execute a development task effectively

## Task
Build a microservice that fulfills Frogo's requirements, to supply Samrise's application with data.

### Requirements
- Connect to a MariaDB instance
- Have an HTTP endpoint at an appropriately-named path, that takes the chosen date as a parameter
- Retrieve the relevant data from a table in MariaDB called `pricing`
- Transform the retrieved data into a form suitable for displaying in a table in Samrise's style (no transformation of the response should be required on the frontend)
- Return the transformed data in the response to the HTTP endpoint call

## Project Run
```
  make all
```
## Prerequisites

	•	Docker
	•	Docker Compose

## Usage

After starting the application, the microservice will be available at **http://0.0.0.0:8080**.

## Sample Request

To get FX rates for a specific date, you can make a POST request to the **/pricing** endpoint with the following JSON body:
```
curl -X POST http://0.0.0.0:8080/pricing -H "Content-Type: application/json" -d '{"date":"2024-01-10"}'
```

**Request**
```
{
    "date": "2024-01-10"
}
```

**Response**
```
[
    {
        "amount": 500,
        "details": [
            {
                "organization_name": "GlobalSettle",
                "rate": 1.3
            }
        ]
    },
    ...
]
```

## Makefile

	•	make all: Build the application and start the containers
	•	make build: Build the application
	•	make up: Start the containers
	•	make down: Stop the containers
	•	make check-db: Check if the database container is running
	•	make restart-db: Restart the database container
	•	make test: Build the application and run tests
