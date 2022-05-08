# Enter your code here. Read input from STDIN. Print output to STDOUT
from collections import Counter
n = int(input())

#for taking multiple input in same line
l=[int(i) for i in input().split()]

#Counting the number occurence of the most repeated number
counter = list(Counter(l).values())
print(max(counter),end=" ")

#converting to set to find number of unique elements
total_towers=len(set(l))
print(total_towers)