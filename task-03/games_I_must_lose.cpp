#include<iostream>
#include<algorithm>
using namespace std;
int no_of_moves(int);

int main(){
    int n,m;
    cin>>n>>m;
    int girl_moves = no_of_moves(n);
    int i;
    int no_of_answers=0;
    for(i=1;i<=m;i++){
        int possible_moves = no_of_moves(i);
        if (possible_moves<girl_moves){
            no_of_answers++;
        }
    }
    int answers[no_of_answers];
    int c=0;
    for(i=1;i<=m;i++){
        int possible_moves = no_of_moves(i);
        if (possible_moves<girl_moves){
            answers[c]=i;
            c++;
        }
    }
    cout<<no_of_answers<<endl;
    for(i=0;i<no_of_answers;i++){
        cout<<answers[i]<<" ";
    }
}

int no_of_moves(int i){
    int counter =0;
    while(i%2==0){
        counter++;
        i/=2;
    }
    return counter;
}