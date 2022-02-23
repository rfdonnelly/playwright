#!/bin/bash

set -e

if [[ -z "${ANDROID_HOME}" ]]; then
    SDKDIR=$PWD/.android-sdk
    export ANDROID_HOME=${SDKDIR}
    export ANDROID_SDK_ROOT=${SDKDIR}
fi

echo "Killing previous emulators"
${ANDROID_HOME}/platform-tools/adb devices | grep emulator | cut -f1 | while read line; do ${ANDROID_HOME}/platform-tools/adb -s $line emu kill; done

echo "Starting emulator"
nohup ${ANDROID_HOME}/emulator/emulator -avd android32 -no-audio -gpu auto -writable-system &
${ANDROID_HOME}/platform-tools/adb wait-for-device shell 'while [[ -z $(getprop sys.boot_completed | tr -d '\r') ]]; do sleep 1; done; input keyevent 82'
${ANDROID_HOME}/platform-tools/adb devices
echo "Emulator started"

# wget https://www.apkmirror.com/wp-content/uploads/2022/02/99/620b115fec8d5/com.android.chrome_98.0.4758.101-475810126_minAPI24(x86_64,x86)(nodpi)_apkmirror.com.apk?verify=1645541018-UsfdqCC3LnGnQ66UoZBvepfL4jm-BirAAQMUuUmnJTA

