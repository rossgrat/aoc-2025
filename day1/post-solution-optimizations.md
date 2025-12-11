# Day 1: Secret Entrance - Solution Methods

## Problem Summary

A dial with positions 0-99 starts at position 50. Given a sequence of rotations (L/R with a distance), solve:

- **Part 1**: Count how many times the dial *ends* at 0 after a rotation
- **Part 2**: Count how many times the dial *passes through or ends* at 0 during all rotations

---

## Method 1: Looping (Simulation)

The most straightforward approach - simulate each click of the dial.

### Part 1

```go
func rotateLeft(start int, val int) int {
    for range val {
        start = start - 1
        if start == -1 {
            start = 99
        }
    }
    return start
}

func rotateRight(start int, val int) int {
    for range val {
        start = start + 1
        if start == 100 {
            start = 0
        }
    }
    return start
}

// After each rotation, check if start == 0
```

### Part 2

```go
func rotateLeft(start int, val int, zeroes *int) int {
    for range val {
        start = start - 1
        if start == -1 {
            start = 99
        }
        if start == 0 {
            *zeroes = *zeroes + 1
        }
    }
    return start
}
```

### Complexity

- **Time**: O(total clicks) - could be very large if rotation values are big
- **Space**: O(1)

### Example

Starting at 50, rotate L68:
- Click through 49, 48, 47... 0 (count!), 99, 98... 82
- Result: position 82, crossed 0 once

### Trade-offs

| Pros | Cons |
|------|------|
| Easy to understand and debug | Slow for large rotation values |
| Directly models the problem | R1000000 would take 1M iterations |

---

## Method 2: Modular Arithmetic (Division & Remainder)

Use math to compute the final position directly, then analyze wrap-arounds.

### Part 1

```go
func universalModulo(start int, val int, total int) int {
    return ((((start + val) % total) + total) % total)
}

// For L rotations, val is negative
// For R rotations, val is positive
newPos := universalModulo(start, delta, 100)
if newPos == 0 {
    zeroes++
}
```

### Part 2

```go
newStart := universalModulo(start, delta, 100)

// Count full rotations (each passes through 0 once)
rotations := val / 100
zeroes += rotations

// Check if we ended at 0
if newStart == 0 {
    zeroes++
} else if dir == "L" && newStart > start && start != 0 {
    // Wrapped around going left
    zeroes++
} else if dir == "R" && newStart < start && start != 0 {
    // Wrapped around going right
    zeroes++
}
```

### Complexity

- **Time**: O(n) where n = number of instructions
- **Space**: O(1)

### Example

Starting at 50, rotate L68:
- `newStart = ((50 - 68) % 100 + 100) % 100 = 82`
- Full rotations: `68 / 100 = 0`
- Did we wrap? `82 > 50` and going left, so yes!
- Total zeros crossed: 1

### Trade-offs

| Pros | Cons |
|------|------|
| O(1) per instruction regardless of rotation size | Wrap-around logic requires careful thought |
| Handles R1000000 instantly | Multiple conditional branches |

---

## Method 3: Closed-Form Distance Formula

Unify all cases into a single formula based on "distance to first zero".

### Core Insight

When rotating from position `p` by `v` clicks:
1. Calculate distance to first encounter with 0
2. If `v >= distance`, we hit 0 at least once
3. After that, we hit 0 again every 100 clicks

### Implementation

```go
func zerosHit(pos int, val int, dir string) int {
    var firstZero int

    if dir == "L" {
        // Going left: first zero is at distance `pos` (or 100 if pos==0)
        firstZero = pos
        if pos == 0 {
            firstZero = 100
        }
    } else {
        // Going right: first zero is at distance `100-pos` (or 100 if pos==0)
        firstZero = 100 - pos
        if pos == 0 {
            firstZero = 100
        }
    }

    if val >= firstZero {
        return 1 + (val-firstZero)/100
    }
    return 0
}
```

### Complexity

- **Time**: O(n) where n = number of instructions
- **Space**: O(1)

### Example Walkthrough

| Start | Rotation | Direction | First Zero Distance | Val >= Distance? | Zeros |
|-------|----------|-----------|---------------------|------------------|-------|
| 50 | L68 | Left | 50 | 68 >= 50 = Yes | 1 + (68-50)/100 = 1 |
| 82 | L30 | Left | 82 | 30 >= 82 = No | 0 |
| 52 | R48 | Right | 48 | 48 >= 48 = Yes | 1 + (48-48)/100 = 1 |
| 0 | L1 | Left | 100 | 1 >= 100 = No | 0 |
| 99 | L99 | Left | 99 | 99 >= 99 = Yes | 1 + (99-99)/100 = 1 |
| 50 | R1000 | Right | 50 | 1000 >= 50 = Yes | 1 + (1000-50)/100 = 10 |

