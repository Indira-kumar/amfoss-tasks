#include <vector>
#include <iostream>
using namespace std;

int main() 
{
    int no_of_cities, no_of_queries;
    cin >> no_of_cities;
    int expenditure[no_of_cities];
    for(int i=0;i<no_of_cities;i++){
        int temp;
        cin>>temp;
        expenditure[i]=temp;
    }
    cin>>no_of_queries;
    int sum[no_of_queries];
    for(int i =0;i<no_of_queries;i++){
        int left_limit, right_limit;
        cin>>left_limit>>right_limit;
        sum[i]=0;
        for(int j=left_limit-1;j<=right_limit-1;j++){
            sum[i]=sum[i]+expenditure[j];
        }
    }
    for(int i=0;i<no_of_queries;i++){
        cout<<sum[i]<<endl;
    }
    return 0;
}