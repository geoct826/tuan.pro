#!/bin/bash
cd public
git status
git add -A -v
git commit -m "$1"  > /dev/null
if [ $? -eq 0 ]
then
  echo "Created git commit" $1
  set -v
  git push origin master
  exit 0
else
  echo "No change to site."
  exit 0
fi
