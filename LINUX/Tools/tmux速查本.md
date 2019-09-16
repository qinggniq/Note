## 啟動新會話：
```shell
tmux [new -s 會話名 -n 視窗名]
```
## 恢復會話：
```shell
tmux at [-t 會話名]
```
## 列出所有會話：
```shell
tmux ls
```
## 關閉會話：
```shell
tmux kill-session -t 會話名
```
## 關閉所有會話：
```shell
tmux ls | grep : | cut -d. -f1 | awk '{print substr($1, 0, length($1)-1)}' | xargs kill
```
## 在 Tmux 中，按下 Tmux 字首 ctrl+b，然後：
### 會話
```shell
:new<回車>  啟動新會話
s           列出所有會話
$           重新命名當前會話
```
### 視窗 (標籤頁)
```shell
c  建立新視窗
w  列出所有視窗
n  後一個視窗
p  前一個視窗
f  查詢視窗
,  重新命名當前視窗
&  關閉當前視窗
```
### 調整視窗排序
```shell
swap-window -s 3 -t 1  交換 3 號和 1 號視窗
swap-window -t 1       交換當前和 1 號視窗
move-window -t 1       移動當前視窗到 1 號
```
### 窗格（分割視窗）
```shell
%  垂直分割
"  水平分割
o  交換窗格
x  關閉窗格
⍽  左邊這個符號代表空格鍵 - 切換佈局
q 顯示每個窗格是第幾個，當數字出現的時候按數字幾就選中第幾個窗格
{ 與上一個窗格交換位置
} 與下一個窗格交換位置
z 切換窗格最大化/最小化
```
### 同步窗格
這麼做可以切換到想要的視窗，輸入 Tmux 字首和一個冒號撥出命令提示行，然後輸入：
```shell
:setw synchronize-panes
```
你可以指定開或關，否則重複執行命令會在兩者間切換。 這個選項值針對某個視窗有效，不會影響別的會話和視窗。 完事兒之後再次執行命令來關閉。幫助

### 調整窗格尺寸
如果你不喜歡預設佈局，可以重調窗格的尺寸。雖然這很容易實現，但一般不需要這麼幹。這幾個命令用來調整窗格：

PREFIX : resize-pane -D          當前窗格向下擴大 1 格
PREFIX : resize-pane -U          當前窗格向上擴大 1 格
PREFIX : resize-pane -L          當前窗格向左擴大 1 格
PREFIX : resize-pane -R          當前窗格向右擴大 1 格
PREFIX : resize-pane -D 20       當前窗格向下擴大 20 格
PREFIX : resize-pane -t 2 -L 20  編號為 2 的窗格向左擴大 20 格
文字複製模式：
按下**字首 [**進入文字複製模式。可以使用方向鍵在螢幕中移動游標。預設情況下，方向鍵是啟用的。在配置檔案中啟用 Vim 鍵盤佈局來切換視窗、調整窗格大小。Tmux 也支援 Vi 模式。要是想啟用 Vi 模式，只需要把下面這一行新增到 .tmux.conf 中：
```shell
setw -g mode-keys vi
```
啟用這條配置後，就可以使用 h、j、k、l 來移動游標了。


想要退出文字複製模式的話，按下回車鍵就可以了。一次移動一格效率低下，在 Vi 模式啟用的情況下，可以輔助一些別的快捷鍵高效工作。


例如，可以使用 w 鍵逐詞移動，使用 b 鍵逐詞回退。使用 f 鍵加上任意字元跳轉到當前行第一次出現該字元的位置，使用 F 鍵達到相反的效果。
```shell
vi             emacs        功能
^              M-m          反縮排
Escape         C-g          清除選定內容
Enter          M-w          複製選定內容
j              Down         游標下移
h              Left         游標左移
l              Right        游標右移
L                           游標移到尾行
M              M-r          游標移到中間行
H              M-R          游標移到首行
k              Up           游標上移
d              C-u          刪除整行
D              C-k          刪除到行末
$              C-e          移到行尾
:              g            前往指定行
C-d            M-Down       向下滾動半屏
C-u            M-Up         向上滾動半屏
C-f            Page down    下一頁
w              M-f          下一個詞
p              C-y          貼上
C-b            Page up      上一頁
b              M-b          上一個詞
q              Escape       退出
C-Down or J    C-Down       向下翻
C-Up or K      C-Up         向下翻
n              n            繼續搜尋
?              C-r          向前搜尋
/              C-s          向後搜尋
0              C-a          移到行首
Space          C-Space      開始選中
               C-t          字元調序
```
### 雜項：
```shell
d  退出 tmux（tmux 仍在後臺執行）
t  視窗中央顯示一個數字時鐘
?  列出所有快捷鍵
:  命令提示符
```
### 配置選項：
```shell
# 滑鼠支援 - 設定為 on 來啟用滑鼠
* setw -g mode-mouse off
* set -g mouse-select-pane off
* set -g mouse-resize-pane off
* set -g mouse-select-window off

# 設定預設終端模式為 256color
set -g default-terminal "screen-256color"

# 啟用活動警告
setw -g monitor-activity on
set -g visual-activity on

# 居中視窗列表
set -g status-justify centre

# 最大化/恢復窗格
unbind Up bind Up new-window -d -n tmp \; swap-pane -s tmp.1 \; select-window -t tmp
unbind Down
bind Down last-window \; swap-pane -s tmp.1 \; kill-window -t tmp
```

### 配置檔案（~/.tmux.conf）：
```shell
# 基礎設定
set -g default-terminal "screen-256color"
set -g display-time 3000
set -g escape-time 0
set -g history-limit 65535
set -g base-index 1
set -g pane-base-index 1

# 字首繫結 (Ctrl+a)
set -g prefix ^a
unbind ^b
bind a send-prefix

# 分割視窗
unbind '"'
bind - splitw -v
unbind %
bind | splitw -h

# 選中視窗
bind-key k select-pane -U
bind-key j select-pane -D
bind-key h select-pane -L
bind-key l select-pane -R

# copy-mode 將快捷鍵設定為 vi 模式
setw -g mode-keys vi

# 啟用滑鼠(Tmux v2.1)
set -g mouse on

# 更新配置檔案
bind r source-file ~/.tmux.conf \; display "已更新"

#<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
# Tmux Plugin Manager(Tmux v2.1)
# Tmux Resurrect
set -g @plugin 'tmux-plugins/tmux-resurrect'

# List of plugins
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'tmux-plugins/tmux-sensible'

# Other examples:
# set -g @plugin 'github_username/plugin_name'
# set -g @plugin '[email protected]/user/plugin'
# set -g @plugin '[email protected]/user/plugin'

# Initialize TMUX plugin manager (keep this line at the very bottom of tmux.conf)
run '~/.tmux/plugins/tpm/tpm'
#>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
```