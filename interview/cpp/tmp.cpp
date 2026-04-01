#include <iostream>
#include <thread>
#include <mutex>
#include <cstdint>

// Use integer representing smallest currency unit (e.g., cents)
class Money {
private:
    int64_t amount;  // Store in smallest unit (cents, satoshis, etc.)
    
public:
    // Constructor with dollars and cents
    Money(int64_t dollars, int64_t cents) : amount(dollars * 100 + cents) {}
    
    // Constructor with cents only
    explicit Money(int64_t cents) : amount(cents) {}
    
    // Get amount in cents
    int64_t getCents() const { return amount; }
    
    // Get amount in dollars (with decimal)
    double getDollars() const { return amount / 100.0; }
    
    // Arithmetic operations with overflow checking
    Money operator+(const Money& other) const {
        if (amount > INT64_MAX - other.amount) {
            throw std::overflow_error("Money addition overflow");
        }
        return Money(amount + other.amount);
    }
    
    Money operator-(const Money& other) const {
        if (amount < INT64_MIN + other.amount) {
            throw std::underflow_error("Money subtraction underflow");
        }
        return Money(amount - other.amount);
    }
    
    bool operator>=(const Money& other) const { return amount >= other.amount; }
    bool operator<(const Money& other) const { return amount < other.amount; }
    
    // Format as string for display
    std::string toString() const {
        int64_t dollars = amount / 100;
        int64_t cents = std::abs(amount % 100);
        return std::to_string(dollars) + "." + 
               (cents < 10 ? "0" : "") + std::to_string(cents);
    }
};

// Thread-safe Account with Money type
class Account {
private:
    int64_t accountId;
    Money balance;
    std::mutex mtx;
    
public:
    Account(int64_t id, const Money& initialBalance) 
        : accountId(id), balance(initialBalance) {}
    
    bool withdraw(const Money& amount) {
        std::lock_guard<std::mutex> lock(mtx);
        if (balance >= amount) {
            balance = balance - amount;
            return true;
        }
        return false;
    }
    
    void deposit(const Money& amount) {
        std::lock_guard<std::mutex> lock(mtx);
        balance = balance + amount;
    }
    
    Money getBalance() {
        std::lock_guard<std::mutex> lock(mtx);
        return balance;
    }
    
    int64_t getAccountId() const { return accountId; }
    std::mutex& getMutex() { return mtx; }
};