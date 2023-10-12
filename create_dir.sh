#!/bin/bash
header="${1/./\\.}"
cd /Users/andriitugai/Documents/LeetCode/leetcode-daily-challenges
num="${1%.*}"
title="${1#*.}"

dirname=$(printf "%04d" "$num"; echo "" $title | tr "[:upper:]" "[:lower:]" | tr " " "-")

mkdir $dirname
echo -e "### $header\n\n" > $dirname/description.md
touch $dirname/solution.go

echo -e "\n\nTo commit the solution to GitHub the followed commands will be issued:"
echo -e '> git add .'
echo -e '> git commit -m "Added solution for $1"'
echo -e '> git push origin\n'
echo -e "Don/'t forget to save the files!"
read -p "Press any key then ready... " -n1 -s

git add .
git commit -m "Added solution for $1"
git push origin
