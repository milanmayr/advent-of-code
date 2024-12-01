# Populates folder structure for a year's challenge

for i in $(seq 1 25)
do

    mkdir 2024
    mkdir 2024/day-$i
    touch 2024/day-$i/README.md
    touch 2024/day-$i/go.mod
    touch 2024/day-$i/main.go
    touch 2024/day-$i/input

    echo "## Problem
    https://adventofcode.com/2024/day/$i
    " > 2024/day-$i/README.md

    echo "module github.com/milanmayr/advent-of-code/2024/day-$i

go 1.23" > 2024/day-$i/go.mod

    echo "package main

import (

)

func main() {
        
}" > 2024/day-$i/main.go

done
