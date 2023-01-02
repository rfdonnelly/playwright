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
import { chromium } from 'playwright';

(async () => {
  // Setup
  const browser = await chromium.launch();
  const context = await browser.newContext();
  const page = await context.newPage();

  await page.goto(`data:text/html,<input type="text" onpointerdown="console.log('pointerdown')">`);
  const locator = page.locator('input');
  const [message] = await Promise.all([
    page.waitForEvent('console'),
    locator.dispatchEvent('pointerdown')
  ]);
  console.log(message.text());

  // Teardown
  await context.close();
  await browser.close();
})();