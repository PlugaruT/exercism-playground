#!/usr/bin/env bash
### get the square number from argument
number=${1}

### check for these errors in input
error(){
    echo "Error: invalid input"
    exit 1
}

errors=(0 -1 65)

for n in ${errors[@]}; do
    ((${number} == ${n})) && error
done

### calculate the number of grains
grains=$((2**(${number}-1)))

### print the grains total or exit with error
echo ${grains#-} || exit 1

exit 0