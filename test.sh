#!/bin/bash

filename="badexample00.txt"
echo "Running: go run . ${filename}"
go run . "${filename}"
printf "\n"
echo "---------------------------------"
printf "\n"
filename="badexample01.txt"
echo "Running: go run . ${filename}"
go run . "${filename}"
printf "\n"
echo "---------------------------------"
printf "\n"
filename="badexample02.txt"
echo "Running: go run . ${filename}"
go run . "${filename}"
printf "\n"
echo "---------------------------------"
printf "\n"
for i in $(seq -w 0 8); do
    filename="example0${i}.txt"
    
    echo "Running: go run . ${filename}"
    go run . "${filename}"
    printf "\n"
if [ "$i" -ne 08 ]; then
        echo "---------------------------------"
        printf "\n"
fi  
done      