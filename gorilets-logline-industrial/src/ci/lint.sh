#!/bin/bash
set -e
npx eslint ../frontend/js/**/*.js
npx stylelint ../frontend/css/**/*.css
