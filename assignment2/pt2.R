library("linprog")
c=c(4,4,7)
b=c(100,110,100)
A=rbind(c(1,2,8),c(7,1,4),c(4,7,1))
res=solveLP(c,b,A,maximum=TRUE)
print(res)



x <- c(9.23,7.66,9.43)

print(sum(c(1,2,8) * x))
print(sum(c(7,1,4) * x))
print(sum(c(4,7,1) * x))
