# Vim Learn

## Vim Modes

- Normal/Command Mode
- Insert Mode
- Line Mode

## Navigating 浏览导航

要向下移动一行，请按`j`。

要上移一行，按`k`。

要移到右侧，请按`l`。

要移到左侧，请按`h`。

也可以按住导航键，使其重复。

要一直移动到文件顶部，按住`k`。

要一直移动到文件底部，按住`j`。

要在文件中前进，使用Ctrl-f。按住控制键并按下`f`。

Ctrl-f与`向下翻页`操作相同。

要在文件中向后移动，使用Ctrl-b。按住控制键并按`b`。

Ctrl-b与`翻页`操作相同。

要通过单词前进，使用`w`。通过使用空格作为单词向前移动
字边界，使用`W`。

要按字移回，使用`b`。使用空格逐字移动背单词作为词边界，使用`B`。

要转到文件的开头，输入`1gg`或`gg`。

要移至文件的最后一行，输入`$ G`或`G`。

要转到特定的行号，使用`<LINE_NUMBER> gg`或`<LINE_NUMBER> G`。对于例如，要去第27行，键入`27gg`或`27G`。

可以使用线路模式去特定的行。例如，要移至第32行，可以输入`：32 <ENTER>`。要转到最后一行，使用`：$ <ENTER>`。

### 打开标尺

`:set ruler` 和 `:set noruler`

或者 `:set ruler!`

### operation(motion}

    dw

    d=The delete operation

    w=The word motion

### [count]operation{motion}

    5dw

    5=The count/how many times to repeat.

    dw=The command(delete word).

### [count]operation[count](motion}

    3w=Repeat word motion 3 times.

    d3w=Delete the 3w motion.

    2d3w=Delete the 3w motion 2 times.

### dot commond

`.` 重复上次操作