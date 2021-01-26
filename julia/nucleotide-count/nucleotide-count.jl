"""
    count_nucleotides(strand)

The frequency of each nucleotide within `strand` as a dictionary.

Invalid strands raise a `DomainError`.

"""
function count_nucleotides(strand)
    result = Dict('A' => 0, 'C' => 0, 'G' => 0, 'T' => 0)

    for char in strand
        if char ∉ valid_chars
            throw(DomainError(char, "Not a valid char"))
        end

        result[char] += 1
    end

    result
end
