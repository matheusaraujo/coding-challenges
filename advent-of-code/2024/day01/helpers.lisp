
(require 'asdf)

(defpackage :aoc-helpers
  (:use :cl :uiop)
  (:export :parse-input))

(in-package :aoc-helpers)

(defun parse-integer-safe (str)
  "Converts a string to an integer, returning NIL on failure."
  (if (stringp str)
      (handler-case (parse-integer str :junk-allowed nil) (error nil nil))
      nil))

(defun parse-input (puzzle-input)
  "Parses the input lines into two sorted lists of integers (LEFT and RIGHT)."
  (let ((input-string (car puzzle-input)) (left 'nil) (right 'nil))
    (loop for line in (split-string input-string :separator '(#\Newline))
          when (not (string= line ""))
          do (let* ((all-parts (split-string line :separator '(#\ )))
                    (non-empty-parts
                     (remove-if (lambda (s) (string= s "")) all-parts)))
               (when (= (length non-empty-parts) 2)
                 (let ((val1 (parse-integer-safe (first non-empty-parts)))
                       (val2 (parse-integer-safe (second non-empty-parts))))
                   (when (and val1 val2)
                     (push val1 left)
                     (push val2 right))))))
    (let ((sorted-left (sort left #'<)) (sorted-right (sort right #'<)))
      (values sorted-left sorted-right))))
