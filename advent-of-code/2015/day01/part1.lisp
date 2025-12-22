
(defun part1 (puzzle-input)
  "Calculates the final floor level. Returns the integer result."
  (let* ((input-string (car puzzle-input)) (floor 0))
    (loop for char across input-string
          do (cond ((char= char #\() (incf floor))
                   ((char= char #\)) (decf floor))))
    floor))
