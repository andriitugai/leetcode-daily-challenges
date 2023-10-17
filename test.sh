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
# -----------------
# for num in 1 2 98 99 100 101 102 198 199 200 201 202 298 299 300 301 302
# do
#     to_value=$(printf "%04d" $((((num-1) / 100 + 1) * 100)))
#     from_value=$(printf "%04d" $((((num-1) / 100 + 1) * 100 - 99)))
#     # echo 'From' $from_value 'To' $to_value
#     dir_to_put=$(printf "from-%s-to-%s" "$from_value" "$to_value")
#     echo $num "--->" $dir_to_put
# done
# -----------------
dirname=$(printf "%04d" "$num"; echo "" $title | tr "[:upper:]" "[:lower:]" | tr " " "-")
# echo $dirname

# header="${1/./\\.}"   

# header=$(echo $1 | tr "." "\\")
# echo $header

for solution in $(ls -d */)
do
    num=${solution:0:4}
    to_value=$(printf "%04d" $((((num-1) / 100 + 1) * 100)))
    from_value=$(printf "%04d" $((((num-1) / 100 + 1) * 100 - 99)))
    dir_to_put=$(printf "from-%s-to-%s" "$from_value" "$to_value")
    echo $num "--->" $dir_to_put

    if [ ! -d $dir_to_put ] 
    then
        mkdir $dir_to_put
    fi
    mkdir $dir_to_put/$solution

    mv -v $solution ./$dir_to_put/$solution/
done