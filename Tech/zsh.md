## ZShell

* http://www.acm.uiuc.edu/workshops/zsh/
* http://zsh.sourceforge.net/Intro/intro_toc.html
* http://zsh.sourceforge.net/Guide/zshguide.html

----

$PROMPT == $PS1

----
`cd OLD NEW`

and zsh replaced any occurance of OLD in the current directory with NEW, and then cd's into it.

A simple example:
```
lyric > pwd
/usr/local/encap/fvwm-2.2/libexec/fvwm/2.2
lyric > cd 2.2 2.0.46
/usr/local/encap/fvwm-2.0.46/libexec/fvwm/2.0.46
lyric > 
```
-----
`dirs -v`
Direct access to paths: `~1`, `~2` ...
`popd`, `pushd`, or reset `dirs ~/work/proj1 ~/work/proj2 ~/work/proj3`
