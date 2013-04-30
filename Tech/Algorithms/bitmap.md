* bitmap compression saves the positions of an element that is repeated very often in the message.
http://www.stoimen.com/blog/2012/01/16/computer-algorithms-data-compression-with-bitmaps/

e.g.
```
'aaaaaaaabbbbbbbb' => 'a8b8'
```

e.g.
```
[1991,1992,1993,1994,1991,1992,1993,1992,1991,1991,1991,1992,1992,1991,1991,1992, ...]

[
  [1991, '1000100011100110'],
  [1992, '0100010100011001'],
  [1993, '0010001000000000'],
  [1994, '0001000000000000']
]
```
