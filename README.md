# Gestión FCTs
## Overview

The Student Internship Documentation Manager is a web application designed to manage and organize documentation related to on-the-job training or internships of students in both middle-grade and higher-degree training cycles. It provides an intuitive platform for institutions, students, and companies to track and handle all necessary documents and interactions for a seamless internship experience.

This project is not only an efficient solution for managing internship documentation but also serves as a demonstration of practical applications using various modern technologies.

🚧 Project Status: The development of this application is currently under construction. Many features are still in progress, and the application may not be fully functional at the moment.
Key Features

- Student Profiles: Store and manage information about each student, their assigned company, and their internship progress. (In Progress)
- Document Management: Upload, categorize, and track various internship-related documents. (Planned)
- User Authentication: Secure login and user management system. (In Progress)
- Intuitive UI: A user-friendly interface built with React.js that allows easy navigation and efficient task completion.

## Technologies Used

- Backend: The server-side logic is built using Golang, offering a robust and efficient backend structure.
- Frontend: The user interface is developed using React.js, providing a reactive and modern UI.
- Containerization: Docker and Docker Compose are employed for easy development, deployment, and scalability.

## Motivation

This project was initiated as a hands-on exercise to practice and showcase skills in Golang, React.js, Docker, and other related technologies. While it serves a practical purpose in managing student internships, the primary goal is to experiment with and understand the underlying tech stack.

## Getting started
To set up the Student Internship Documentation Manager locally for development purposes, follow the steps below:
Prerequisites

Ensure you have Docker and Docker Compose installed on your machine.
### Setting Up the Backend

1. Navigate to the backend directory:
```
cd backend
```

3. Build and start the backend services:
```
docker-compose -f docker-compose-dev.yml up -d
```

This command will build the backend services based on the configurations in docker-compose-dev.yml and start them in detached mode.
### Setting Up the Databases and Database Admin Tool (pgAdmin)

1. Navigate to the databases directory:
```
cd databases
```

2. Build and start the database services:
```
docker-compose up -d
```

This command will set up your databases and also start the pgAdmin tool for database management.

### Accessing the Application

Once both the backend services and databases are running, you can access the application through your browser at the designated port (e.g., http://localhost:8080 for the backend if you've mapped to port 8080).

## Contributing

Feel free to fork this repository, submit issues, or pull requests if you think you can contribute in any way. All contributions are welcome!
