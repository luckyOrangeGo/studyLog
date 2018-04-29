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

## 获得帮助

`:help` 获取帮助

`:q<ENTER>` 退出帮助

`:help <COMMAND>` = 特殊查找该命令的帮助

例如 `:help dd<ENTER>`

`:help {subject}` = 查找特殊主题的帮助

例如 `:help count`

`:help :help<ENTER>` = 获取`:`键入类型的帮助

`:help` = `:h`

在帮助文档与书写文档之间切换 `CTRL-W CTRL-W (双击W)`

`CTRL-D` 显示当前`:h []` 中所有可能满足的指令

`Tab` 可以直接选择下一条可能的指令

`Shift-Tab`选择上一条可能的指令

`:set nowildmenu` 取消指令提示行

`CTRL-O` 返回上一个help

`CTRL-I` 前进到刚才的help

`CTRL-]` 下一个单词

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

### cut-copy-paste

cut-copy-paste = delete-yank-put.

Registers are storage locations.

`""`contains last operated on text.

Numbered registers:`"0`...`"9`.

Named registers:`"a`....`"z`.

`:reg [register(s)]`

### UNDO 撤销 REDO 重做

undo = `u`, redo = `Ctrl-R`

### Insert 插入

`Shift-I` = 将光标移动到该行**最前方**并进入Insert Mode

`Shift-A` = 将光标移动到该行**最后方**并进入Insert Mode

`o` = 将光标移动到该行**下方新建一行**并进入Insert Mode

`Shift-O` = 将光标移动到该行**上方新建一行**并进入Insert Mode

`{NUMBER} + i` = 插入字符***NUMBER 遍**，按下 Escape 键后生效

`{NUMBER} + o` = 新建一行插入字符后***NUMBER 行重复**，按下 Escape 键后生效

### Replace 替换

`Shift-R` 进入光标位置处的Replace mode

`r + {另一个字符}` 替换光标下的**一个字符**

`c + ({NUMBER}) + w` = **捕获光标后的该({NUMBER})个单词**并进入Insert Mode，被捕获的部分进入寄存器中，可以通过`:reg` 查看或使用`p`粘贴。
也可以**在该命令前面使用`"a`等指定的寄存器存入这些字符**。

`c + $` 或者 `Shift-C`= 捕获该行*光标后*的所有字符

`c + c` = 捕获*该行*所有字符

### Upper 大小写转换

`~` = **切换**光标下的*该字符*为大写（小写）

`g~$` = 将*光标后的*该行所有字符全部**大小写切换**

`g~~` = 将*该行*所有字符全部**大小写切换**

`gU + w` = 将*word*转换为**大写**

`gUU` = 将*该行*转换为**大写**

`gu + w` = 将*word*转换为**小写**

`guu` = 将*该行*转换为**小写**

### Join 连接两行

`Shift-J` = 连接该行和下一行**（自动添加空格在其之间）**

`g + J` = 连接该行和下一行**（连接处无空格）**

`{NUM} + J` = 连接{NUM}行

### 查找

`/ + {Word}` 查找光标后的内容至该{Word} 处

`//` 重复上次查找

