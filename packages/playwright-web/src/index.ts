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

import { PlaywrightClient } from './playwrightClient';

window.addEventListener('DOMContentLoaded', () => {
  const submitButton = document.querySelector('button');
  submitButton?.addEventListener('click', async () => {
    const client = await PlaywrightClient.connect({
      wsEndpoint: 'ws://' + window.location.host + '/pw-ws',
    });
    const playwright = client.playwright();
    const browser = await playwright.chromium.launch({
      headless: false,
      slowMo: 1000,
    });
    const page = await browser.newPage();
    const myURL = document.querySelector('input')?.value;
    await page.goto(myURL!);
    const data = await page.screenshot();
    console.log(data);

    const myImage = document.querySelector('img');
    myImage!.src = URL.createObjectURL(new Blob([data], { type: 'image/png' }));
    myImage!.style.display = 'block';

    await browser.close();
    await client.close();
  });
});
