#include <functional>
#include <iostream>
#include <mutex>
#include <thread>
#include <vector>

class Account{
    public:
    int id;
    double balance;
    std::mutex mtx;

    Account(int _id,double _balance):id(_id),balance(_balance){}
};


void transfer(Account& f,Account& t,double amount){
    std::lock(f.mtx, t.mtx);
    std::lock_guard<std::mutex> lock1(f.mtx,std::adopt_lock);
    std::lock_guard<std::mutex> lock2(t.mtx,std::adopt_lock);

    if (f.balance>=amount) {
        f.balance-=amount;
        t.balance+=amount;
    }
}

int main(){
    Account a1(1,1000);
    Account a2(2,1000);

    std::vector<std::thread> threads;

    // Simulate concurrent transfers
    for (int i = 0; i < 1000; ++i) {
        threads.emplace_back(transfer, std::ref(a1), std::ref(a2), 2);
        threads.emplace_back(transfer, std::ref(a2), std::ref(a1), 1);
    }

    for (auto& t : threads) {
        t.join();
    }

    std::cout << "Final Balance:\n";
    std::cout << "Account 1: " << a1.balance << "\n";
    std::cout << "Account 2: " << a2.balance << "\n";
}