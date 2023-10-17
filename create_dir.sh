#!/bin/bash
while true
do
    echo "Press ^C to exit or"
    read -p  "Enter the number and the title of the task: " input

    if [ "$input" != "" ]; then
        header="${input/./\\.}" # header for description.md

        cd /Users/andriitugai/Documents/LeetCode/leetcode-daily-challenges

        num="${input%.*}"
        title="${input#*.}"

        dirname=$(printf "%04d" "$num"; echo "" $title | tr "[:upper:]" "[:lower:]" | tr " " "-")

        mkdir $dirname
        echo -e "### $header\n\n" > $dirname/description.md
        touch $dirname/solution.py

        echo -e "\n\nTo commit the solution to GitHub the followed commands will be issued:"
        echo -e '> git add .'
        echo -e '> git commit -m "Added solution for $1"'
        echo -e '> git push origin\n'
        echo -e "Don/'t forget to save the files"
        read -p "Press any key then ready... " -n1 -s

        git add .
        git commit -m "Added solution for $input"
        git push origin
    fi
done