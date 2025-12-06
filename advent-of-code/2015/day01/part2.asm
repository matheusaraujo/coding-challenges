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
    mov r12, 0         ; r12 = Global 0-based index (i), starts at 0
    
read_loop:
    ; --- 1. Read input chunk from stdin (sys_read = 0) ---
    mov rax, 0
    mov rdi, 0
    mov rsi, input_buffer   
    mov rdx, BUF_SIZE       
    syscall
    
    ; rax now holds the number of bytes read
    
    ; --- 2. Check for End of File (EOF) or Error ---
    cmp rax, 0
    jle not_found           ; If rax <= 0 (EOF or error), exit loop and return 0
    
    ; Store the number of bytes read (RAX) into R13 for processing limit
    mov r13, rax 
    
    ; --- 3. Process the current chunk ---
    mov rsi, input_buffer ; rsi = pointer to start of current chunk
    mov r14, 0           ; r14 = Chunk 0-based index (j)
    
process_chunk_loop:
    ; Loop termination check: Stop when we've processed all valid bytes (r14 >= r13)
    cmp r14, r13
    jge read_loop           ; Finished this chunk, go read the next one
    
    ; Load character
    movzx r10, byte [rsi] 
    
    ; Check and process
    cmp r10, '('           
    je increment_floor
    
    cmp r10, ')'           
    je decrement_floor
    
    ; If character is neither, ignore it and continue
    jmp next_char
    
increment_floor:
    inc rbx                ; floor++
    jmp check_basement
    
decrement_floor:
    dec rbx                ; floor--
    
check_basement:
    ; Check floor *AFTER* processing the character
    cmp rbx, -1             
    je found_basement       ; If -1, we found the basement!
    
next_char:
    inc rsi                ; Move pointer to next character
    inc r12                ; Increment Global 0-based index (i)
    inc r14                ; Increment Chunk 0-based index (j)
    jmp process_chunk_loop ; Continue processing this chunk

    ; --- 4. Basement Found ---
found_basement:
    ; R12 is the 0-based index of the character that caused the drop.
    inc r12                 ; Result is the 1-based index (i + 1)
    mov rax, r12            ; RAX holds the result
    jmp convert_value

    ; --- 5. Not Found (Return 0) ---
not_found:
    mov rax, 0              ; RAX holds the result
    
    ; --- 6. Convert Integer to ASCII String (itoa) ---
convert_value:
    ; rsi: Temporary pointer that moves backward
    ; r10: The length counter (digits + sign + newline)
    ; r11: The fixed end address
    
    mov rsi, output_buffer + 19 
    mov byte [rsi], 0xA         ; Place newline at the very end
    mov r10, 1                  ; r10 = current length (just the newline)
    dec rsi                     ; Move pointer back

    ; RAX holds the result value to convert
    mov r8, 0                   ; r8 is a boolean flag (0 = positive, 1 = negative)
    
    ; Note: The basement index (i+1) is always positive, so we skip sign handling.
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

final_setup:
    ; Calculate the correct start address: (output_buffer + 20) - total_length
    mov r11, output_buffer + 20 
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