### Trade-offs

| Pros | Cons |
|------|------|
| Single unified formula | Less intuitive than conditional checks |
| No branching for wrap-around detection | Requires understanding the math |
| Elegant mathematical model | Edge case at pos==0 still needs handling |

---

## Method 4: Cumulative Sum Perspective

Reframe the problem in terms of prefix sums.

### Core Insight

Position after k rotations = `(50 + sum of first k deltas) mod 100`

For Part 1, we want `position == 0`, which means:
```
(50 + cumulative_delta) mod 100 == 0
cumulative_delta mod 100 == 50
```

### Implementation

```go
func solvePart1(rotations []Rotation) int {
    cumSum := 0
    zeros := 0

    for _, r := range rotations {
        delta := r.val
        if r.dir == "L" {
            delta = -delta
        }
        cumSum += delta

        // Check if cumulative sum mod 100 equals 50
        if ((cumSum % 100) + 100) % 100 == 50 {
            zeros++
        }
    }
    return zeros
}
```

### Complexity

- **Time**: O(n)
- **Space**: O(1)

### Example

| Step | Rotation | Delta | Cumulative Sum | Cum % 100 | Position | At Zero? |
|------|----------|-------|----------------|-----------|----------|----------|
| 0 | - | - | 0 | 0 | 50 | No |
| 1 | L68 | -68 | -68 | 32 | 82 | No |
| 2 | L30 | -30 | -98 | 2 | 52 | No |
| 3 | R48 | +48 | -50 | 50 | 0 | Yes! |
| 4 | L5 | -5 | -55 | 45 | 95 | No |

### Trade-offs

| Pros | Cons |
|------|------|
| Elegant mathematical reframing | Only works directly for Part 1 |
| Shows the underlying structure | Part 2 still needs interval analysis |

---

## Method 5: Interval/Arc Counting

Think geometrically - each rotation sweeps an arc on the circular dial.

### Core Insight

A rotation from position `p` by `v` clicks sweeps an arc. Count how many times the arc contains position 0.

- **Arc definition**: All positions visited (exclusive of start, inclusive of end)
- **Zeros in arc**: `floor(v/100) + (1 if 0 is in the "remainder arc")`

### Visual Example

```
Dial positions (circular):

        0
    99     1
  98         2
 .             .
 .             .
95               5
  .           .
   50       49
      ...

Starting at 50, L68:
Arc sweeps: 49 → 48 → ... → 0 → 99 → ... → 82
Contains 0? Yes, once.
```

### Implementation

```go
func arcContainsZero(start, end int, dir string) bool {
    if dir == "L" {
        // Going left: arc from start-1 down to end
        // Contains 0 if we cross from positive to negative (wrap)
        return end > start || end == 0
    } else {
        // Going right: arc from start+1 up to end
        // Contains 0 if we cross from 99 to 0 (wrap)
        return end < start || end == 0
    }
}

func zerosInRotation(start, val int, dir string) int {
    end := computeEnd(start, val, dir)
    fullRotations := val / 100
    remainder := val % 100

    zeros := fullRotations

    // Check if remainder arc contains 0
    remainderEnd := computeEnd(start, remainder, dir)
    if remainder > 0 && arcContainsZero(start, remainderEnd, dir) {
        zeros++
    }

    return zeros
}
```

### Trade-offs

| Pros | Cons |
|------|------|
| Intuitive geometric interpretation | More code than formula approach |
| Easy to visualize | Same complexity as other O(n) methods |

---

## Comparison Summary

| Method | Time Complexity | Best For | Readability |
|--------|-----------------|----------|-------------|
| Looping | O(total clicks) | Small inputs, debugging | Very High |
| Modular Arithmetic | O(n) | Production code | High |
| Closed-Form Formula | O(n) | Mathematical elegance | Medium |
| Cumulative Sum | O(n) | Part 1 specifically | High |
| Interval Counting | O(n) | Geometric intuition | Medium |

---

## Conclusion

All non-looping methods achieve **O(n)** time complexity, which is optimal since we must read each instruction. The choice between them is primarily about:

1. **Readability** - Modular arithmetic with explicit conditionals is easiest to understand
2. **Elegance** - Closed-form formula is mathematically cleanest
3. **Debugging** - Looping is best for verifying correctness on small inputs

For this problem size (likely ~1000 instructions with moderate rotation values), any O(n) approach will be effectively instant.
