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

import type { Page } from '@playwright/test';

export class PlaywrightDevPage {
  readonly page: Page;

  constructor(page: Page) {
    this.page = page;
  }

  async goto() {
    await this.page.goto('https://playwright.dev');
  }

  async toc() {
    const text = await this.page.innerText('article ul');
    return text.split('\n').filter(line => !!line);
  }

  async getStarted() {
    await this.page.click('text=Get started');
    await this.page.waitForSelector(`text=Core concepts`);
  }

  async coreConcepts() {
    await this.getStarted();
    await this.page.click('text=Core concepts');
    await this.page.waitForSelector(`h1:has-text("Core concepts")`);
  }
}