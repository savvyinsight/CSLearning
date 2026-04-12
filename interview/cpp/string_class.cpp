#include <cstddef>
#include <cstring>

class MyString{
private:
    char* data_;
    size_t size_;

public:
    // 1.Default Constructor
    MyString():data_(new char[1]{'\0'}),size_(0){}

    // 2.Constructor from C-string
    MyString(const char *str){
        size_ = std::strlen(str);
        data_ = new char[size_+1];
        std::memcpy(data_,str,size_+1);
    }

    // 3.Copy constructor
    MyString(const MyString &other){
        size_ = other.size_;
        data_ = new char[size_+1];
        std::memcpy(data_,other.data_,size_+1);
    }

    // 4.Move constructor
    MyString(MyString &&other) noexcept :data_(other.data_),size_(other.size_){
        other.data_ = new char[1]{'\0'};
        other.size_ = 0;
    }

    // 5.Copy Assignment
    MyString& operator=(const MyString &other){
        if(this == &other) return *this;

        char *new_data = new char[other.size_+1];
        std::memcpy(new_data,other.data_,size_+1);
        delete[] data_;
        data_ = new_data;
        size_ = other.size_;
        return *this;
    }

    // 6.Move Assignment
    MyString& operator=(MyString &&other){
        if(this == &other) return *this;
        
        delete []data_;
        data_ = other.data_;
        size_ = other.size_;

        other.data_ = new char[1]{'\0'};
        other.size_ = 0;
        return *this;
    }

    // 7.Destructor
    ~MyString(){
        delete [] data_;
    }

    // 8.size
    size_t size(){
        return size_;
    }

    // 9.access raw string
    const char* c_str(){
        return data_;
    }

    // 10.Index operator
    char& operator[](size_t index){
        return data_[index];
    }

    const char& operator[](size_t index) const{
        return  data_[index];
    }

    // 11. push_back
    void push_back(char c){
        char* new_data = new char[size_+2];
        std::memcpy(new_data,data_,size_);
        new_data[size_] = c;
        new_data[size_+1] = '\0';

        delete[] data_;
        data_ = new_data;
        size_++;
    }

    // 12.append
    void append(const MyString &other){
        char* newdata = new char[size_ + other.size_ + 1];

        std::memcpy(newdata,data_,size_);
        std::memcpy(newdata+size_, other.data_, other.size_ + 1);

        delete[] data_;
        data_ = newdata;
        size_+=other.size_;
    }
};