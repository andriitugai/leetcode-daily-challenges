#!/bin/bash
input="200. Number of Students Unable to Eat Lunch\n"

# read -p "Enter the number and the title of the task" input

# echo $input
input=$(echo $input | sed 's/.\{1\}$//')
# echo $input

num="${input%.*}"
title="${input#*.}"
# echo $num
# echo $title

dirname=$(printf "%04d" "$num"; echo "" $title | tr "[:upper:]" "[:lower:]" | tr " " "-")
# echo $dirname

# header="${1/./\\.}"   

# header=$(echo $1 | tr "." "\\")
# echo $header

for solution in $(ls -d */)
do
    num=$(expr ${solution:0:4} - 1)
    to_val=$(( $(($num / 100 + 1 )) * 100 ))
    from_val=$(( $(( $(($num / 100 + 1 )) * 100 )) - 99 ))

    dir_to_put=$(printf "from-%04d-to-%04d" "$from_val" "$to_val")
    echo $num "--->" $dir_to_put

    if [ ! -d $dir_to_put ] 
    then
        mkdir $dir_to_put
    fi

    mv -v $solution ./$dir_to_put/
done