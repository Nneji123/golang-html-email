# Email Sender with Template Inheritance in Go

This project is a Go application that sends emails using HTML templates with inheritance. It supports sending emails with attachments and plain text versions of the HTML content. The email templates are styled using Google Fonts and include a header, footer, and disclaimer.

## Features

- Send emails with HTML content
- Template inheritance for reusable HTML structures
- Send plain text versions of the emails
- Attach files to emails
- Use environment variables for configuration

## Prerequisites

- Go 1.16+
- SMTP server credentials
- Environment variables setup

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/Nneji123/golang-html-email.git
cd golang-html-email
```

### Install Dependencies

```sh
go get -u github.com/jordan-wright/email
go get -u github.com/joho/godotenv
```

### Setup Environment Variables

Create a `.env` file in the root of your project and add your SMTP server credentials:

```
EMAIL_HOST=smtp.your-email-provider.com
EMAIL_HOST_USER=your-email@example.com
EMAIL_HOST_PASSWORD=your-email-password
```

### HTML Templates

Create two HTML template files, `base.html` and `content.html`, in the same directory as the Go program. These templates will define the structure and content of your emails.

### Running the Program

1. Make sure you have the necessary environment variables set up in the `.env` file.
2. Ensure your HTML templates (`base.html` and `content.html`) and any attachment files are in the same directory as the Go program.
3. Run the program:

```sh
go run main.go
```

### Example Usage

The `main.go` file includes an example usage of the `SendEmail` function. Modify the example to suit your needs, including setting the email subject, recipients, context variables, and attachments.

## Project Structure

```
.
├── .env
├── base.html
├── content.html
├── main.go
├── sample.pdf
```

- `.env`: Environment variables file containing SMTP server credentials.
- `base.html`: Base HTML template with common email structure.
- `content.html`: HTML template extending the base template with specific email content.
- `main.go`: Go program file containing the `SendEmail` function and example usage.
- `sample.pdf`: Example attachment file.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
