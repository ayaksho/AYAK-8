# AYAK-8

Little assembler-interpretator, writen on Go

## Usage
In AYAK-8...
 - you can use 3 block's (ax, bx, cx).
 - Using ```db``` you can store variable's.
 - Using ```int``` you can interrupt with devices.
 - ```add``` & ```addx``` used for increasing value by one, or by custom value (can be used ONLY for integers).
 - And other stuff!

## Examples
Printing "Hello World!" message:
```
main:
  db msg, "Hello World!\n"

  mov ax, $msg
  mov bx, 0

  int 0x10
  hlt
```

And that's all... (in the future, there is gonne be MORE function's, stay tune!)
