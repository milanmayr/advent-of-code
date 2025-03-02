#!/bin/bash

# Initialize variables with default empty values
year=""
day=""

# Parse named parameters
while [[ $# -gt 0 ]]; do
    case $1 in
        --year)
            year="$2"
            shift 2
            ;;
        --day)
            day="$2"
            shift 2
            ;;
        *)
            echo "Unknown parameter: $1"
            exit 1
            ;;
    esac
done

# Validate parameters
if [ -z "$year" ] || [ -z "$day" ]; then
    echo "Usage: $0 --year <year> --day <day>"
    echo "Example: $0 --year 2024 --day 1"
    exit 1
fi

# Create directory structure
mkdir -p $year/day-$day
touch $year/day-$day/README.md
touch $year/day-$day/go.mod
touch $year/day-$day/main.go
touch $year/day-$day/input

# Create README.md content
echo "## Problem
https://adventofcode.com/$year/day/$day" > $year/day-$day/README.md

# Create go.mod content
echo "module github.com/milanmayr/advent-of-code/$year/day-$day

go 1.23" > $year/day-$day/go.mod

# Create main.go content
echo "package main

import (

)

func main() {
    
}" > $year/day-$day/main.go

# Add directory to go.work file
if [ -f "go.work" ]; then
    go work use $year/day-$day
else
    go work init $year/day-$day
fi
