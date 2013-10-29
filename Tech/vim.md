# Vim

* Use Vundle to manage plugins.
* iTerm set to 256 colors.
* Don't use prepackaged vim config. But take a look at square/maximum-awesome
* key bindings - (mode)map <key> <command>
* <leader> to trigger keys
* grammar - <command> [<count>] <motion> e.g. f4a, ciw, de

## Resources for learning.

* https://github.com/shawncplus/vim-classes
* :help (everything) - shorthand :h
* vimtutor (comes with vim)
* http://vimcast.org, also "Practical Vim" book.
* http://vim-adventures.com/ (game)
* Sublime vintage mode.

## Modes

WIP

## Registers

WIP

### Run terminal in background 

`:!`
e.g. `:!rspec spec/model/user.rb`

__Mac OS X__
> If you have administrator privileges, you must fix an Apple-introduced problem in Mac OS X 10.5 Leopard by executing the following command, or BASH and Zsh will have the wrong PATH when executed non-interactively.

`sudo chmod ugo-x /usr/libexec/path_helper`
