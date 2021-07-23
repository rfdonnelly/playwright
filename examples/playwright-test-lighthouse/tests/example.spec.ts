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
import { lighthouseTest, expect } from './lighthouseFixtures';
import lighthouse from 'lighthouse';

lighthouseTest('will pass the lighthouse test', async ({ cdpPort, page, browserName }) => {
  lighthouseTest.slow();
  // do e.g. your login or page navigation here, the session gets then re-used inside the lighthouse test
  await page.goto('https://github.com');
  const result = await lighthouse(page.url(), {
    port: cdpPort,
  });
  expect(result.lhr.categories.accessibility.score).toBeGreaterThan(0.9);
  expect(result.lhr.categories['best-practices'].score).toBeGreaterThan(0.9);
  expect(result.lhr.categories.performance.score).toBeGreaterThan(0.3);
  expect(result.lhr.categories.pwa.score).toBeGreaterThan(0.4);
  expect(result.lhr.categories.seo.score).toBeGreaterThan(0.9);
});
