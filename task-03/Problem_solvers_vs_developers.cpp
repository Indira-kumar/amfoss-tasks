#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    short n,m;
    cin>>n>>m;
    if(m<n){
        vector<short> developer;
        for(int i=0;i<n;i++){
        short temp;
        cin>>temp;
        developer.push_back(temp);
        }
        vector<short> problem_solver;
        for(int i=0;i<m;i++){
        short temp;
        cin>>temp;
        problem_solver.push_back(temp);
        }
        cout<<"NO";
    }
    else{
        vector<short> developer;
        for(int i=0;i<n;i++){
        short temp;
        cin>>temp;
        developer.push_back(temp);
        }
        vector<short> problem_solver;
        for(int i=0;i<m;i++){
        short temp;
        cin>>temp;
        problem_solver.push_back(temp);
        }
    sort(developer.begin(),developer.end());
    sort(problem_solver.begin(),problem_solver.end());
    short minimum_power = 0;
    int i=0;
    for(int j=0;j<m;j++){
        if(developer[i]<problem_solver[j]){
            minimum_power+=problem_solver[j];
            i++;
            if(i==n){
                break;
            }
        }
    }
    // while(i<n || j<m){
    //     if(developer[i]<problem_solver[j]){
    //         minimum_power+=problem_solver[j];
    //         j++;
    //         i++;
    //         counter++;
    //         continue;
    //     }
    //     else if(developer[i]>=problem_solver[j]){
    //         j++;
    //         continue;
    //     }
    // }
    if(i==n){
        cout<<"YES"<<endl;
        cout<<minimum_power;
    }
    else{
        cout<<"NO";
    }
    
    }
    return 0;
}