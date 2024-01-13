## Ascii-art-justify Project

## Objectives
The objective of the ascii-art program is to receive a string as an argument and output the string in a graphic representation using ASCII characters. The program should be able to handle input with numbers, letters, spaces, special characters, and line breaks (\n).
A graphic representation using ASCII characters means that the string will be written using ASCII characters to form a visual representation. Each character in the string will be represented by a series of ASCII characters, following a specific graphical template.

## Instructions
- To change the alignment of the output it must be possible to use a flag --align=<type>, in which type can be :
center
left
right
justify
- The representation should be adapted to the terminal size. 
If you reduce the terminal window the graphical representation should be adapted to the terminal size.
Only text that fits the terminal size will be tested.

## Usage
To run the program, use the following command on Git Bash terminal:

```
Usage: go run . [OPTION] [STRING] [BANNER]

```

Example: 
```
go run . --align=right something standard

```

To run the Test File:

```
./Test.sh
```

The program should accept other correctly formatted [OPTION] and/or [BANNER].
Additionally, the program must still be able to run with a single [STRING] argument.

## AUTHORS

* Fatema Alawadhi
* Asma Ahmed


## LICENSES
This program developed within the scope of Reboot.