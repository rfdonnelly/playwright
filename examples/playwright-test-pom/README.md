# Page object model with Playwright Test example

This example shows how to use POMs with Playwright Test. Each POM is represented as a Playwright Test test fixture which you can use in your tests.

## How to use

### 1. Create the POM classes

Each POM ([example](./tests/poms/playwrightDevPage.ts)) is a class that accepts in the constructor the Playwright page class.

### 2. Create the test fixtures

A test fixture gets created inside the [baseFixtures.ts](./tests/baseFixtures.ts) which passes the created POM instance to the `use()` function, its value gets available inside your tests.

### 3. Use the POMs inside your test

Inside your test, you can then get the POM by the fixture name and call methods on it, see in the [example.spec.ts](./tests/example.spec.ts).

## References

- [Playwright Test](https://playwright.dev/docs/intro)
- [Page Object Model guide](https://playwright.dev/docs/test-pom)
