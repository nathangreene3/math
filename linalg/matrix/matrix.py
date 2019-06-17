def chain_matrix(matrices):
    # https://medium.com/future-vision/dynamic-programming-deep-dive-chain-matrix-multiplication-a3b8e207b201
    # Ideally do error checking to make sure the columns of
    # each matrix match the rows of the next matrix.

    def cols(i): return matrices[i][1]
    def rows(i): return matrices[i][0]

    n = len(matrices)
    f = {}  # cached values of the recurrence relation

    for l in range(1, n + 1):
        for start in range(0, n - l + 1):
            # Base case
            if l == 1:
                f[(start, start)] = 0
                continue

            # Recursive case
            end = start + l - 1
            f[(start, end)] = min(
                f[(start, mid)] +
                f[(mid + 1, end)] +
                rows(start) * cols(mid) * cols(end)
                for mid in range(start, end)  # end is exclusive
            )
    
    return f[(0, n - 1)]

print(chain_matrix([(2, 10), (10, 3), (3, 8)]))