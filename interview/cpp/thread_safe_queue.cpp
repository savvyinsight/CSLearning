#include <condition_variable>
#include <mutex>
#include <queue>


template <typename T>
class ThreadSafeQueue{
private:
    std::queue<T> q;
    mutable std::mutex mtx;
    std::condition_variable cv;

public:
    ThreadSafeQueue() = default;

    // Disable copy(important for safety)
    ThreadSafeQueue(const ThreadSafeQueue&) = delete;
    ThreadSafeQueue operator=(const ThreadSafeQueue&) = delete;

    // Push
    void push(T v){
        {
            std::lock_guard<std::mutex> lock(mtx);
            q.push(std::move(v));
        }
        cv.notify_one(); // notify after unlock
    }

    // Blocking Pop
    T pop(){
        std::unique_lock<std::mutex> lock(mtx);

        cv.wait(lock,[this]{  //wait() calls lock.unlock() internally, then lock.lock() when waking
            return !q.empty();
        });
        // 1. lock.unlock() - release mutex so others can push
        // 2. Sleep until notify
        // 3. lock.lock() - re-acquire mutex before returning
        T value = std::move(q.front());
        q.pop();
        return value;
    }

    // Non-Blocking pop
    T try_pop(T &result){
        std::lock_guard<std::mutex> lock(mtx);
        if(!q.empty()) return false;
        result = std::move(q.front());
        q.pop();
        return true;
    }

    // Check empty
    bool empty() const{
        std::lock_guard<std::mutex> lock(mtx);
        return q.empty();
    }
};