# EAS Client

This project implements an HTTP client for the API of the Otris EAS developed
by Otris AG.

## Testing

Since this is an API client for a proprietary software, we have decided to not include any docker containers / images
which would allow you to test the client against a running EAS instance.

That is why the CI pipeline skips tests which require a running EAS instance.
