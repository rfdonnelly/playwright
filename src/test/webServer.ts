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
import { ChildProcess, exec } from 'child_process';
import net from 'net';
import stream from 'stream';
import { monotonicTime, raceAgainstDeadline } from './util';
import { WebServerConfig } from '../../types/test';
import { assert } from '../utils/utils';

const DEFAULT_ENVIRONMENT_VARIABLES = {
  'BROWSER': 'none', // Disable that create-react-app will open the page in the browser
};

const newProcessLogPrefixer = () => new stream.Transform({
  transform(this: stream.Transform, chunk: any, encoding: string, callback: stream.TransformCallback) {
    this.push(`[Playwright Test WebServer] ${chunk.toString()}`);
    callback();
  },
});

export class WebServer {
  _process?: ChildProcess;
  _port?: number;
  private _processExitedWithNonZeroStatusCode!: Promise<unknown>;
  private _config: WebServerConfig;
  constructor(config: WebServerConfig) {
    this._config = config;
    this._port = config.port;
  }
  public static async create(config: WebServerConfig): Promise<WebServer> {
    const webServer = new WebServer(config);
    if (webServer._port)
      await webServer._verifyFreePort();
    await webServer._startWebServer();
    await webServer._waitForAvailability();
    process.env.PLAYWRIGHT_BASE_URL = `http://localhost:${webServer._port}`;
    return webServer;
  }
  private async _verifyFreePort() {
    assert(this._port);
    const cancellationToken = { canceled: false };
    const portIsUsed = await Promise.race([
      new Promise(resolve => setTimeout(() => resolve(false), 100)),
      waitForSocket(this._port, 100, cancellationToken),
    ]);
    cancellationToken.canceled = true;
    if (portIsUsed)
      throw new Error(`Port ${this._port} is used, make sure that nothing is running on the port`);
  }
  private async _startWebServer(): Promise<void> {
    let collectPortResolve = (port: number) => { };
    const collectPortPromise = new Promise<number>(resolve => collectPortResolve = resolve);
    function collectPort(data: Buffer) {
      const regExp = /http:\/\/localhost:(\d+)/.exec(data.toString());
      if (regExp)
        collectPortResolve(parseInt(regExp[1], 10));
    }

    this._process = exec(this._config.command, {
      env: {
        ...DEFAULT_ENVIRONMENT_VARIABLES,
        ...process.env,
        ...this._config.env,
      },
    });
    this._process.stdout.pipe(newProcessLogPrefixer()).pipe(process.stdout);
    this._process.stderr.pipe(newProcessLogPrefixer()).pipe(process.stderr);
    if (!this._port)
      this._process.stdout.on('data', collectPort);
    let processExitedWithNonZeroStatusCodeCallback = (error: Error) => { };
    this._processExitedWithNonZeroStatusCode = new Promise((_, reject) => processExitedWithNonZeroStatusCodeCallback = reject);
    this._process.on('exit', code => processExitedWithNonZeroStatusCodeCallback(new Error(`WebServer was not able to start. Exit code: ${code}`)));
    if (!this._port)
      this._port = await collectPortPromise;
  }
  private async _waitForAvailability() {
    assert(this._port);
    const launchTimeout = this._config.timeout || 60 * 1000;
    const cancellationToken = { canceled: false };
    const { timedOut } = (await Promise.race([
      raceAgainstDeadline(waitForSocket(this._port, 100, cancellationToken), launchTimeout + monotonicTime()),
      this._processExitedWithNonZeroStatusCode,
    ])) as { timedOut: number };
    cancellationToken.canceled = true;
    if (timedOut) {
      this.kill();
      throw new Error(`failed to start web server on port ${this._port} via "${this._config.command}"`);
    }
  }
  public async kill() {
    assert(this._process);
    const waitForClose = new Promise(resolve => this._process?.on('exit', resolve));
    this._process.kill('SIGINT');
    await waitForClose;
  }
}

async function waitForSocket(port: number, delay: number, cancellationToken: { canceled: boolean }) {
  while (!cancellationToken.canceled) {
    const connected = await new Promise(resolve => {
      const conn = net
          .connect(port)
          .on('error', () => {
            resolve(false);
          })
          .on('connect', () => {
            conn.end();
            resolve(true);
          });
    });
    if (connected)
      return;
    await new Promise(x => setTimeout(x, delay));
  }
}
