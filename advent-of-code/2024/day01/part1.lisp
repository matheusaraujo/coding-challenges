
(in-package :cl-user)

(defun part1 (puzzle-input)
  "Calculates the sum of the absolute differences between sorted left and right lists."
  (multiple-value-bind (left right)
      (aoc-helpers:parse-input puzzle-input)
    (let ((sum 0))
      (loop for l in left
            for r in right
            do (incf sum (abs (- l r))))
      sum)))
