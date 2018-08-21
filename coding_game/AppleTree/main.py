# https://www.codingame.com/training/hard/apple-tree
import sys
import math

# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

n, index = [int(i) for i in input().split()]

falling_apples = []
apples = []

print(f"index: {index}", file=sys.stderr)
print(f"N: {n}", file=sys.stderr)

for i in range(n):
    x, y, z, r = [int(j) for j in input().split()]
    apples.append([x, y, z, r])

falling_apples.append(apples[index])
del apples[index]

def count(apples, falling_apples):
    new_falling_apples = []
    new_apples = []

    for apple in apples:
        falling = False
        for falling_apple in falling_apples:
            if falling_apple[2] > apple[2]:
                distance = math.pow(math.pow(apple[0] - falling_apple[0], 2) + math.pow(apple[1] - falling_apple[1], 2), 0.5)
                if (falling_apple[3] + apple[3]) > distance:
                    falling = True
        if falling:
            new_falling_apples.append(apple)
        else:
            new_apples.append(apple)
    return new_apples, new_falling_apples


while len(falling_apples) > 0:
    apples, falling_apples = count(apples=apples, falling_apples=falling_apples)

# Write an action using print
# To debug: print("Debug messages...", file=sys.stderr)
print(f"{len(apples)}")
