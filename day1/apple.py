import re
text = open("day1-input", "r")
calibration_data = text.readlines()

# Function to process each line and extract the calibration value
def extract_calibration_value(line):
    # Find the first and last digit in the line
    first_digit = next((char for char in line if char.isdigit()), None)
    last_digit = next((char for char in reversed(line) if char.isdigit()), None)

    # Combine them to form a two-digit number, if both digits are found
    if first_digit and last_digit:
        return int(first_digit + last_digit)
    return 0


# Calculate the sum of all calibration values
total_calibration_value = sum(
    extract_calibration_value(line) for line in calibration_data
)

print(f"Solution Part 1: {total_calibration_value}")
