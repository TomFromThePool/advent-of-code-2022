using DelimitedFiles

function p(pair)
    return [parse(Int64, x) for x in eachsplit(pair, '-')]
end

function lContainsR(l, r)
    return minimum(l) <= minimum(r) && maximum(l) >= maximum(r)
end

function lOverlapsR(l, r)
    return lContainsR(l, r) || maximum(l) - minimum(r) >= 0 && maximum(r) >= minimum(l)
end

function compareContains(l, r)
    return lContainsR(l, r) || lContainsR(r, l) 
end

function compareOverlaps(l, r)
    return lOverlapsR(l, r)
end

input = readdlm("input.txt", ',')
g = eachrow(input)
contains = [compareContains(p(first(x)), p(last(x))) for x in g]
overlap = [compareOverlaps(p(first(x)), p(last(x))) for x in g]

println(count(contains[contains .== true]))
println(count(overlap[overlap .== true]))