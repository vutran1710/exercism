"""
    is_leap_year(year)

Return `true` if `year` is a leap year in the gregorian calendar.

"""
function is_leap_year(year)
    is_divisible(x, y) = ===(mod(x, y), 0)
    is_divisible(year, 400) || (is_divisible(year, 4) && !is_divisible(year, 100))
end
