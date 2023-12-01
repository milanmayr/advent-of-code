# Populates folder structure for a year's challenge

for i in $(seq 1 25)
do

mkdir day-$i
touch day-$i/README.md
touch day-$i/go.mod
touch day-$i/main.go
touch day-$i/input

echo "## Problem
https://adventofcode.com/2023/day/$i
" > day-$i/README.md

echo "module github.com/milanmayr/advent-of-code/2023/day-$i

go 1.21" > day-$i/go.mod

echo "package main

import (

)

func main() {
	
}" > day-$i/main.go

done
