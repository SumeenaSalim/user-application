# User Application
This repository employs a robust and scalable architecture pattern known as `Clean Architecture` to systematically organize functionality into well-defined layers. This approach significantly enhances code readability, testability, and maintainability, ensuring a solid foundation for the application.

To maintain a seamless database schema evolution, the repository advocates the use of the `Flyway` package. Flyway streamlines the database migration process, allowing for version control and efficient management of changes to the database schema over time. This ensures a consistent and reproducible database state across different environments.

For database querying, the application leverages the `Squirrel` package. Squirrel provides a fluent API for building SQL queries in a type-safe manner. This not only enhances code expressiveness but also minimizes the risk of SQL injection vulnerabilities. Squirrel seamlessly integrates with the Clean Architecture approach, ensuring efficient and secure data retrieval from the database.


 
# PreRequisites
- docker, docker compose

# Run
    - `make up` will start the app
    - `make migration` to set the migration
