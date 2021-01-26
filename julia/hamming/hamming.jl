"Your optional docstring here"
function distance(a, b)
    if length(a) !== length(b)
        throw(ArgumentError("Unequal lengthb"))
    end

    ham = 0

    for i in range(1, length=length(a))
        char_x = a[i]
        char_y = b[i]
        inc = char_x === char_y ? 0 : 1
        ham += inc
    end

    ham
end
