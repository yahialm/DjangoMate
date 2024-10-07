# DjangoMate CLI Tool

DjangoMate (`dm`) is a powerful CLI tool designed to simplify and enhance Django project management. It brings several helpful features to streamline common operations, enforce security best practices, and ensure code quality in Django-based projects. This tool is aimed at developers who want to be more productive by automating repetitive tasks and integrating essential development tools.

## Features

1. **Effortless Django Commands**:
    - Simplifies commonly used Django commands like `runserver`, `migrate`, and `test`.
    - Reduces command complexity by offering shorthand versions of Django's long management commands.

2. **Automated Project Setup**:
    - A single command to initialize Django projects with optional Docker and Kubernetes configurations.
    - Sets up `INSTALLED_APPS` and other configurations without manually editing the `settings.py`.

3. **Static Application Security Testing (SAST)**:
    - Integrated with tools like Bandit for security checks on the codebase.
    - Quickly scans your code for potential security issues.

4. **Code Formatting**:
    - Automates the process of running code formatters like Black and linters like Flake8 to ensure consistent code quality.
    - Can be invoked with a single command.

5. **Pre-commit Hook Setup**:
    - Sets up pre-commit hooks to ensure no insecure or poorly formatted code gets committed.
    - Automatically runs Bandit, Black, and Flake8 on every commit to ensure security and code quality are enforced at the source.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yahialm/django-simplifier.git
    ```

2. Navigate to the project directory:
    ```bash
    cd django-simplifier
    ```

3. Install the tool:
    ```bash
    go install
    ```

4. Ensure that the `dm` command is in your `PATH`.

## Usage

Once installed, you can access the DjangoMate tool using the `dm` command. Below are the available subcommands and their explanations:

### General Django Commands

- **Run the Django Development Server**:
    ```bash
    dm serve
    ```
    This command simplifies running the Django development server (`python manage.py runserver`).

- **Migrations**:
    ```bash
    dm mig
    dm mkmig
    dm shmig
    ```
    These commands handle Django migrations (`migrate`, `makemigrations`, and `showmigrations`).

- **Database Shell**:
    ```bash
    dm dbsh
    ```
    This command opens the Django database shell (`python manage.py dbshell`).

- **Django Shell**:
    ```bash
    dm sh
    ```
    This command opens the Django interactive shell (`python manage.py shell`).

- **Run Tests**:
    ```bash
    dm test
    ```
    Run your project's tests easily (`python manage.py test`).

- **Create Superuser**:
    ```bash
    dm su
    ```
    Creates a new superuser for the Django admin panel (`python manage.py createsuperuser`).

- **Create a New App**:
    ```bash
    dm app <app_name>
    ```
    Creates a new Django app and automatically adds it to `INSTALLED_APPS` in `settings.py`.

### SAST and Code Quality

- **Scan for Security Vulnerabilities**:
    ```bash
    dm scan
    ```
    Runs Bandit, a security analysis tool to find common security issues in your Python codebase.

- **Format Code**:
    ```bash
    dm format
    ```
    Automatically formats your Python code using Black and checks it with Flake8 to ensure compliance with style guides.

### Pre-commit Hooks

- **Set up Pre-commit Hooks**:
    ```bash
    dm precommit-setup
    ```
    This command sets up a pre-commit hook to automatically run Bandit, Black, and Flake8 before any commit. It ensures no insecure or badly formatted code is committed to the repository.

### Example Workflow

1. **Initialize your Django project**:
    ```bash
    dm setup
    ```

2. **Run the development server**:
    ```bash
    dm serve
    ```

3. **Create new apps**:
    ```bash
    dm app blog
    ```

4. **Run security and code checks**:
    ```bash
    dm scan
    dm format
    ```

5. **Set up pre-commit hooks**:
    ```bash
    dm precommit-setup
    ```

6. **Commit code**:
    ```bash
    git add .
    git commit -m "Added new features"
    ```

   The pre-commit hooks will automatically scan the code for security vulnerabilities and ensure it is formatted correctly.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests on GitHub. The goal of this project is to make Django development faster, more secure, and less error-prone for everyone.

### How to Contribute:

1. Fork the repository.
2. Create a new feature branch.
3. Submit a pull request with a detailed explanation of your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
