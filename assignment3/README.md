# Assigment 3

## Finding the majority

1. ./majority.go
2. We can use dynamic programming by saving and counting all the different number occurances and then see if we get majority on one of them.


## Maximizing exam marks
To maximize the score based on points {p1,p2,...,pn}=P and assumed time needed for each problem {t1,t2,...,tn}=T, we can count the score for each problem based on the actual score we get for finishing and our assumed time spent on each question. We count the score by taking the points pn/tn = score. the total points divided by the amount of time is the score we should get for each problem in general -> P/T . We can check consecutive problems if the score we get for it is "good enough" equal or higher to the average score we will solve this problem, otherwise move to the next one, adding the time we spent on each task untill we exceed the time limit.
In short, we look at the next problem and see if its "worth" starting solving it if it will give lower score than the average total score we should skip that problem.


## Huffman
Before compression 344
k 1 110010
T 1 01001
q 1 01110
s 1 101101
a 1 110110
m 1 110111
f 1 111010
w 1 111100
v 1 01000
g 1 101111
n 1 111011
r 2 0101
o 4 100
e 3 1010
c 1 110100
u 2 11100
y 1 111110
x 1 01100
b 1 101110
i 1 110101
j 1 110011
l 1 111111
  8 00
p 1 01111
t 1 101100
h 2 11000
z 1 01101
d 1 111101
After compression 194
