// by gemini

package main

func part3(puzzleInput []string) any {
	start, bones := ParseGrid(puzzleInput)

	visited := map[Pos]bool{}
	pos := start
	visited[pos] = true

	// Track the bounding box
	minR, maxR := start.R, start.R
	minC, maxC := start.C, start.C

	for b := range bones {
		if b.R < minR {
			minR = b.R
		}
		if b.R > maxR {
			maxR = b.R
		}
		if b.C < minC {
			minC = b.C
		}
		if b.C > maxC {
			maxC = b.C
		}
	}

	dir := 0
	steps := 0

	// Reusable structures for our fast "Outside-In" BFS
	seen := make(map[Pos]int)
	var queue []Pos
	runID := 0

	for {
		// 1. Progress the spiral until we place a valid move
		for {
			next := Add(pos, SpiralDirs[dir])

			if !visited[next] && !bones[next] {
				pos = next
				visited[pos] = true
				steps++

				// Expand bounding box if necessary
				if pos.R < minR {
					minR = pos.R
				}
				if pos.R > maxR {
					maxR = pos.R
				}
				if pos.C < minC {
					minC = pos.C
				}
				if pos.C > maxC {
					maxC = pos.C
				}

				dir = (dir + 1) % len(SpiralDirs)
				break
			}
			dir = (dir + 1) % len(SpiralDirs)
		}

		// 2. Optimization: Did we potentially close a loop?
		// A loop can only close if our new cell touches >1 existing solid cells
		// (the one we just moved from, plus at least one other).
		touchCount := 0
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				adj := Pos{pos.R + dr, pos.C + dc}
				if visited[adj] || bones[adj] {
					touchCount++
				}
			}
		}

		// 3. If we touched multiple things, check for trapped spaces and fill them
		if touchCount > 1 {
			runID++
			queue = queue[:0]

			// Pad the bounding box by 1 to create a guaranteed "outside" perimeter
			padMinR, padMaxR := minR-1, maxR+1
			padMinC, padMaxC := minC-1, maxC+1

			// Add the entire perimeter to the BFS queue
			for r := padMinR; r <= padMaxR; r++ {
				queue = append(queue, Pos{r, padMinC}, Pos{r, padMaxC})
				seen[Pos{r, padMinC}] = runID
				seen[Pos{r, padMaxC}] = runID
			}
			for c := padMinC + 1; c <= padMaxC-1; c++ {
				queue = append(queue, Pos{padMinR, c}, Pos{padMaxR, c})
				seen[Pos{padMinR, c}] = runID
				seen[Pos{padMaxR, c}] = runID
			}

			// Run BFS from the outside in
			head := 0
			for head < len(queue) {
				curr := queue[head]
				head++

				for _, d := range Cardinal {
					next := Add(curr, d)
					// Stay within padded bounds, walk only on empty space
					if next.R >= padMinR && next.R <= padMaxR && next.C >= padMinC && next.C <= padMaxC {
						if seen[next] != runID && !visited[next] && !bones[next] {
							seen[next] = runID
							queue = append(queue, next)
						}
					}
				}
			}

			// 4. Check for trapped cells and surrounded bones
			allBonesSurrounded := true

			for r := minR; r <= maxR; r++ {
				for c := minC; c <= maxC; c++ {
					p := Pos{r, c}

					if bones[p] {
						// A bone is NOT surrounded if any neighboring cell was reached by the outside BFS
						for _, d := range Cardinal {
							if seen[Add(p, d)] == runID {
								allBonesSurrounded = false
							}
						}
					} else if !visited[p] {
						// Empty cell. If the outside BFS didn't see it, it's trapped!
						if seen[p] != runID {
							visited[p] = true // Instantly fill it!
						}
					}
				}
			}

			// If all bones are blocked off from the outside BFS, we are done.
			if allBonesSurrounded {
				return steps
			}
		}
	}
}
