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

-----
__brace expansion__
```
> touch logfile.9908{01..31}.tmp
> ls logfile*
logfile.990301.tmp  logfile.990309.tmp  logfile.990317.tmp ogfile.990325.tmp
logfile.990302.tmp  logfile.990310.tmp  logfile.990318.tmp ogfile.990326.tmp
logfile.990303.tmp  logfile.990311.tmp  logfile.990319.tmp ogfile.990327.tmp
logfile.990304.tmp  logfile.990312.tmp  logfile.990320.tmp ogfile.990328.tmp
logfile.990305.tmp  logfile.990313.tmp  logfile.990321.tmp ogfile.990329.tmp
logfile.990306.tmp  logfile.990314.tmp  logfile.990322.tmp ogfile.990330.tmp
logfile.990307.tmp  logfile.990315.tmp  logfile.990323.tmp ogfile.990331.tmp
logfile.990308.tmp  logfile.990316.tmp  logfile.990324.tmp
```
---
__Tab completion__

`compctl [ -CDT ] options [command ...]`

* Completion definitions start with the command compctl on the command line.
* Completion definitions end with the names of command(s) you'd like to complete.
* The list of options specify all types of items you'd like eligible for completion. (Variables, options, filenames, jobs, etc.)

