# Lighthouse with Playwright Test example

This example shows how to use Lighthouse with Playwright Test. A single persistent context is shared so that you can e.g. perform your login beforehand and the state gets shared between lighthouse.

## How to use

### 1. Create the fixtures inside your project

We create two fixtures to use it in combination with Lighthouse. The first one is `cdpPort` which contains the CDP port where Lighthouse will connect to. The second one is the `context` which we override with our custom `launchPersistentContext`. There we also pass a custom argument to Chromium so the CDP server will run on our given port.

Both get implemented inside the [lighthouseFixtures.ts](./tests/lighthouseFixtures.ts).

### 2. Write a lighthouse test

For writing tests we import the lighthouse fixtures, navigate to the URL where we want to run the tests on and then call the `lighthouse` function. You are also able to do e.g. logins beforehand since the context and its state get shared with Lighthouse.

See in the [example.spec.ts](./tests/example.spec.ts) for the full code.

## References

- [Using the Lighthouse API](https://github.com/GoogleChrome/lighthouse/blob/HEAD/docs/readme.md#using-programmatically)
- [About Fixtures](https://playwright.dev/docs/next/test-fixtures)
