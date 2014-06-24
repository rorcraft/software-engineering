The 2010 Census puts populations of 26 largest US metro areas at 18897109, 12828837, 9461105, 6371773, 5965343, 5946800, 5582170, 5564635, 5268860, 4552402, 4335391, 4296250, 4224851, 4192887, 3439809, 3279833, 3095313, 2812896, 2783243, 2710489, 2543482, 2356285, 2226009, 2149127, 2142508, and 2134411.

Can you find a subset of these areas where a total of exactly 100,000,000 people live, assuming the census estimates are exactly right? Provide the answer and code or reasoning used.

Solution: 

```ruby

stats = [18897109, 12828837, 9461105, 6371773, 5965343, 5946800,
         5582170, 5564635, 5268860, 4552402, 4335391, 4296250, 4224851,
         4192887, 3439809, 3279833, 3095313, 2812896, 2783243, 2710489,
         2543482, 2356285, 2226009, 2149127, 2142508, 2134411]
@target = 100_000_000
@combos = []
@sum = 0
@min = 0
@max = 0

# limit search space finding max and min number of combos required to add up to 100_000_000
stats.each_with_index do |n, i|
  @sum += n
  if @sum >= @target
    @min = i + 1
    break
  end
end
# puts @min
@sum = 0
stats.sort.each_with_index do |n, i|
  @sum += n
  if @sum >= @target
    @max = i + 1
    break
  end
end

# search through all possible combos and find ones that adds up to exactly 100_000_000
(@min..@max).each do |i|
  stats.combination(i).map do |combo|
    sum = combo.inject(:+)
    if sum == @target
      puts combo.inspect
    end
  end
end

# Running times:
# ruby 1.9.3
# 48.82s user 0.32s system 99% cpu 49.302 total
# jruby 1.7.11
# 35.83s user 0.47s system 138% cpu 26.254 total
```

Answer:

```
[18897109, 12828837, 9461105, 6371773, 5946800, 5582170, 5268860, 4552402, 4335391, 4296250, 4224851, 3279833, 3095313, 2812896, 2543482, 2226009, 2142508, 2134411]
```
