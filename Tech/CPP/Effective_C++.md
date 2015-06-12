# Effective C++ 
Meyers, Scott (2014-11-11). Effective Modern C++: 42 Specific Ways to Improve Your Use of C++11 and C++14

### 6. `auto` watch out for 'invisible' proxy types
```c++
std::vector<bool> features(const Widget& w);

auto highPriority = features(w)[5]; // deduced std::vector<bool>::reference
// use explicit typed initializer idiom
auto highPriority = static_cast<bool>(features(w)[5])
```


