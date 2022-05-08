test_cases=int(input())
for i in range(0,test_cases):
    n,k=map(int,input().split())
    values = list(map(float,input().split()))
    #if sign can be changed, then max average will be the subset containing absolute maximum
    if k>0:
        absolute_values =[abs(i) for i in values]
        print(max(absolute_values))
    if k==0:
        print(max(values))