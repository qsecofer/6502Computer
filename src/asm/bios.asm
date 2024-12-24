PROMPT          .EQ     "\"
DSP             .EQ     $6000

    .org $8000

reset:
    lda #$80
loop:
    nop
    ror A
    jmp loop

echo:
    BIT DSP
    BMI echo
    STA DSP
    RTS
    
irq:
    lda #$55
    rti

nmi:
    rti

    .org $fffa
    .word nmi
    .word reset
    .word irq  