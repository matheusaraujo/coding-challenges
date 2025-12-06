
(in-package :cl-user)

(defun part2 (puzzle-input)
  "Calculates the sum of matching elements in RIGHT, weighted by frequency in LEFT."
  (multiple-value-bind (left right)
      (aoc-helpers:parse-input puzzle-input)
    (let ((count-map (make-hash-table :test 'eql)) (result 0))
      (dolist (x left) (incf (gethash x count-map 0)))
      (dolist (item right)
        (let ((count-value (gethash item count-map)))
          (when count-value (incf result (* item count-value)))))
      result)))
