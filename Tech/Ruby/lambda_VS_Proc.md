http://en.wikibooks.org/wiki/Ruby_Programming/Syntax/Method_Calls

> ```return``` inside a block or Proc[1] will return from the method, in which the block or Proc is defined, not from the block itself
```
def func_one
    proc_new = Proc.new {return "123"}
    proc_new.call
    return "456"
end

def func_two
    lambda_new = lambda {return "123"}
    lambda_new.call
    return "456"
end

-----

The result of running func_one is 123
The result of running func_two is 456

```
