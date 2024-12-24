;-------------------------------------------------------------------------
;
;  The WOZ Monitor for the Apple 1
;  Written by Steve Wozniak 1976
;
;-------------------------------------------------------------------------


;-------------------------------------------------------------------------
;  Memory declaration
;-------------------------------------------------------------------------

XAML            .EQ     $24             
XAMH            .EQ     $25             
STL             .EQ     $26             
STH             .EQ     $27             
L               .EQ     $28             
H               .EQ     $29             
YSAV            .EQ     $2A             
MODE            .EQ     $2B             
IN              .EQ     $0200
KBD             .EQ     $D010           
KBDCR           .EQ     $D011           
DSP             .EQ     $6000
DSPCR           .EQ     $6001           

; KBD b7..b0 are inputs, b6..b0 is ASCII input, b7 is constant high
;     Programmed to respond to low to high KBD strobe
; DSP b6..b0 are outputs, b7 is input
;     CB2 goes low when data is written, returns high when CB1 goes high
; Interrupts are enabled, though not used. KBD can be jumpered to IRQ,
; whereas DSP can be jumpered to NMI.

;-------------------------------------------------------------------------
;  Constants
;-------------------------------------------------------------------------

BS              .EQ     $DF
CR              .EQ     $8D
ESC             .EQ     $9B
PROMPT          .EQ     "\"

;-------------------------------------------------------------------------
;  Let's get started
;
;  Remark the RESET routine is only to be entered by asserting the RESET
;  line of the system. This ensures that the data direction registers
;  are selected.
;-------------------------------------------------------------------------

    .org $8000

RESET           CLD                     
                CLI
                LDY     #%01111111     
                STY     DSP             
                LDA     #%10100111     
                STA     KBDCR           
                STA     DSPCR           

; Program falls through to the GETLINE routine to save some program bytes
; Please note that Y still holds $7F, which will cause an automatic Escape

;-------------------------------------------------------------------------
; The GETLINE process
;-------------------------------------------------------------------------

NOTCR           CMP     #BS             
                BEQ     BACKSPACE       
                CMP     #ESC            
                BEQ     ESCAPE          
                INY                     
                BPL     NEXTCHAR        
ESCAPE          LDA     #PROMPT         
                JSR     ECHO            
GETLINE         LDA     #CR             
                JSR     ECHO
                LDY     #0+1            
BACKSPACE       DEY                     
                BMI     GETLINE         
NEXTCHAR        LDA     KBDCR           
                BPL     NEXTCHAR        
                LDA     KBD             
                STA     IN,Y            
                JSR     ECHO            
                CMP     #CR
                BNE     NOTCR           

; Line received, now let's parse it
                LDY     #-1             
                LDA     #0              
                TAX                     
SETSTOR         ASL                     
SETMODE         STA     MODE            
BLSKIP          INY                     
NEXTITEM        LDA     IN,Y            
                CMP     #CR
                BEQ     GETLINE         
                CMP     #"."
                BCC     BLSKIP          
                BEQ     SETMODE         
                CMP     #":"
                BEQ     SETSTOR         
                CMP     #"R"
                BEQ     RUN             
                STX     L               
                STX     H
                STY     YSAV            
; Here we're trying to parse a new hex value

NEXTHEX         LDA     IN,Y            
                EOR     #$B0            
                CMP     #9+1            
                BCC     DIG             
                ADC     #$88            
                CMP     #$FA            
                BCC     NOTHEX          
DIG             ASL
                ASL                     
                ASL
                ASL
                LDX     #4              
HEXSHIFT        ASL                     
                ROL     L               
                ROL     H               
                DEX                     
                BNE     HEXSHIFT        
                INY                     
                BNE     NEXTHEX         
NOTHEX          CPY     YSAV            
                BEQ     ESCAPE          
                BIT     MODE            
                BVC     NOTSTOR         

; STOR mode, save LSD of new hex byte

                LDA     L               
                STA     (STL,X)         
                INC     STL             
                BNE     NEXTITEM        
                INC     STH             
TONEXTITEM      JMP     NEXTITEM        

;-------------------------------------------------------------------------
;  RUN user's program from last opened location
;-------------------------------------------------------------------------

RUN             JMP     (XAML)          

;-------------------------------------------------------------------------
;  We're not in Store mode
;-------------------------------------------------------------------------

NOTSTOR         BMI     XAMNEXT         

; We're in XAM mode now

                LDX     #2              
SETADR          LDA     L-1,X           
                STA     STL-1,X         
                STA     XAML-1,X        
                DEX                     
                BNE     SETADR          

; Print address and data from this address, fall through next BNE.

NXTPRNT         BNE     PRDATA          
                LDA     #CR             
                JSR     ECHO
                LDA     XAMH            
                JSR     PRBYTE
                LDA     XAML            
                JSR     PRBYTE
                LDA     #":"            
                JSR     ECHO

PRDATA          LDA     #" "            
                JSR     ECHO
                LDA     (XAML,X)        
                JSR     PRBYTE          
XAMNEXT         STX     MODE            
                LDA     XAML            
                CMP     L
                LDA     XAMH
                SBC     H
                BCS     TONEXTITEM      
                INC     XAML            
                BNE     MOD8CHK         
                INC     XAMH
MOD8CHK         LDA     XAML            
                AND     #%0000011
                BPL     NXTPRNT         

;-------------------------------------------------------------------------
;  Subroutine to print a byte in A in hex form (destructive)
;-------------------------------------------------------------------------

PRBYTE          PHA                     
                LSR
                LSR
                LSR                     
                LSR
                JSR     PRHEX           
                PLA                     

; Fall through to print hex routine

;-------------------------------------------------------------------------
;  Subroutine to print a hexadecimal digit
;-------------------------------------------------------------------------

PRHEX           AND     #%00001111     
                ORA     #"0"            
                CMP     #"9"+1          
                BCC     ECHO            
                ADC     #6              

; Fall through to print routine

;-------------------------------------------------------------------------
;  Subroutine to print a character to the terminal
;-------------------------------------------------------------------------

ECHO            BIT     DSP             
                BMI     ECHO            
                STA     DSP             
                RTS


;-------------------------------------------------------------------------

    .org $fffc
    .word RESET