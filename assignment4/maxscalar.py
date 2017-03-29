import numpy as np
def matrix_chain_order(p):

	n = len(p) - 1

	m = np.zeros( (n,n) )
	s = np.zeros( (n,n) )

	for l in range(1,n): #l is the chain length
		for i in range(0, n-l):
			j = i + l
			m[i,j] = float("-inf")
			for k in range(i, j):
				q = m[i,k] + m[k+1,j] + p[i]*p[k+1]*p[j+1]
				if q > m[i,j]:
					m[i,j] = q
					s[i,j] = k
	return (m, s)


m, s = matrix_chain_order([30,35,15,5,10,25])

print(m)
print(s)
