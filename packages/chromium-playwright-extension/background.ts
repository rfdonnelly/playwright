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
import './globals';

chrome.action.onClicked.addListener(async tab => {
  if (!tab.id)
    return;

  const [{ result: shouldLoadContentScript }] = await chrome.scripting.executeScript({
    target: {
      tabId: tab.id,
    },
    func: () => {
      return !window.__pwExtension;
    },
    world: 'MAIN',
  });

  if (shouldLoadContentScript) {
    await chrome.scripting.executeScript({
      files: ['content.js'],
      target: {
        tabId: tab.id,
      },
      world: 'MAIN',
    });
    return;
  }

  await chrome.scripting.executeScript({
    target: {
      tabId: tab.id,
    },
    func: () => window.__pwExtension.toggleMode(),
    world: 'MAIN',
  });
});
