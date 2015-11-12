```
$ awk '{print $0}' file #  Print the all column.
$ awk '{print $3}' file #  Print the 3rd column.
$ awk '{print $1 $3}' file #  Print the 1st and the 3rd columns.
```

`tee` to print stdout and log to file.
```
cmd xxx | tee file
```

```
2>&1 # redirect stderr to stdout
>/dev/null 2>&1 # and redirect to black hole
```
### Sum
```
cat numbers.log | python -c "import sys; print sum(int(l) for l in sys.stdin)"
cat numbers.log | ruby -e 'puts STDIN.read.split("\n").map(&:to_i).reduce(:+)'
cat numbers.log | paste -s -d+ infile | bc
```
### Regex match
```
cat $1 | awk '{ gsub("\t", "  ") ; print $0 }'
```

```
import re
import sys
import argparse

parser = argparse.ArgumentParser(description='print matched by regex')
parser.add_argument('regex', type=str, help='regex pattern')
args = parser.parse_args()

for line in sys.stdin:
    matched = re.match(args.regex, line)
    if matched:
        print matched.group(1)
```
### Bash
```
my_array=(one two three four)
for i in ${my_array[@]}; do
  echo $i
done
```
Find
```
find . -name "abc.jpg"
find . -regextype sed -regex ".*/[a-f0-9\-]\{36\}\.jpg"
```

### Network

List connections
```
lsof -i
```
