# Vim

* Use Vundle to manage plugins.
* iTerm set to 256 colors.
* Don't use prepackaged vim config. But take a look at square/maximum-awesome
* key bindings - (mode)map <key> <command>
* <leader> to trigger keys
* grammar - `[register][num/range]<verb><motion|(i|a)<text object>>` e.g. 4fa, ciw, de
* magic key - `.` (repeat last action)
* `brew install vim` for system clipboard support.
```  
   $ vim --version | grep clipboard  
   +clipboard  
```

## Resources for learning.

* https://github.com/shawncplus/vim-classes
* :help (everything) - shorthand :h
* vimtutor (comes with vim)
* http://vimcast.org, also "Practical Vim" book.
* http://vim-adventures.com/ (game)
* http://vimgolf.com
* Sublime vintage mode.
* https://github.com/lunixbochs/actualvim
* Practical Vim - https://github.com/gitig/Practical-Vim-Notes 

## Modes

* Normal
  * `/` or `?` to search, `n` next , `shift+n` previous.
  * `shift+}` `shift+{` block
  * `%s/oldword/newworld/gc` - substitute, global, confirmation
  * `i`, `a`, `I`, `A`
  * `r` - replace
  * `ciw` change in word
  * `shift+h`, `shift+l`, `shift+m`, `zz`
* Insert
  * `ctrl+n` autocomplete
* Command
* Visual
  * `ctrl+v` block
  * `shift+v` line
  * `viw ` word, `vap` paragraph

## Registers

`:reg` - lists all registers

`"0p` - to paste from `0` register

`set clipboard=unnamed` - yank and paste with the system clipboard

Insert mode - `ctrl-r + # of register` to paste.

## Buffers

* `:ls` - show buffers
* `:b # of buffer` - go to buffer
* `:copen` - Quickfix list - syntax errors, search in project
* `ctrl+o`, `ctrl+i` to jump between last jumps. `g,`, `g;` - change lists

### Macro

TBD

### Dispatch in background 

`:!`
e.g. `:!rspec spec/model/user.rb`

or `ctrl-z` to put vim in background, do your thing and `fg` to resume

__Mac OS X__
> If you have administrator privileges, you must fix an Apple-introduced problem in Mac OS X 10.5 Leopard by executing the following command, or BASH and Zsh will have the wrong PATH when executed non-interactively.

`sudo chmod ugo-x /usr/libexec/path_helper`

### Protip

__Esc key timeout__
* vimrc `set timeoutlen=1000 ttimeoutlen=0`
* zsh # 10ms for key sequences `KEYTIMEOUT=1`
* tmux `set -sg escape-time 0`

__folding__
```
set foldmethod=syntax
set nofoldable
```

