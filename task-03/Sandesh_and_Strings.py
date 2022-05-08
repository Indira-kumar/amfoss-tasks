l_word = int(input())
word = input()
possible_letters = []
count = []
odd_no_of_letters = ''

for j in word:
    if j not in possible_letters:
        possible_letters.append(j)
        n = word.count(j)
        count.append(n)
k = 0
for j in count:
    if len(possible_letters)>1:
        if (j%2==1):
            odd_no_of_letters+=possible_letters[k]
            break
        k+=1
if len(possible_letters)>1:
    for i in range(len(count)):
        if count[i]%2==1:
            count[i]-=1
palindrome=''
if len(possible_letters)>1:
    for i in range(len(count)):
        if count[i]>0:
            for j in range(count[i]//2):
                palindrome+=possible_letters[i]
    stringLast = palindrome[::-1]
    palindrome += odd_no_of_letters
    palindrome +=stringLast
    print(len(palindrome))
    print(palindrome)
else:
    print(count[0])
    for i in range(0,count[0]):
        print(possible_letters[0],end='')
    

           