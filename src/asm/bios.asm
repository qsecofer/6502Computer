PROMPT          .EQ     "\"
DSP             .EQ     $6000

    .org $8000

reset:
    lda #$80
    pha

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
    nop
    nop
    nop
    rti

nmi:
    lda #$EA
    nop
    nop
    rti

    .org $fffa
    .word nmi
    .word reset
    .word irq  