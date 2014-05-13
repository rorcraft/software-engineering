__Vectors__: 0 indexed collection, allow mixed types.

```clojure
[3 2 1]

(get [3 2 1] 0)
; => 3

;; adds to end of vector
(conj [3 2 1] 4)
; => [3 2 1 4]
```

__List__: 

```clojure
(= '(1 2 3) (list 1 2 3))

(nth '(1 2 3) 0)
=> 1

;; adds to head of list
(conj '(1 2 3) 4)
; => '(4 1 2 3)
```

__Sets__: Unique collection, allow mixed types.

```clojure
(= #{"hello" "world" 1 10}
   (set ["hello" 1 10 "world"]))

(conj #{:a :b} :b)
; => #{:a :b}

(= (get #{:a :b} :a)
   (:a #{:a :b})
   (:a))
   
(hash-set [1 1 3 3 2 2])
(sorted-set [1 1 2 2 3 3])
```

__Map__: similar to dictionaries or hashes

```clojure

(get {:a 0 :b 1} :b)
; => 1

(get {:a 0 :b {:c "ho hum"}} :b)
; => {:c "ho hum"}

```
