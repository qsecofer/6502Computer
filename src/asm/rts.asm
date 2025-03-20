PORTB = $6000

  .org $8000

reset:
  nop
  jsr lcd_instruction
  nop
  
loop:
  jmp reset

lcd_instruction:
  sta PORTB
  lda #0        
  rts

  .org $fffc
  .word reset
  .word $0000
