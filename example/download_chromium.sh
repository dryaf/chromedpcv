#!/usr/bin/env bash

rm -rf chrome-mac
curl -Lf https://download-chromium.appspot.com/dl/Mac?type=snapshots --output chrome-mac.zip
unzip -q chrome-mac.zip
rm chrome-mac.zip
