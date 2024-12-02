#! /usr/bin/env bash
set -eu

DAY=$1
LONGDAY="$( printf '%02d' "$1" )"
FOLDER="day$LONGDAY"
YEAR=`date "+%Y"`

mkdir -p $YEAR/$FOLDER
echo "Created $YEAR/$FOLDER"

COOKIE=`cat session.txt`
curl -s --compressed -H "Cookie: $COOKIE" "https://adventofcode.com/$YEAR/day/$DAY/input" > $YEAR/$FOLDER/input.txt
echo "Downloadedd input.txt"

cat <<EOT >> $YEAR/$FOLDER/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        fmt.Println(scanner.Text())
	}
}
EOT
echo "Started main.go"
