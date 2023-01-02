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

import type { UIState } from '@recorder/recorderTypes';
import { generateSelector } from 'playwright-core/lib/server/injected/selectorGenerator';
import { asLocator } from 'playwright-core/lib/server/isomorphic/locatorGenerators';

let __pw_recorderState: UIState = {
  mode: 'none',
  language: 'javascript',
  testIdAttributeName: 'test-id',
};
window.__pw_recorderState = async () => __pw_recorderState;

window.__pw_recorderSetSelector = async (selector: string) => {
  if (window.parent !== window)
    window.parent.postMessage({ type: 'setSelector', selector }, '*');
  else
    sendLocatorToOverlay(asLocator(__pw_recorderState.language, selector));
};
window.__pw_recorderPerformAction = async () => { };

chrome.runtime.onMessage.addListener(message => {
  if (message.type === 'updateRecorderState')
    __pw_recorderState = { ...__pw_recorderState, ...message.state };
  if (message.type === 'toggleRecorderMode')
    __pw_recorderState = { ...__pw_recorderState, mode: __pw_recorderState.mode === 'inspecting' ? 'none' : 'inspecting' };
  // Forward to the InjectedScript.
  window.dispatchEvent(new CustomEvent('pw:recorder:event', { detail: message }));
});

// @ts-ignore
globalThis.module = {};
const InjectedScript = require('../playwright-core/src/server/injected/injectedScript').InjectedScript;
const injectedScript = new InjectedScript(false, 'javascript', 'test-id', 0, 'chromium', []);
const ConsoleAPI = require('../playwright-core/src/server/injected/consoleApi');
new ConsoleAPI(injectedScript);

const Recorder = require('../playwright-core/src/server/injected/recorder');
new Recorder(injectedScript);

// listen on all frame messages
window.addEventListener('message', ev => {
  const frameElement = ev.source && (ev.source as Window).frameElement;
  if (!frameElement)
    return;
  if (typeof ev.data !== 'object' || !ev.data)
    return;
  if (ev.data.type === 'setSelector') {
    const frameSelector = generateSelector(injectedScript, frameElement, __pw_recorderState.testIdAttributeName).selector;
    const selector = ev.data.selector;
    if (window.parent !== window) {
      window.parent.postMessage({ type: 'generateFrameChain', chain: [frameSelector, selector] }, '*');
    } else {
      const locator = asLocator(__pw_recorderState.language, frameSelector + ' >> internal:control=enter-frame >> ' + selector, true);
      sendLocatorToOverlay(locator);
    }
  } else if (ev.data.type === 'generateFrameChain') {
    const frameSelector = generateSelector(injectedScript, frameElement, __pw_recorderState.testIdAttributeName).selector;
    if (window.parent !== window) {
      window.parent.postMessage({ type: 'generateFrameChain', chain: [frameSelector, ...ev.data.chain] }, '*');
      return;
    } else {
      const frameSelectors = [frameSelector, ...ev.data.chain];
      const selector = frameSelectors.pop();
      const locator = asLocator(__pw_recorderState.language, frameSelectors.join(' >> internal:control=enter-frame >> ') + ' >> ' + selector, true);
      sendLocatorToOverlay(locator);
    }
  }
});

function sendLocatorToOverlay(locator: string) {
  const message = {
    type: 'setLocator',
    locator,
  };
  window.dispatchEvent(new CustomEvent('pw:recorder:event', { detail: message }));
}
