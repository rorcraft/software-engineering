__Polymorphism__

* java
* golang interfaces

__Rust__

* http://www.rustforrubyists.com/book/chapter-10.html

Example
```rust
fn print_vec<T: std::fmt::Show>(v: &[T]) {
    for i in v.iter() {
        println!("{}", i)
    }
}

fn main() {
    let vec = [1i ,2i ,3i];

    print_vec(vec);

    let str_vec = ["hey", "there", "yo"];

    print_vec(str_vec);
}
```

The `<T: std::fmt::Show>` says: “We take any type `T` that implements the `Show` trait. Traits are sort of like ‘static duck typing’ or ‘structural typing.’ 

```rust
  trait Monster {
    fn attack(&self);
  }

    struct IndustrialRaverMonkey {
        strength: int
    }

    impl Monster for IndustrialRaverMonkey {
        fn attack(&self) {
            println!("The monkey attacks for {:d}.", self.strength)
        }
    }

    impl Monster for int {
        fn attack(&self) {
            println!("The int attacks for {:d}.", *self)
        }
    }

    fn main() {
        let monkey = IndustrialRaverMonkey {strength:35};
        monkey.attack();

        let i = 10;
        i.attack();
    }    
```

```
$ rustc dwemthy.rs && ./dwemthy
The monkey attacks for 35.
The int attacks for 10.
```

Generics
```rust
   fn monsters_attack(monsters: &[&Monster]) {
        for monster in monsters.iter() {
            monster.attack();
        }
    }
```
