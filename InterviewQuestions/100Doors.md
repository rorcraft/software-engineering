### 100 Doors in a row.

_Problem:_

you have 100 doors in a row that are all initially closed. you make 100 passes by the doors starting with the first door every time. the first time through you visit every door and toggle the door (if the door is closed, you open it, if its open, you close it). the second time you only visit every 2nd door (door #2, #4, #6). the third time, every 3rd door (door #3, #6, #9), etc, until you only visit the 100th door.

_question:_ what state are the doors in after the last pass? which are open which are closed?

_Solution:_

For example, after the first pass every door is open. on the second pass you only visit the even doors (2,4,6,8…) so now the even doors are closed and the odd ones are opened. the third time through you will close door 3 (opened from the first pass), open door 6 (closed from the second pass), etc..

question: what state are the doors in after the last pass? which are open which are closed?

solution: you can figure out that for any given door, say door #42, you will visit it for every divisor it has. so 42 has 1 & 42, 2 & 21, 3 & 14, 6 & 7. so on pass 1 i will open the door, pass 2 i will close it, pass 3 open, pass 6 close, pass 7 open, pass 14 close, pass 21 open, pass 42 close. for every pair of divisors the door will just end up back in its initial state. so you might think that every door will end up closed? well what about door #9. 9 has the divisors 1 & 9, 3 & 3. but 3 is repeated because 9 is a perfect square, so you will only visit door #9, on pass 1, 3, and 9… leaving it open at the end. only perfect square doors will be open at the end.
