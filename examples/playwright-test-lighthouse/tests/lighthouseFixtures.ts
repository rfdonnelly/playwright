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
import os from 'os';
import path from 'path';
import fs from 'fs';
import { test as base } from '@playwright/test';

type WorkerFixtures = {
  cdpPort: number;
};

export const lighthouseTest = base.extend<{}, WorkerFixtures>({
  cdpPort: [async ({ }, run, workerInfo) => {
    await run(9000 + workerInfo.workerIndex);
  }, { scope: 'worker' }],

  context: async ({ playwright, browserName, cdpPort, headless }, use, workerInfo) => {
    if (browserName !== 'chromium')
      workerInfo.skip();
    const userDataDir = path.join(os.tmpdir(), 'playwright-test-lighthouse');
    const context = await playwright.chromium.launchPersistentContext(userDataDir, {
      headless,
      args: [`--remote-debugging-port=${cdpPort}`],
    });
    await use(context);
    await context.close();
    await fs.promises.rmdir(userDataDir);
  }
});

export const expect = lighthouseTest.expect;
