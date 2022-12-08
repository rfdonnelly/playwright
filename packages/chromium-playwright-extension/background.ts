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

import type { CallLog, EventData, Mode, Source } from '@recorder/recorderTypes';
import type { Language } from 'playwright-core/src/server/recorder/language';
import * as recorderSource from '../playwright-core/src/generated/recorderSource';


class Recorder  {
  private _recorderApp: IRecorderApp | null = null;
  private _mode: Mode = 'none';
  private _highlightedSelector = '';
  private _currentLanguage: Language = 'javascript';
  constructor() {

  }

  async install() {
    const recorderApp = new RecorderApp();
    this._recorderApp = recorderApp;
    recorderApp.onMessage = ((data: EventData) => {
      if (data.event === 'init') {
        const contextRecorder = new ContextRecorder(recorderApp.tabId);
        contextRecorder.install().catch(e => console.error(e));
        return;
      }

      if (data.event === 'setMode') {
        this.setMode(data.params.mode);
        return;
      }
      /* if (data.event === 'selectorUpdated') {
        this.setHighlightedSelector(data.params.language, data.params.selector);
        return;
      }
      if (data.event === 'step') {
        this._debugger.resume(true);
        return;
      }
      if (data.event === 'fileChanged') {
        this._currentLanguage = this._contextRecorder.languageName(data.params.file);
        this._refreshOverlay();
        return;
      }
      if (data.event === 'resume') {
        this._debugger.resume(false);
        return;
      }
      if (data.event === 'pause') {
        this._debugger.pauseOnNextStatement();
        return;
      }
      if (data.event === 'clear') {
        this._contextRecorder.clearScript();
        return;
      } */
    });
  }

  setMode(mode: Mode) {
    if (this._mode === mode)
      return;
    this._highlightedSelector = '';
    this._mode = mode;
    this._recorderApp?.setMode(this._mode);
  }

}

export interface IRecorderApp  {
  onMessage?: (data: EventData) => void;
  close(): Promise<void>;
  setPaused(paused: boolean): Promise<void>;
  setMode(mode: 'none' | 'recording' | 'inspecting'): Promise<void>;
  setFileIfNeeded(file: string): Promise<void>;
  setSelector(selector: string, focus?: boolean): Promise<void>;
  updateCallLogs(callLogs: CallLog[]): Promise<void>;
  setSources(sources: Source[]): Promise<void>;
}


class RecorderApp implements IRecorderApp {
  public tabId: number = null!;

  public onMessage?: (data: EventData) => void;

  constructor() {
    chrome.runtime.onMessage.addListener((data, sender) => {
      this.tabId = sender.tab?.id!;
      this.onMessage?.(data);
    });
  }

  close(): Promise<void> {
    throw new Error('Method not implemented.');
  }
  setPaused(paused: boolean): Promise<void> {
    throw new Error('Method not implemented.');
  }
  setMode(mode: 'inspecting' | 'recording' | 'none'): Promise<void> {
    chrome.tabs.sendMessage(this.tabId, { event: 'setMode', params: { mode } });
    throw new Error('Method not implemented.');
  }
  setFileIfNeeded(file: string): Promise<void> {
    throw new Error('Method not implemented.');
  }
  setSelector(selector: string, focus?: boolean | undefined): Promise<void> {
    throw new Error('Method not implemented.');
  }
  updateCallLogs(callLogs: CallLog[]): Promise<void> {
    throw new Error('Method not implemented.');
  }
  setSources(sources: Source[]): Promise<void> {
    throw new Error('Method not implemented.');
  }

  setHighlightedSelector(language: Language, selector: string) {}
  _refreshOverlay() {}
}

class ContextRecorder {
  private _tabId: number;

  constructor(tabId: number) {
    this._tabId = tabId;
  }

  async install() {
   /*  await chrome.scripting.executeScript({
      func: () => console.log('hello world'),
      target: {
        tabId: this._tabId
      },
      world: 'MAIN'
    }); */
  }
}

chrome.action.onClicked.addListener((tab) => {
  debugger
  console.log('action clicked', tab);
});

chrome.runtime.onInstalled.addListener(() => {

  new Recorder().install();

});
