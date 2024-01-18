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

        num=$(expr ${num} - 1)
        to_val=$(( $(( $(( $num / 100 )) + 1 )) * 100 ))
        from_val=$(( $(( $(( $(( $num / 100 )) + 1 )) * 100 )) - 99 ))

        dir_to_put=$(printf "from-%04d-to-%04d" "$from_val" "$to_val")
        if [ ! -d $dir_to_put ] 
        then
            mkdir $dir_to_put
        fi

        if [ ! -d "$dir_to_put/$dirname" ]
        then
            mkdir $dir_to_put/$dirname
            echo -e "### $header\n\n" >> $dir_to_put/$dirname/description.md
        fi
        
        touch $dir_to_put/$dirname/solution.py
        touch $dir_to_put/$dirname/solution.go

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