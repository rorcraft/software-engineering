Question:

Add arithmetic operators (plus, minus, times, divide) to make the following expression true:
3 1 3 6 = 8. You can use any parentheses you’d like.

Answer:

(3-1)/(3/6) = 8

-------

Question:

There is an 8x8 chess board in which two diagonally opposite corners have been cut off.
You are given 31 dominos, and a single domino can cover exactly two squares.
Can you use the 31 dominos to cover the entire board?
Prove your answer (by providing an example, or showing why it’s impossible).
https://en.wikipedia.org/wiki/Mutilated_chessboard_problem

Answer:

The puzzle is impossible to complete. A domino placed on the chessboard will always cover one white square and one black square. Therefore a collection of dominoes placed on the board will cover an equal numbers of squares of each colour. If the two white corners are removed from the board then 30 white squares and 32 black squares remain to be covered by dominoes, so this is impossible. If the two black corners are removed instead, then 32 white squares and 30 black squares remain, so it is again impossible.

-------

Question:

You have a five quart jug and a three quart jug, and an unlimited supply of water (but no measuring cups).
How would you come up with exactly four quarts of water?
NOTE: The jugs are oddly shaped, such that filling up exactly 'half' of the jug would be impossible.

Answer:

* 2x 3 quart - 5 quart = 1 quart remaining in 3 quart jug.
* clear 5 quart, fill with 1 quart
* fill 3 quart again.
* total of 4 quarts.

or

* fill 3 quart jug from 5 quart jug, 2 remaining in 5 quart jug.
* clear 3 quart jug, fill with 2 quart
* fill 5 quart, and poor 1 quart in the 2/3 full qart jug.
* total 4 quarts remaining in 5 quart jug.

-------

Question:

A bunch of men are on an island. A genie comes down and gathers everyone together and places a
magical hat on some people's heads ( i.e., at least one person has a hat). The hat is magical:
it can be seen by other people, but not by the wearer of the hat himself.
To remove the hat, those (and only those who have a hat) must dunk themselves underwater
at exactly midnight. If there are n people and c hats, how long does it take the men to remove the hats?
The men cannot tell each other (in any way) that they have a hat.

FOLLOW UP

Prove that your solution is correct.

Answer:

* if 1c, 1 night - must be himself.
* if 2c, (each c would see 1c). 1st night, no dunk, so it must be 2. - 2nd night, both will dunk.
* if 3c, (each c would see 2c). 1st night, no dunk, 2nd night no dunk. etc....

It'll take c number of nights.

--------

Question:

There is a building of 100 floors. If an egg drops from the Nth floor or above it will break.
If it's dropped from any floor below, it will not break. You're given 2 eggs.
Find N, while minimizing the number of drops for the worst case.

Answer:

http://datagenetics.com/blog/july22012/index.html

Key is _closest_ N with 2 eggs, and eggs _dont_ break if below.
Binary search would mean if 50th breaks, you'll have to start from 1.

_Minimization of Maximum Regret_
If it doesn’t break, rather than jumping up another n floors, instead we should step up just (n-1) floors (because we have one less drop available if we have to switch to one-by-one floors), so the next floor we should try is floor n + (n-1)

```
n + (n-1) + (n-2) .. + 1 >= 100
n(n + 1) / 2 >= 100
n ~= 14

The optimal strategy is to work our way through the table until the first egg breaks, then back up to one floor higher than the line above and then proceed floor-by-floor until we find the exact solution.

14, 27, 39, 50, 60, 69, 77, 84, 90, 95, 99, 100

worst case is 14 drops for 2 eggs.
```

--------

Question:

There are one hundred closed lockers in a hallway. A man begins by opening all one hundred lockers. 
Next, he closes every second locker. Then he goes to every third locker and closes it if it is open 
or opens it if it is closed (e.g., he toggles every third locker). After his one hundredth pass in the hallway, 
in which he toggles only locker number one hundred, how many lockers are open?

Answer:

1: 11111111111111...
2: 10101010101010...
3: 10001110001110...
4: 10011111001010...
5: 10010111101011...

only if n is perfect square. see [100Doors.md]





