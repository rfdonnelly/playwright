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
import { asLocator } from 'playwright-core/lib/server/isomorphic/locatorGenerators';
import './globals';

chrome.action.onClicked.addListener(async tab => {
  if (!tab.id)
    return;

  const [{ result: shouldInjectOverlayScript }] = await chrome.scripting.executeScript({
    target: {
      tabId: tab.id,
    },
    func: () => {
      return !window.__pwExtension;
    },
    world: 'MAIN',
  });

  if (shouldInjectOverlayScript) {
    await chrome.scripting.executeScript({
      files: ['overlay.js'],
      target: {
        tabId: tab.id,
      },
      world: 'MAIN',
    });
    await chrome.scripting.executeScript({
      target: {
        tabId: tab.id,
      },
      world: 'MAIN',
      func: extensionId => window.__pwExtension.setExtensionId(extensionId),
      args: [chrome.runtime.id],
    });
    await chrome.tabs.sendMessage(tab.id, { type: 'updateRecorderState', state: { mode: 'inspecting' } });
    return;
  }

  await chrome.tabs.sendMessage(tab.id, { type: 'toggleRecorderMode' });
});

chrome.runtime.onMessageExternal.addListener(async message => {
  const forwardMessage = async () => {
    const tabs = await chrome.tabs.query({ active: true, currentWindow: true });
    for (const tab of tabs) {
      if (!tab.id || tab.id === chrome.tabs.TAB_ID_NONE)
        continue;
      chrome.tabs.sendMessage(tab.id, message).catch(e => {});
    }
  };
  if (message?.type === 'updateRecorderState')
    await forwardMessage();
  else if (message?.type === 'toggleRecorderMode')
    await forwardMessage();
});


chrome.runtime.onMessage.addListener(async (message, sender) => {
  if (message?.type === 'setSelector')
    await setSelector(message.selector, sender.tab?.id, sender.frameId);
});

async function setSelector(selector: string, tabId?: number, frameId?: number) {
  if (!tabId || isNaN(frameId))
    return;
  // build the frame tree and find the frame with the given id.
  asLocator('javascript', selector, false);
}