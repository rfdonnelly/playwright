/**
 * Copyright (c) Microsoft Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
import { test as base } from '@playwright/test';

type WorkerFixtures = {
  cdpPort: number;
};

export const test = base.extend<{}, WorkerFixtures>({
  browser: async ({ playwright, browserName }, use) => {
    if (browserName !== 'chromium')
      throw new Error('Only Chromium is supported');
    const browser = await playwright.chromium.connectOverCDP({
      endpointURL: 'ws://127.0.0.1:9123/devtools/browser/a9c47a67-192b-401d-9d92-5e8f09098279'
    });
    await use(browser);
    await browser.close();
  }
});

export const expect = test.expect;
