n=int(input())
counter=0
for i in range(1,int(n**(0.5))+1):
  if n%i==0  :
    if n/i==i:
      counter+=1
    else:
      counter+=2
print(counter)