/*
Tacks:
Implementing a generic Coordinate or NDimensionalPoint class in C++ 
that can handle points in any number of dimensions. 
This class will provide basic functionalities such as setting and getting coordinates, 
calculating the distance between two points, and overloading operators for easy manipulation.
*/

/*
needs : understanding of templates (generics), operator overloading, and memory management.
*/ 

/*
Key requirements:
1. Genericity: The class should be able to handle any data type (e.g., int, float, double) and
 any number of dimensions.
2. Basic functionalities: The class should allow setting and getting coordinates, calculating
 the distance between two points, and overloading operators for easy manipulation.
3. Efficiency: The implementation should be efficient in terms of memory and performance.(Dimentions
should be fixed at compile time for efficiency, using a fixed-size array instead of dynamic memory allocation.)
 */
#include <iostream>
#include <complex>
#include <cstddef>
template<typename T,std::size_t N>
class Coordinate{
private:
    T componets[N]; // Fixed-size array for efficiency

public:
    Coordinate(){
        for(std::size_t i=0;i<N;i++) componets[i] = T();
    }

    // Variadic coordinate to allow : Coordiante<int,3> p(1,2,3)
    template<typename... Args>
    Coordinate(Args... args):componets{static_cast<T>(args)...}{
        static_assert(sizeof...(args) == N, "Number of arguments must match dimensions.");
    }

    // Acceess operator (Getter/Setter)
    T& operator[](std::size_t index){
        if (index >= N) throw std::out_of_range("Index out of range");
        return componets[index];
    }

    const T& operator[](std::size_t index) const{
        if (index >= N) throw std::out_of_range("Index out of range");
        return componets[index];
    }

    // Vector addition
    Coordinate operator+(const Coordinate& other) const{
        Coordinate result; 
        for(std::size_t i=0;i<N;i++){
            result[i] = componets[i] + other[i];
        }
        return result;
    }

    // Euclidean distance
    static double distance(const Coordinate& a, const Coordinate& b){
        double sum = 0.0;
        for(std::size_t i=0;i<N;i++){
            double diff = static_cast<double>(a[i]) - static_cast<double>(b[i]);
            sum += diff * diff;
        }
        return std::sqrt(sum);
    }

    void print() const{
        std::cout<<"(";
        for(std::size_t i=0;i<N;i++){
            std::cout<<componets[i]<<(i<N-1?", ":"");
        }
        std::cout<<")\n";
    }
};

int main(){
    Coordinate<int,3> p1(1,2,3);
    Coordinate<int,3> p2(4,5,6);
    Coordinate<int,3> p3 = p1 + p2;
    p3.print(); // Output: (5, 7, 9)

    std::cout<<"Distance between p1 and p2: "<<Coordinate<int,3>::distance(p1,p2)<<"\n"; // Output: Distance between p1 and p2: 5.19615

    return 0;
}