# Simple Work Log Project

A basic work log project implemented in Go, utilizing HTML templates, HTMX for dynamic updates, and a minimal HTTP server.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Display a list of work log entries with titles and users.
- Add new work log entries dynamically through a simple form.

## Technologies Used

- [Go](https://golang.org/) - Backend server implementation.
- [HTML Templates](https://golang.org/pkg/html/template/) - For rendering HTML content.
- [HTMX](https://htmx.org/) - For dynamic updates on the frontend.
- [Tailwind CSS](https://tailwindcss.com/) - A utility-first CSS framework.

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/0asaca0rum0/work-logs-using-go.git
    ```

2. Navigate to the project directory:

    ```bash
    cd work-logs-using-go
    ```

3. Run the application:

    ```bash
    go run main.go
    ```
    or
   ```bash
    air
    ```

4. Open your web browser and go to [http://localhost:3000](http://localhost:3000) to view the work log.

## Usage

1. The main page displays a list of predefined work log entries.
2. To add a new entry, click on the "Add Entry" button.
3. Enter the title and user details in the form and submit.
4. The new entry will dynamically appear in the list.

## Contributing

If you would like to contribute to the project, feel free to fork the repository and create a pull request. Please follow the existing code style and provide a clear description of your changes.

## License

This project is open source under the [MIT License](LICENSE).
