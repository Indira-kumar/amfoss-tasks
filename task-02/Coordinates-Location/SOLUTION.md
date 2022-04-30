
# Solutions

## Step 1:
mkdir Coordinates-Location
cd coordinates-Location

## Step 2:
mkdir North
cd North

i) cat NDegree.txt
type "9"
and type "ctrl + shift + u" and then followed by 00b0 and then a spacebar
save the file and exist using ctrl + C on terminal

ii) cat NMinutes.txt
type "5" followed by a '
save the file and exist using ctrl + C on terminal

iii)cat NSeconds.txt
let the largest even number not divisible by "x" 
bruteforce looping of
x/2 = prime number
and x-1 = prime number
Now, 38 is the answer
now I typed 38" in the NSeconds.txt
save the file and exist using ctrl + C on terminal

iv)cat NDegree.txt NMinutes.txt NSeconds.txt >> NorthCoordinate.txt

v)cp -i NorthCoordinate.txt ../
rm NorthCoordinates.txt
cd ../
mv NorthCoordinates.txt North.txt

## Step 4:
mkdir East
cd East

i) cat EDegree.txt
type "76"
and type "ctrl + shift + u" and then followed by 00b0 and then a spacebar
save the file and exist using ctrl + C on terminal

ii) cat EMinutes.txt
type "29" followed by a '
save the file and exist using ctrl + C on terminal

iii)cat ESeconds.txt
Minimum age of a US senator is 30
and 30 + 0.8 = 30.8
Type "30.8" followed by a "
save the file and exist using ctrl + c on terminal

iv)cat EDegree.txt EMinutes.txt ESeconds.txt >> EastCoordinate.txt

v)cp -i EastCoordinate.txt ../
rm EastCoordinates.txt
cd ../
mv EastCoordinates.txt East.txt

## Step 5:
cat North.txt East.txt >> Location-Coordinate.txt

## gitCommands

- git add .
- git commit -m "complete task-02"
- git push origin main

