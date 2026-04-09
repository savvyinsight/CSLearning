
// request - implement heap
// design - what is the core function, what is the essence
// priority_queue use template. function : push, pop,top , heapify , empty, size
// core : heapify up or down , need to understand it firsst.


// so if we store the heap in vector,  use indices.
/*

1.understanding heap indices(0-indexed)
parent : (i-1)/2 , left : 2*i + 1, 2*i + 2
2.When do we use heapify_up vs heapify_down?
    - which operation push or pop use heapify_up?
    - which would use heapify_down?
    - why?

    think about where the new elements starts vs where the removed elements was.

3.When push value(heapify_up). Add it to the end of the vector. to maintain the heap property(parent > children)
we need to heapify up.
    - write heapify_up.   core : index, parent=(i-1)/2 , continue heapify when heap[parent]<heap[index];

4.implement pop value (heapify_down).   --->swap with last, remove last, heapify_down.(sink from  the top)
left = 2*i+1 , right = 2*i+2. 
find the largest among heap[index], heap[left], heap[right]

5.implement push and pop, top,size.



6. what missing compare to std::priority_queue? How does it allow both `max-heap` and `min-heap`.
--->the core logic : comparision logic.
--->now hardcoded >, need to make the comparision configurable.
use std::less<T> as a prameter, and use custom comparator

*/
#include <cstddef>
#include <functional>
#include <iostream>
#include <stdexcept>
#include <utility>
#include <vector>

template<typename T>
struct greater{ //min-heap
    bool operator()(const T& a,const T& b){
        return a>b;
    }
};

template<typename T>
struct less{ //max heap
    bool operator()(const T& a,const T& b){
        return a<b;
    }
};

template<typename T,typename Container = std::vector<T>,typename Compare = std::less<T>>
class PriorityQueue{
private:
    Container heap;
    Compare comp;

    void heapify_up(int index){
        while (index>0) {
            int parent = (index-1)/2;
            if (comp(heap[parent],heap[index])) {
                std::swap(heap[index],heap[parent]);
                index = parent;
            }else {
                break;
            }
        }
    }

    void heapify_down(int index){
        int n = heap.size();
        while (true) {
            int left  = 2*index + 1;
            int right = 2*index + 2;
            int best = index ;
            if (left<n && comp(heap[best],heap[left])) {
                best = left;
            }

            if (right<n && comp(heap[best],heap[right])) {
                best = right;
            }

            if (best != index) {
                std::swap(heap[best],heap[index]);
                index = best;
            }else {
                break;
            }
        }
    }

public:
    PriorityQueue() = default;
    ~PriorityQueue() = default;

    void push(const T& value){
        heap.push_back(value);
        heapify_up(heap.size()-1);
    }

    void pop(){
        if (heap.empty()) {
            throw std::out_of_range("priority queue is empty.");
        }
        std::swap(heap[0],heap.back());
        heap.pop_back();
        heapify_down(0);        
    }

    T& top(){
        if (heap.empty()) {
            throw std::out_of_range("priority queue is empty.");
        }   
        return heap[0];
    }

    std::size_t size() const{
        return heap.size();
    }

    bool empty(){
        return heap.empty();
    }
};


void test(){
    PriorityQueue<int,std::vector<int>,less<int>> pq;
    pq.push(4);
    pq.push(2);
    pq.push(9);
    pq.push(1);
    pq.push(7);

    auto t = pq.top();
    pq.pop();
    auto t1 = pq.top();
    pq.pop();
    auto t2 = pq.top();
    pq.pop();
    auto t3 = pq.top();
    pq.pop();
    auto t4 = pq.size();
    std::cout<<t<<" "<<t1<<" "<<t2<<" "<<t3<<" "<<t4<<"\n"; 
}

int main(){
    test();
    return 0;
}