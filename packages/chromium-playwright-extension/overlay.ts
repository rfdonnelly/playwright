/* eslint-disable no-console */
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

import '@recorder/recorderTypes';
import type { Language } from 'playwright-core/lib/server/isomorphic/locatorGenerators';
import { asLocator } from 'playwright-core/lib/server/isomorphic/locatorGenerators';
import './globals';

export class PlaywrightExtension {
  private _overlay: PlaywrightExtensionOverlay;
  private _extensionId?: string;

  constructor() {
    this._overlay = this._initOverlay();
    this._loadInjectedScripts();
    window.addEventListener('pw:recorder:event', ((ev: CustomEvent) => {
      if (ev.detail.type === 'toggleRecorderMode')
        this._toggleOverlayVisiblity();
    }) as EventListener);
  }

  public setExtensionId(extensionId: string) {
    this._extensionId = extensionId;
  }

  private _initOverlay(): PlaywrightExtensionOverlay {
    const overlay = new PlaywrightExtensionOverlay();
    overlay.onLanguageChanged = (language: Language) => {
      chrome.runtime.sendMessage(this._extensionId!, { type: 'updateRecorderState', state: { language } });
    };
    overlay.onStopPicking = () => chrome.runtime.sendMessage(this._extensionId!, { type: 'toggleRecorderMode' });
    return overlay;
  }

  private async _loadInjectedScripts() {
    // @ts-ignore
    globalThis.module = {};
    const InjectedScript = require('../playwright-core/src/server/injected/injectedScript').InjectedScript;
    const injectedScript = new InjectedScript(false, 'javascript', 'test-id', 0, 'chromium', []);
    const ConsoleAPI = require('../playwright-core/src/server/injected/consoleApi');
    new ConsoleAPI(injectedScript);
  }

  private _toggleOverlayVisiblity() {
    this._overlay.visible = !this._overlay.visible;
  }
}

class PlaywrightExtensionOverlay {
  private _locatorInput: HTMLInputElement;
  private _visible = true;
  private _root = document.createElement('playwright-extension-overlay');
  private _shadowRoot: ShadowRoot;

  public onLanguageChanged: (language: Language) => void = () => { };
  public onStopPicking: () => void = () => { };

  constructor() {
    document.body.appendChild(this._root);
    this._shadowRoot = this._root.attachShadow({ mode: 'open' });
    this._shadowRoot.innerHTML = `
        <style>
          :host {
            cursor: move;
            position: absolute;
            z-index: 2147483646;
            background-color: hsl(0 0% 10% / 0.9);
            border-radius: 10px;
            color: #ecf0f1;
            padding: 10px;
            font-family: Helvetica;
            box-shadow: rgb(0 0 0 / 50%) 0px 0.25em 0.5em;
          }
        </style>
        <div>
          🎭 Playwright Locators
          <button>Stop Picking</button>
        </div>
        <div>
          Target
          <select class="recorder-chooser">
            <option value="javascript">Node.js</option>
            <option value="python">Python</option>
            <option value="csharp">C#</option>
            <option value="java">Java</option>
          </select>
        </div>
        <div>
          Locator
          <input type="text">
        </div>
      `;
    const selectElement = this._shadowRoot.querySelector('select')!;
    selectElement.addEventListener('change', () => {
      this.onLanguageChanged(selectElement.value as Language);
      selectElement.blur();
    });
    this._locatorInput = this._shadowRoot!.querySelector('input') as HTMLInputElement;
    this._shadowRoot.querySelector('button')!.addEventListener('click', () => this.onStopPicking());
    this._addDragListeners();
    window.addEventListener('resize', () => this._setInitialPosition());
    this._setInitialPosition();
  }

  private _setInitialPosition() {
    const rect = this._root.getBoundingClientRect();
    const padding = 10;
    const y = window.innerHeight - rect.height - padding;
    this._root.style.left = padding + 'px';
    this._root.style.top = y + 'px';
  }

  private _addDragListeners() {
    let pos1 = 0, pos2 = 0, pos3 = 0, pos4 = 0;

    const dragMouseDown = (ev: MouseEvent) => {
      const target = ev.composedPath()[0] as HTMLElement;
      if (target.tagName === 'INPUT' || target.tagName === 'SELECT')
        return;

      ev.preventDefault();
      pos3 = ev.clientX;
      pos4 = ev.clientY;
      document.addEventListener('mouseup', closeDragElement);
      document.addEventListener('mousemove', elementDrag);
    };

    const elementDrag = (ev: MouseEvent) => {
      ev.preventDefault();
      pos1 = pos3 - ev.clientX;
      pos2 = pos4 - ev.clientY;
      pos3 = ev.clientX;
      pos4 = ev.clientY;
      this._root.style.top = (this._root.offsetTop - pos2) + 'px';
      this._root.style.left = (this._root.offsetLeft - pos1) + 'px';
    };

    const closeDragElement = () => {
      document.removeEventListener('mouseup', closeDragElement);
      document.removeEventListener('mousemove', elementDrag);
    };

    this._root.addEventListener('mousedown', dragMouseDown);
  }

  get locator(): string {
    return this._locatorInput.value;
  }

  set locator(val: string) {
    this._locatorInput.value = val;
  }

  get visible(): boolean {
    return this._visible;
  }

  set visible(val: boolean) {
    this._visible = val;
    this._root.style.display = val ? 'block' : 'none';
  }
}

window.__pwExtension = new PlaywrightExtension();