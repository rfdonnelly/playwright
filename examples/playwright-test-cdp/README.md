# Connect to existing CDP server with Playwright Test example

This example shows how to connect to an existing CDP server with Playwright Test.

## How to use

### 1. Create the fixtures inside your project

We override the `browser` fixture so it uses `browserType.connectOverCDP` instead of a normal launch. This gives us the option to connect to a Chrome DevTools Protocol server.

See in the [baseFixtures.ts](./tests/baseFixtures.ts) for the full code.

### 2. Write a test

Since a CDP connection ends up in a normal Playwright browser instance, you are not limited in Playwright related features. The only limitation is that it only works for Chromium and not for other browsers.

## References

- [Playwright Test](https://playwright.dev/docs/intro)
- [browserType.connectOverCDP](https://playwright.dev/docs/api/class-browsertype#browser-type-connect-over-cdp)
