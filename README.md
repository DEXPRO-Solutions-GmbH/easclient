# EAS Client

This project implements an HTTP client for the API of the Otris EAS developed
by Otris AG.

## Testing

Since this is an API client for a proprietary software, we have decided to not include any docker containers / images
which would allow you to test the client against a running EAS instance.

That is why the CI pipeline skips tests which require a running EAS instance.

If you want to run tests:

1. Have an EAS instance running and accessible from your machine. We recommend Docker.
2. Create a `.env` file in the root of the project with the following content (adjusted to your setup):

```shell
EAS_HOST=localhost:8090

EAS_STORE=store1
EAS_USER=eas_administrator
EAS_PASSWORD=changeme
```
