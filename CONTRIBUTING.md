# Contributing to BinVault

Thank you for your interest in contributing to BinVault! We welcome all contributions, whether they are bug fixes, feature requests, documentation improvements, or anything else that helps improve the project.

## Getting Started

1. **Fork the Repository**
   - Click the "Fork" button at the top of the [BinVault repository](https://github.com/polovni-mk/binvault) on GitHub.
   - Clone your fork locally:
     ```sh
     git clone https://github.com/your-username/binvault.git
     cd binvault
     ```

2. **Set Up the Development Environment**
   - Ensure you have Go installed ([Download Go](https://golang.org/dl/)).
   - Install dependencies:
     ```sh
     go mod tidy
     ```
   - Run the project:
     ```sh
     go run main.go
     ```

## Contribution Guidelines

### Reporting Issues
- Check if the issue has already been reported.
- Provide a **clear and descriptive** title.
- Include steps to reproduce the issue, expected behavior, and actual behavior.

### Submitting Pull Requests
1. **Create a new branch** for your work:
   ```sh
   git checkout -b feature-branch
   ```
2. **Make your changes and commit**:
   ```sh
   git commit -m "Description of changes"
   ```
3. **Push to your fork**:
   ```sh
   git push origin feature-branch
   ```
4. **Open a Pull Request (PR)** on GitHub.

### Code Guidelines
- Follow Go best practices and idioms.
- Run `go fmt ./...` before committing to ensure code formatting.
- Write meaningful commit messages.
- Add tests for new functionality where applicable.

## Community Guidelines
- Be respectful and inclusive.
- Provide constructive feedback.
- Follow the project's [Code of Conduct](CODE_OF_CONDUCT.md) (if applicable).

## Need Help?
If you need assistance, feel free to open a GitHub Discussion or join our community channels.

We appreciate your contributions and efforts to improve BinVault! 🚀

