import numpy as np
import sys
def lcs_length(x, y):
	m = len(x)+1
	n = len(y)+1
	b = np.zeros( (m,n) )
	c = np.zeros( (m,n) )
	for l in range(1,len(x)+2):
		b[0,l] = l
	for l in range(1,len(y)):
		b[l,0] = l

	b = np.char.mod('%d', b)
	for i in range(1,m):
		for j in range (1,n):
			if x[i-1] == y[j-1]:
				c[i,j] = c[i-1, j-1] + 1
				b[i,j] = "x"
			elif c[i-1, j] >= c[i, j-1]:
				c[i,j] = c[i-1, j]
				b[i,j] = "^"
			else:
				c[i,j] = c[i, j-1]
				b[i,j] = "<"
	return (c, b)

x = ['C','A','C','A','Q']
y = ['C','A','D','A','C','A']

c, b = lcs_length(x,y)

print("cost")
print(c)
print("sub structure")
print(b)

def print_lcs(b,x, i,j):
	if i == 0 or j == 0:
		return
	if b[i,j] == "x":
		print_lcs(b,x,i-1,j-1)
		sys.stdout.write(x[i-1])
	elif b[i,j] == "^":
		print_lcs(b,x,i-1, j)
	else:
		print_lcs(b, x, i, j-1)

sys.stdout.write("LCS: ")
print_lcs(b, x, len(x), len(y))
print