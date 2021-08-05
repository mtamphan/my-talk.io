import sys, random
import datetime
import sys

n = int(sys.argv[1])

A = [[random.random() for row in range(n)] for col in range(n)]
B = [[random.random() for row in range(n)] for col in range(n)]
C = [[random.random() for row in range(n)] for col in range(n)]

start = datetime.datetime.now()
for i in range(n):
    for j in range(n):
        for k in range(n):
            C[i][j] += A[i][k] * B[k][i]

end = datetime.datetime.now()
print('program took {}'.format(end - start))