PROMPT          .EQ     "\"
DSP             .EQ     $6000
GDSP            .EQ     $7000

    .org $8000

reset:
    lda #"H"
    sta DSP
    sta GDSP

    lda #"E"
    sta DSP
    sta GDSP

    jmp reset

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