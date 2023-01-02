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

import chokidar from 'chokidar';
import { build } from 'esbuild';

async function buildSourceFiles() {
  const tsFiles = ['background.ts', 'overlay.ts', 'content.ts'];
  for (const file of tsFiles) {
    await build({
      entryPoints: [file],
      bundle: true,
      outfile: file.replace('.ts', '.js'),
      minify: process.argv.includes('--minify'),
      sourcemap: true,
    });
  }
}

await buildSourceFiles();

if (process.argv.includes('--watch')) {
  chokidar.watch('*.ts').on('change', () => {
    console.log('Rebuilding...');
    buildSourceFiles().catch(e => {
      console.error(e);
    });
  });
}
