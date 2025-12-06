(defun part2 (puzzle-input)
  "Finds the 1-based index of the character that first reaches floor -1. Returns the integer result."
  (let* ((input-string (car puzzle-input)) (floor 0))
    (loop for char across input-string
          for index from 1
          do (cond ((char= char #\() (incf floor))
                   ((char= char #\)) (decf floor))) (when (= floor -1)
                                                      (return-from part2
                                                        index)))
    0))
