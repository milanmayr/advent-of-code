#!/bin/bash

# Initialize variables for colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

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
echo -e "${BLUE}Creating directory structure for Year $year, Day $day...${NC}"
mkdir -p $year/day-$day
touch $year/day-$day/README.md
touch $year/day-$day/go.mod
touch $year/day-$day/main.go
touch $year/day-$day/input
echo -e "${GREEN}✓ Directory structure created${NC}"

# Create README.md content
echo -e "${BLUE}Creating README.md...${NC}"
echo "## Problem
https://adventofcode.com/$year/day/$day" > $year/day-$day/README.md
echo -e "${GREEN}✓ README.md created${NC}"

# Create go.mod content
echo -e "${BLUE}Creating go.mod...${NC}"
echo "module github.com/milanmayr/advent-of-code/$year/day-$day

go 1.23" > $year/day-$day/go.mod
echo -e "${GREEN}✓ go.mod created${NC}"

# Create main.go content
echo -e "${BLUE}Creating main.go...${NC}"
echo "package main

import (

)

func main() {
    
}" > $year/day-$day/main.go
echo -e "${GREEN}✓ main.go created${NC}"

# Add directory to go.work file
echo -e "${BLUE}Updating go.work...${NC}"
if [ -f "go.work" ]; then
    go work use $year/day-$day
else
    go work init $year/day-$day
fi
echo -e "${GREEN}✓ go.work updated${NC}"

# Final message with input URL
echo -e "\n${GREEN}Setup complete! 🎄${NC}"
echo -e "${BLUE}Get your input at: ${NC}https://adventofcode.com/$year/day/$day/input"

# 
