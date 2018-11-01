# None of the following seem to work, even when run directly
# from the cli (and not executing this file).
#
# If I run these commands directly from an interactive bash
# then the program starts, I get an sh shell (so far so good),
# the first command runs and returns fine
# but then the shell doesn't run any more commands,
# pressing enter just adds new lines, ctrl+c doesn't kill it,
# neither does sending EOF (ctrl+d).
#
# What's wrong?

#(./solve.sh ; cat) | ./02
cat input_sequence.txt - | ./02
