# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability in BinVault, we appreciate your help in disclosing it responsibly. Please do the following:

1. **Do not disclose it publicly.** Avoid posting security issues in public forums, GitHub issues, or discussions.
2. **Report the vulnerability via GitHub Security Advisories** in the repository.
3. **Provide as much detail as possible**, including:
   - A description of the vulnerability
   - Steps to reproduce the issue
   - Potential impact
   - Any suggestions for a fix (if available)
4. **Expect an acknowledgment within 48 hours** and a resolution timeline depending on severity.

## Supported Versions
We actively maintain and apply security patches to the latest stable release. The versions currently supported are:

| Version  | Supported |
|----------|-----------|
| Latest   | ✅ Yes    |
| Older than 6 months | ❌ No |

## Security Best Practices
- Keep your dependencies up to date by running `go mod tidy && go mod verify` regularly.
- Do not expose sensitive information in logs or error messages.
- Follow the principle of least privilege when configuring access permissions.
- Use HTTPS and secure authentication methods for API communication.

## Third-Party Dependencies
We regularly monitor and update third-party dependencies for security vulnerabilities. If you identify an issue with a dependency, report it using GitHub Security Advisories.

Thank you for helping us keep BinVault secure!

