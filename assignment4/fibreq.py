def req_fib(n):
	if (n == 0):
		return 0
	if (n == 1):
		return 1
	return req_fib(n-2) + req_fib(n-1)


def dyn_fib(n, l):
	# init values not part of algorithm


	## algorithm start
	if (n == 0):
		return 0
	if (n == 1):
		return 1
	if (l[n-1] >= 0):
		return l[n-1]
	l[n-1] = dyn_fib(n-2, l) + dyn_fib(n-1, l)
	return l[n-1]


fibNo = 10
print("fibonnacci no:" + str(fibNo))
print("recursive : ")
print(req_fib(fibNo))	# complexity = O(2^n)
print("dynamic : ")
l = []
if (len(l)<1):
	l = []
	for i in range(fibNo):
		l.append(-1)

print(dyn_fib(fibNo, l)) # complexity = O(n)
