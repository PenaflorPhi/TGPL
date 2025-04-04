# Generate a list of numbers to use as arguments for `echo`
args = " ".join([str(i) for i in range(100_000)])

with open("args", "wt") as file:
    file.write(args)
