#Receiving the variable
x,n = input().split()
x= int(x)
n = int(n)
positions=[int(i) for i in input().split()]
direction=[int(i) for i in input().split()]

#here Law of Conservation of Momentum can be used
#So technically we can imagine the cars are passing through each other
#the answer is the car that is the farthest from the end
max_distance = 0
k=0
for i in positions:
    if direction[k]>0:
        distance = n+1-positions[k]
    else:
        distance = positions[k]
    if(distance>max_distance):
        max_distance=distance
    k+=1
        
print(max_distance)
