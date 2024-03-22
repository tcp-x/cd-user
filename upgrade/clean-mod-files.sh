#!/bin/bash

# Function to search for lines starting with a keyword and delete them
delete_lines_starting_with_keyword() {
    local file="$1"
    local keyword="$2"

    # Create a temporary file
    local temp_file=$(mktemp)

    # Use grep to search for lines starting with the keyword and invert the match
    # This will output lines that do not start with the keyword
    grep -v "^$keyword" "$file" > "$temp_file"

    # Replace the original file with the temporary file
    mv "$temp_file" "$file"
}

# Main function
main() {
    local filename="$1"
    local keyword="$2"

    # Check if filename and keyword are provided
    if [ -z "$filename" ] || [ -z "$keyword" ]; then
        echo "Usage: $0 <filename> <keyword>"
        exit 1
    fi

    # Check if the file exists
    if [ ! -f "$filename" ]; then
        echo "File '$filename' does not exist"
        exit 1
    fi

    # Delete lines starting with the keyword
    delete_lines_starting_with_keyword "$filename" "$keyword"
    
    echo "Lines starting with the keyword deleted successfully"
}

# Call the main function with command-line arguments
main "$@"
