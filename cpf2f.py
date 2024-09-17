#!/usr/bin/env python

import sys
import os

def copy_file_contents(source_file, destination_file):
    try:
        # Read the contents of the source file
        with open(source_file, 'r') as src:
            content = src.read()
        
        # Write the content to the destination file
        with open(destination_file, 'w') as dst:
            dst.write(content)
        
        print(f"Contents of {source_file} copied to {destination_file}")
    except FileNotFoundError:
        print(f"Error: File '{source_file}' or '{destination_file}' not found.")
    except IOError:
        print("Error: Unable to read or write files.")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: copyfiletofile.py <source_file>")
        sys.exit(1)

    # Get the current working directory
    current_dir = os.getcwd()
    
    # Join the current directory with the samples folder
    samples_folder = os.path.join(current_dir, "source2copy")
    
    # Construct the full path for the source file
    source_file = os.path.join(samples_folder, sys.argv[1]+".txt")
    
    # Construct the full path for the destination file
    destination_file = os.path.join(current_dir, "sample.txt")

    # if destination file does not exist, create it
    if not os.path.exists(destination_file):
        open(destination_file, 'w').close()

    # Copy the contents of the source file to the destination file
    copy_file_contents(source_file, destination_file)
