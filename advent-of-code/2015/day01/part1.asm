section .data
    BUF_SIZE equ 4096

section .bss
    input_buffer resb BUF_SIZE ; Buffer to hold chunks of puzzle input
    output_buffer resb 20      ; Buffer for the ASCII representation

section .text
global _start

_start:
    ; rbx holds the floor value (accumulator across all reads)
    xor rbx, rbx       ; rbx = floor = 0
    
read_loop:
    ; --- 1. Read input chunk from stdin (sys_read = 0) ---
    ; rax = 0 (sys_read), rdi = 0 (stdin), rsi = buffer, rdx = size
    mov rax, 0
    mov rdi, 0
    mov rsi, input_buffer   ; Start reading into the beginning of the buffer
    mov rdx, BUF_SIZE       ; Read up to 4096 bytes
    syscall
    
    ; rax now holds the number of bytes read
    
    ; --- 2. Check for End of File (EOF) or Error ---
    cmp rax, 0
    jle convert_result      ; If rax <= 0 (EOF or error), exit loop and convert result
    
    ; Store the number of bytes read (RAX) into R12 for processing limit
    mov r12, rax 
    
    ; --- 3. Process the current chunk ---
    mov rsi, input_buffer ; rsi = pointer to start of current chunk
    mov r13, 0           ; r13 = index counter (i) for this chunk
    
process_chunk_loop:
    ; Loop termination check: Stop when we've processed all valid bytes (r13 >= r12)
    cmp r13, r12
    jge read_loop           ; Finished this chunk, go read the next one
    
    ; 4. Load character
    movzx r10, byte [rsi] ; Load character into r10
    
    ; 5. Check and process
    cmp r10, '('           
    je increment_floor
    
    cmp r10, ')'           
    je decrement_floor
    
    ; If character is neither, ignore it and continue
    jmp continue_chunk
    
increment_floor:
    inc rbx                ; floor++
    jmp continue_chunk
    
decrement_floor:
    dec rbx                ; floor--
    
continue_chunk:
    inc rsi                ; Move pointer to next character
    inc r13                ; Increment the processed count
    jmp process_chunk_loop ; Continue processing this chunk

    ; --- 6. Convert Integer to ASCII String (itoa) ---
convert_result:
    ; rsi: Temporary pointer that moves backward
    ; r10: The length counter (digits + sign + newline)
    ; r11: The fixed end address
    
    mov rsi, output_buffer + 19 ; Start writing from the end of the temp buffer (index 19)
    mov byte [rsi], 0xA         ; Place newline at the very end
    mov r10, 1                  ; r10 = current length (just the newline)
    dec rsi                     ; Move pointer back

    mov rax, rbx                ; rax = floor value (the number to convert)
    mov r8, 0                   ; r8 is a boolean flag (0 = positive, 1 = negative)
    
    ; Handle negative sign
    cmp rax, 0
    jge convert_abs_value
    neg rax                     ; rax = abs(floor)
    mov r8, 1                   ; Set negative flag
    
convert_abs_value:
    mov r9, 10                  ; r9 = divisor (10)
    
convert_loop:
    xor rdx, rdx                ; Clear rdx for division
    div r9                      ; rax = rax / 10, rdx = rax % 10 (remainder)
    
    add dl, '0'                 ; Convert remainder to ASCII
    mov byte [rsi], dl          ; Store digit in buffer
    dec rsi                     ; Move pointer back
    inc r10                     ; Increment total length
    
    cmp rax, 0                  ; Check if quotient is zero
    jnz convert_loop            ; If not zero, keep dividing

    ; Add back the negative sign if floor was negative
    cmp r8, 1                   ; Check negative flag
    jnz final_setup             ; If not negative, skip
    
    mov byte [rsi], '-'         ; Write the sign at the front
    inc r10                     ; Increment total length
    
final_setup:
    ; Calculate the correct start address: (output_buffer + 20) - total_length
    mov r11, output_buffer + 20 ; R11 = Fixed address just past the end of the buffer (20)
    sub r11, r10                ; R11 is now the exact starting address
    
    ; --- 7. Write result to stdout (sys_write = 1) ---
    mov rax, 1
    mov rdi, 1
    mov rsi, r11                ; RSI = final start address
    mov rdx, r10                ; RDX = final length
    syscall
    
    ; --- 8. Exit (sys_exit = 60) ---
    mov rax, 60
    xor rdi, rdi                ; rdi = 0 (exit status)
    syscall